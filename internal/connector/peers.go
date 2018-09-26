package connector

var peers = []PeerReceiver{}

type PeerReceiver struct {
	Peer     Peer
	Receiver Peer
}

func AddPeer(peer Peer, receiver Peer) {
	peerReceiver := PeerReceiver{peer, receiver}
	peers = append(peers, peerReceiver)
}

func RemovePeer(peer *Peer) {
	index := findPeerIndex(peer)
	removePeerWithIndex(index)
}

func findPeerIndex(peer *Peer) (index int) {
	for i, iterator := range peers {
		if &iterator.Peer == peer {
			index = i
			break
		}
	}
	return
}

func removePeerWithIndex(index int) {
	peers[index] = peers[len(peers)-1]
	peers = peers[:len(peers)-1]
}
