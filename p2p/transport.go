package p2p

// Peer represents the remote node
type Peer interface {

}

// Transport handles communication between nodes in the network
type Transport interface {
	ListenAndAccept() error
}