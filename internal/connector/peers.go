package connector

type PeerReceiver struct {
	Peer     Peer
	Receiver Peer
}

type PeerReceivers []PeerReceiver

func (peerReceivers *PeerReceivers) AddPeer(peer Peer, receiver Peer) {
	peerReceiver := PeerReceiver{peer, receiver}
	*peerReceivers = append(*peerReceivers, peerReceiver)
}

func (peerReceivers *PeerReceivers) RemovePeer(peer *Peer) {
	index := peerReceivers.findPeerIndex(peer)
	peerReceivers.removePeerWithIndex(index)
}

func (peerReceivers *PeerReceivers) findPeerIndex(peer *Peer) (index int) {
	for i, iterator := range *peerReceivers {
		if &iterator.Peer == peer {
			index = i
			break
		}
	}
	return
}

func (peerReceivers *PeerReceivers) removePeerWithIndex(index int) {
	(*peerReceivers)[index] = (*peerReceivers)[len(*peerReceivers)-1]
	*peerReceivers = (*peerReceivers)[:len(*peerReceivers)-1]
}
