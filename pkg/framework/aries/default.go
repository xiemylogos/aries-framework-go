/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package aries

import (
	"fmt"

	"github.com/hyperledger/aries-framework-go/pkg/didcomm/dispatcher"
	"github.com/hyperledger/aries-framework-go/pkg/didcomm/protocol/didexchange"
	"github.com/hyperledger/aries-framework-go/pkg/didmethod/peer"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/api"
	"github.com/hyperledger/aries-framework-go/pkg/framework/aries/factory/transport"
	"github.com/hyperledger/aries-framework-go/pkg/framework/didresolver"
	"github.com/hyperledger/aries-framework-go/pkg/storage"
	"github.com/hyperledger/aries-framework-go/pkg/storage/leveldb"
)

// DBPath Level DB Path.
// TODO - Need to configure the path externally (#148 & #175)
var DBPath = "/tmp/peerstore/"

// defFramework provides default framework configs
type defFramework struct {
	storeProv storage.Provider
}

// transportProviderFactory provides default Outbound Transport provider factory
func (d *defFramework) transportProviderFactory() api.TransportProviderFactory {
	return transport.NewProviderFactory()
}

// didResolverProvider provides default DID resolver.
func (d *defFramework) didResolverProvider() (DIDResolver, error) {
	dbprov, err := d.storeProvider()
	if err != nil {
		return nil, fmt.Errorf("resolver initialization failed : %w", err)
	}

	dbstore, err := dbprov.GetStoreHandle()
	if err != nil {
		return nil, fmt.Errorf("storage initialization failed : %w", err)
	}

	resl := didresolver.New(didresolver.WithDidMethod(peer.NewDIDResolver(peer.NewDIDStore(dbstore))))
	return resl, nil
}

func (d *defFramework) storeProvider() (storage.Provider, error) {
	if d.storeProv != nil {
		return d.storeProv, nil
	}

	// TODO - Need to configure the path externally
	storeProv, err := leveldb.NewProvider(DBPath)
	if err != nil {
		return nil, fmt.Errorf("leveldb provider initialization failed : %w", err)
	}

	d.storeProv = storeProv
	return storeProv, nil
}

// defFrameworkOpts provides default framework options
func defFrameworkOpts() ([]Option, error) {
	// get the default framework configs
	def := defFramework{}

	var opts []Option
	// protocol provider factory
	opt := WithTransportProviderFactory(def.transportProviderFactory())
	opts = append(opts, opt)

	reslv, err := def.didResolverProvider()
	if err != nil {
		return nil, fmt.Errorf("resolver initialization failed : %w", err)
	}
	opts = append(opts, WithDIDResolver(reslv))

	storeProv, err := def.storeProvider()
	if err != nil {
		return nil, fmt.Errorf("resolver initialization failed : %w", err)
	}
	opts = append(opts, WithStoreProvider(storeProv))

	store, err := storeProv.GetStoreHandle()
	if err != nil {
		return nil, fmt.Errorf("resolver initialization failed : %w", err)
	}

	// default protocols
	newExchangeSvc := func(prv api.Provider) (dispatcher.Service, error) { return didexchange.New(store, prv), nil }

	opts = append(opts, WithProtocols(newExchangeSvc))

	return opts, nil
}