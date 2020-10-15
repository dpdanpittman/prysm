package helpers

import (
	"github.com/pkg/errors"
	stateTrie "github.com/prysmaticlabs/prysm/beacon-chain/state"
	"github.com/prysmaticlabs/prysm/shared/params"
	"github.com/prysmaticlabs/prysm/shared/types"
)

// BlockRootAtSlot returns the block root stored in the BeaconState for a recent slot.
// It returns an error if the requested block root is not within the slot range.
//
// Spec pseudocode definition:
//  def get_block_root_at_slot(state: BeaconState, slot: Slot) -> Hash:
//    """
//    Return the block root at a recent ``slot``.
//    """
//    assert slot < state.slot <= slot + SLOTS_PER_HISTORICAL_ROOT
//    return state.block_roots[slot % SLOTS_PER_HISTORICAL_ROOT]
func BlockRootAtSlot(state *stateTrie.BeaconState, slot types.Slot) ([]byte, error) {
	if slot >= state.Slot() || state.Slot() > slot+params.BeaconConfig().SlotsPerHistoricalRoot {
		return []byte{}, errors.Errorf("slot %d out of bounds", slot)
	}
	return state.BlockRootAtIndex(slot.Uint64() % params.BeaconConfig().SlotsPerHistoricalRoot.Uint64())
}

// BlockRoot returns the block root stored in the BeaconState for epoch start slot.
//
// Spec pseudocode definition:
//  def get_block_root(state: BeaconState, epoch: Epoch) -> Hash:
//    """
//    Return the block root at the start of a recent ``epoch``.
//    """
//    return get_block_root_at_slot(state, compute_start_slot_at_epoch(epoch))
func BlockRoot(state *stateTrie.BeaconState, epoch types.Epoch) ([]byte, error) {
	s, err := StartSlot(epoch)
	if err != nil {
		return nil, err
	}
	return BlockRootAtSlot(state, s)
}
