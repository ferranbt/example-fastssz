package example

import (
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/holiman/uint256"
)

type BidTrace struct {
	Slot                 uint64
	ParentHash           phase0.Hash32              `ssz-size:"32"`
	BlockHash            phase0.Hash32              `ssz-size:"32"`
	BuilderPubkey        phase0.BLSPubKey           `ssz-size:"48"`
	ProposerPubkey       phase0.BLSPubKey           `ssz-size:"48"`
	ProposerFeeRecipient bellatrix.ExecutionAddress `ssz-size:"20"`
	GasLimit             uint64
	GasUsed              uint64
	Value                *uint256.Int `ssz-size:"32"`
}
