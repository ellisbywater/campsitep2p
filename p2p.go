package campsitep2p

// Peer is the interface that wraps the basic methods of a remote peer.
type Peer interface {
}

// Transport is the interface that wraps the basic methods of Transport.
// Transport is responsible for sending and receiving messages to and from
// a remote peer.
// Transport can communicate via any protocol, e.g. TCP, UDP, HTTP, RPC, gRPC, etc.
type Transport interface {
	// ListenAndAccept listens for incoming connections from remote peers.
	ListenAndAccept() error
}
