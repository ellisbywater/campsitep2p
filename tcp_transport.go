package campsitep2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// TCPTransport implements the Transport interface.
type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshake     HandshakeFunc
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddress,
		peers:         make(map[net.Addr]Peer),
		handshake:     NOPHandshake,
	}
}

func (tt *TCPTransport) ListenAndAccept() error {
	l, err := net.Listen("tcp", tt.listenAddress)
	if err != nil {
		return err
	}
	tt.listener = l
	go tt.acceptConnections()
	return nil
}

func (tt *TCPTransport) Close() error {
	return tt.listener.Close()
}

// acceptConnections accepts incoming connections from remote peers.
func (tt *TCPTransport) acceptConnections() {
	for {
		conn, err := tt.listener.Accept()
		if err != nil {
			fmt.Println("acceptConnections error: ", err)
			continue
		}
		go tt.handleConnection(conn)
	}
}

func (tt *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	if err := tt.handshake(conn); err != nil {
		fmt.Println("handshake error: ", err)
		return
	}

	fmt.Printf("New connection from %s\n", peer.conn.RemoteAddr())
}
