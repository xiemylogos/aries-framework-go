/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

import (
	"io"
	"net/http"

	"github.com/hyperledger/aries-framework-go/pkg/controller/command"
	cmdkms "github.com/hyperledger/aries-framework-go/pkg/controller/command/kms"
	"github.com/hyperledger/aries-framework-go/pkg/controller/internal/cmdutil"
	"github.com/hyperledger/aries-framework-go/pkg/controller/rest"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	"github.com/hyperledger/aries-framework-go/pkg/kms/legacykms"
)

const (
	kmseOperationID           = "/kms"
	legacykmseOperationID     = "/legacykms"
	createKeySetPath          = kmseOperationID + "/keyset"
	importKeyPath             = kmseOperationID + "/import"
	createKeySetLegacyKMSPath = legacykmseOperationID + "/keyset"
)

// provider contains dependencies for the kms command and is typically created by using aries.Context().
type provider interface {
	KMS() kms.KeyManager
	LegacyKMS() legacykms.KeyManager
}

type kmsCommand interface {
	CreateKeySet(rw io.Writer, req io.Reader) command.Error
	CreateKeySetLegacyKMS(rw io.Writer, req io.Reader) command.Error
	ImportKey(rw io.Writer, req io.Reader) command.Error
}

// Operation contains basic common operations provided by controller REST API
type Operation struct {
	handlers []rest.Handler
	command  kmsCommand
}

// New returns new kms operations rest client instance
func New(p provider) *Operation {
	cmd := cmdkms.New(p)

	o := &Operation{command: cmd}
	o.registerHandler()

	return o
}

// GetRESTHandlers get all controller API handler available for this service
func (o *Operation) GetRESTHandlers() []rest.Handler {
	return o.handlers
}

// registerHandler register handlers to be exposed from this protocol service as REST API endpoints
func (o *Operation) registerHandler() {
	o.handlers = []rest.Handler{
		cmdutil.NewHTTPHandler(createKeySetPath, http.MethodPost, o.CreateKeySet),
		cmdutil.NewHTTPHandler(importKeyPath, http.MethodPost, o.ImportKey),
		cmdutil.NewHTTPHandler(createKeySetLegacyKMSPath, http.MethodPost, o.CreateKeySetLegacyKms),
	}
}

// CreateKeySetLegacyKms swagger:route POST /legacykms/keyset legacykms createKeySetLegacyKMS
//
// Create key set.
//
// Responses:
//    default: genericError
//        200: createKeySetRes
func (o *Operation) CreateKeySetLegacyKms(rw http.ResponseWriter, req *http.Request) {
	rest.Execute(o.command.CreateKeySetLegacyKMS, rw, req.Body)
}

// CreateKeySet swagger:route POST /kms/keyset kms createKeySet
//
// Create key set.
//
// Responses:
//    default: genericError
//        200: createKeySetRes
func (o *Operation) CreateKeySet(rw http.ResponseWriter, req *http.Request) {
	rest.Execute(o.command.CreateKeySet, rw, req.Body)
}

// ImportKey swagger:route POST /kms/import kms importKey
//
// Import key.
//
// Responses:
//    default: genericError
func (o *Operation) ImportKey(rw http.ResponseWriter, req *http.Request) {
	rest.Execute(o.command.ImportKey, rw, req.Body)
}
