/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package ws

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/internal/test/transportutil"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"

	commontransport "github.com/hyperledger/aries-framework-go/pkg/didcomm/common/transport"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/transport"
	mockpackager "github.com/hyperledger/aries-framework-go/pkg/internal/mock/didcomm/packager"
)

type mockProvider struct {
	packagerValue commontransport.Packager
}

func (p *mockProvider) InboundMessageHandler() transport.InboundMessageHandler {
	return func(message []byte) error {
		logger.Infof("message received is %s", string(message))
		if string(message) == "invalid-data" {
			return errors.New("error")
		}
		return nil
	}
}

func (p *mockProvider) Packager() commontransport.Packager {
	return p.packagerValue
}

func TestInboundTransport(t *testing.T) {
	t.Run("test inbound transport - with host/port", func(t *testing.T) {
		port := ":" + strconv.Itoa(transportutil.GetRandomPort(5))
		externalAddr := "http://example.com" + port
		inbound, err := NewInbound("localhost"+port, externalAddr)
		require.NoError(t, err)
		require.Equal(t, externalAddr, inbound.Endpoint())
	})

	t.Run("test inbound transport - with host/port, no external address", func(t *testing.T) {
		internalAddr := "example.com" + ":" + strconv.Itoa(transportutil.GetRandomPort(5))
		inbound, err := NewInbound(internalAddr, "")
		require.NoError(t, err)
		require.Equal(t, internalAddr, inbound.Endpoint())
	})

	t.Run("test inbound transport - without host/port", func(t *testing.T) {
		inbound, err := NewInbound(":"+strconv.Itoa(transportutil.GetRandomPort(5)), "")
		require.NoError(t, err)
		require.NotEmpty(t, inbound)
		mockPackager := &mockpackager.Packager{UnpackValue: &commontransport.Envelope{Message: []byte("data")}}
		err = inbound.Start(&mockProvider{packagerValue: mockPackager})
		require.NoError(t, err)

		err = inbound.Stop()
		require.NoError(t, err)
	})

	t.Run("test inbound transport - nil context", func(t *testing.T) {
		inbound, err := NewInbound(":"+strconv.Itoa(transportutil.GetRandomPort(5)), "")
		require.NoError(t, err)
		require.NotEmpty(t, inbound)

		err = inbound.Start(nil)
		require.Error(t, err)
	})

	t.Run("test inbound transport - invalid port number", func(t *testing.T) {
		_, err := NewInbound("", "")
		require.Error(t, err)
		require.Contains(t, err.Error(), "websocket address is mandatory")
	})
}

func TestInboundDataProcessing(t *testing.T) {
	t.Run("test inbound transport - multiple invocation with same client", func(t *testing.T) {
		port := ":" + strconv.Itoa(transportutil.GetRandomPort(5))

		// initiate inbound with port
		inbound, err := NewInbound(port, "")
		require.NoError(t, err)
		require.NotEmpty(t, inbound)

		// start server
		mockPackager := &mockpackager.Packager{UnpackValue: &commontransport.Envelope{Message: []byte("valid-data")}}
		err = inbound.Start(&mockProvider{packagerValue: mockPackager})
		require.NoError(t, err)

		// create ws client
		client, cleanup := websocketClient(t, port)
		defer cleanup()

		for i := 1; i <= 5; i++ {
			err = client.WriteMessage(websocket.TextMessage, []byte("random"))
			require.NoError(t, err)

			messageType, val, err := client.ReadMessage()
			require.NoError(t, err)
			require.Equal(t, messageType, websocket.TextMessage)
			require.Equal(t, "", string(val))
		}
	})

	t.Run("test inbound transport - unpacking error", func(t *testing.T) {
		port := ":" + strconv.Itoa(transportutil.GetRandomPort(5))

		// initiate inbound with port
		inbound, err := NewInbound(port, "")
		require.NoError(t, err)
		require.NotEmpty(t, inbound)

		// start server
		mockPackager := &mockpackager.Packager{UnpackErr: errors.New("error unpacking")}
		err = inbound.Start(&mockProvider{packagerValue: mockPackager})
		require.NoError(t, err)

		// create ws client
		client, cleanup := websocketClient(t, port)
		defer cleanup()

		err = client.WriteMessage(websocket.TextMessage, []byte(""))
		require.NoError(t, err)

		messageType, val, err := client.ReadMessage()
		require.NoError(t, err)
		require.Equal(t, messageType, websocket.TextMessage)
		require.Equal(t, processFailureErrMsg, string(val))
	})

	t.Run("test inbound transport - message handler error", func(t *testing.T) {
		port := ":" + strconv.Itoa(transportutil.GetRandomPort(5))

		// initiate inbound with port
		inbound, err := NewInbound(port, "")
		require.NoError(t, err)
		require.NotEmpty(t, inbound)

		// start server
		mockPackager := &mockpackager.Packager{UnpackValue: &commontransport.Envelope{Message: []byte("invalid-data")}}
		err = inbound.Start(&mockProvider{packagerValue: mockPackager})
		require.NoError(t, err)

		// create ws client
		client, cleanup := websocketClient(t, port)
		defer cleanup()

		err = client.WriteMessage(websocket.TextMessage, []byte(""))
		require.NoError(t, err)

		messageType, val, err := client.ReadMessage()
		require.NoError(t, err)
		require.Equal(t, messageType, websocket.TextMessage)
		require.Equal(t, processFailureErrMsg, string(val))
	})
}

func websocketClient(t *testing.T, port string) (*websocket.Conn, func()) {
	require.NoError(t, transportutil.VerifyListener("localhost"+port, time.Second))

	u := url.URL{Scheme: "ws", Host: "localhost" + port, Path: ""}
	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	require.NoError(t, err)
	require.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)
	require.NoError(t, resp.Body.Close())

	return c, func() {
		require.NoError(t, c.Close())
	}
}