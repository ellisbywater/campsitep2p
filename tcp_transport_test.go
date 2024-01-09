package campsitep2p

import "testing"

func TestNewTCPTransport(t *testing.T) {
	tt := NewTCPTransport(":4000")
	if tt == nil {
		t.Error("NewTCPTransport() should not return nil")
	}

	if tt.listenAddress != ":4000" {
		t.Errorf("NewTCPTransport() should set listenAddress to %s, got %s", ":4000", tt.listenAddress)
	}

	if tt.peers == nil {
		t.Error("NewTCPTransport() should initialize peers map")
	}
	error := tt.ListenAndAccept()
	if error != nil {
		t.Errorf("NewTCPTransport() should not return error")
	}
}
