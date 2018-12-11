package custom

import (
	"github.com/coschain/go-bft/message"
)

/*
 * A validator is a node in a distributed system that participates in the
 * bft consensus process. It proposes and votes for a certain proposal.
 * Each validator should maintain a set of all the PubValidators so that
 * it can verifies messages sent by other validators. Each validator should
 * have exactly one PrivValidator which contains its private key so that
 * it can sign a message. A validator can be a proposer, the rules of which
 * validator becomes a valid proposer at a certain time is totally decided by
 * user.
 */

// PubKey is binary representation of the public key
type PubKey []byte

// Proposer stages a candidate data so that other validators can vote for it
type Proposer interface {
	GetCurrentProposer()
	Propose() (*message.Proposal, error)
	// Each Validator will vote for the POLed proposal if there's any. Otherwise it
	// votes for the first proposal it sees in default unless user explicitly calls
	// BoundVotedData(data), in which case it votes for the bounded data.
	BoundVotedData(data []byte)
}

// PubValidator verifies if a message is properly signed by the right validator
type PubValidator interface {
	VerifySig(digest, signature []byte) bool
	GetPubKey() PubKey
}

// PrivValidator signs a message
type PrivValidator interface {
	GetPubKeyBytes() PubKey
	Sign(digest []byte) []byte
}

// Committer defines the actions the users taken when consensus is reached
type Committer interface {

}