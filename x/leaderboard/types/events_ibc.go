package types

// IBC events
const (
	EventTypeTimeout          = "timeout"
	EventTypeIbcTopRankPacket = "ibcTopRank_packet"
	EventTypeCandidatePacket  = "candidate_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
