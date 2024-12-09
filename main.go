package main

import (
	"fmt"

	v1 "github.com/attestantio/go-builder-client/api/v1"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/bellatrix"
	"github.com/attestantio/go-eth2-client/spec/capella"
	deneb_869bb "github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
)

func do_panic() {
	/*
		github.com/attestantio/go-eth2-client/spec/phase0.(*AttesterSlashing).SizeSSZ(...)
			/0xtylerholmes/git/go-eth2-client/spec/phase0/attesterslashing_ssz.go:101
		github.com/attestantio/go-eth2-client/spec/deneb.(*BeaconBlockBody).SizeSSZ(0xc0000d6640)
			/0xtylerholmes/git/go-eth2-client/spec/deneb/beaconblockbody_ssz.go:413 +0x75
		github.com/attestantio/go-eth2-client/spec/deneb.(*BeaconBlock).SizeSSZ(0xc00008eeb0?)
			/0xtylerholmes/git/go-eth2-client/spec/deneb/beaconblock_ssz.go:101 +0x5b
		github.com/ferranbt/fastssz.MarshalSSZ({0x712ae0, 0xc0000be1e0})
			/home/user/go/pkg/mod/github.com/ferranbt/fastssz@v0.1.4/encode.go:12 +0x22
		github.com/attestantio/go-eth2-client/spec/deneb.(*BeaconBlock).MarshalSSZ(...)
			/0xtylerholmes/git/go-eth2-client/spec/deneb/beaconblock_ssz.go:13
		github.com/attestantio/go-eth2-client.do_panic()
	*/

	/*
		Similiar to previous one we enter into .SizeSSZ() which doesn't know that the receiver is nil
		@ go-eth2-client/spec/phase0/attesterslashing_ssz.go:101
			// Field (0) 'Attestation1'
			if a.Attestation1 == nil { 			<- receiver is nil
				a.Attestation1 = new(IndexedAttestation)
			}
	*/

	block := &deneb_869bb.BeaconBlock{
		Slot:          0,
		ProposerIndex: 0,
		ParentRoot:    phase0.Root{},
		StateRoot:     phase0.Root{},
		Body: &deneb_869bb.BeaconBlockBody{
			RANDAOReveal:      phase0.BLSSignature{},
			ETH1Data:          &phase0.ETH1Data{},
			Graffiti:          [32]byte{},
			ProposerSlashings: []*phase0.ProposerSlashing{},
			AttesterSlashings: []*phase0.AttesterSlashing{
				new(phase0.AttesterSlashing),
				nil,
				new(phase0.AttesterSlashing),
			},
			Attestations:          []*phase0.Attestation{},
			Deposits:              []*phase0.Deposit{},
			VoluntaryExits:        []*phase0.SignedVoluntaryExit{},
			SyncAggregate:         &altair.SyncAggregate{},
			ExecutionPayload:      &deneb_869bb.ExecutionPayload{},
			BLSToExecutionChanges: []*capella.SignedBLSToExecutionChange{},
			BlobKZGCommitments:    []deneb_869bb.KZGCommitment{},
		},
	}

	_, _ = block.MarshalSSZ()

}

func do_panic_2() {
	bidTrace := v1.BidTrace{
		Slot:                 0,
		ParentHash:           phase0.Hash32{},
		BlockHash:            phase0.Hash32{},
		BuilderPubkey:        phase0.BLSPubKey{},
		ProposerPubkey:       phase0.BLSPubKey{},
		ProposerFeeRecipient: bellatrix.ExecutionAddress{},
		GasLimit:             0,
		GasUsed:              0,
		Value:                nil,
	}

	_, err := bidTrace.MarshalSSZ()

	if err != nil {
		fmt.Println(fmt.Sprintf("failed to marshal object: %s", err.Error()))
	}
}

func main() {
	do_panic_2()
}
