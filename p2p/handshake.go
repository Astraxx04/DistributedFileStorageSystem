package p2p

import "errors"

// This is returned if the connection between local and remote node could not be established
var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func (any) error

func NOPHandshakeFunc(any) error { 
	return nil 
}