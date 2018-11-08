package babble

import (
	"github.com/republicprotocol/babble-go/adapter/db"
	"github.com/republicprotocol/babble-go/adapter/rpc"
	"github.com/republicprotocol/babble-go/core/addr"
	"github.com/republicprotocol/babble-go/core/gossip"
)

type (
	Gossiper  = gossip.Gossiper
	Message   = gossip.Message
	Client    = gossip.Client
	Signer    = gossip.Signer
	Observer  = gossip.Observer
	Verifier  = gossip.Verifier
	Store     = gossip.Store
	AddrStore = addr.Store
	Book      = addr.Book
)

var (
	NewGossiper     = gossip.NewGossiper
	NewRPCClient    = rpc.NewClient
	NewRPCService   = rpc.NewService
	NewBook         = addr.NewBook
	NewMessageStore = db.NewMessageStore
	NewAddrs        = db.NewAddrStore
)
