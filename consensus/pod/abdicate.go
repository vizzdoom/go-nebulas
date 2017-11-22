// Copyright (C) 2017 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package pod

import (
	"fmt"

	"github.com/nebulasio/go-nebulas/consensus"
	"github.com/nebulasio/go-nebulas/util/byteutils"
	log "github.com/sirupsen/logrus"
)

// AbdicateContext carray abdicate info
type AbdicateContext struct {
	Voter     byteutils.Hash
	BlockHash byteutils.Hash
}

// AbdicateState presents the prepare stage in pod
type AbdicateState struct {
	sm *consensus.StateMachine
}

// NewAbdicateState create a new prepare state
func NewAbdicateState(sm *consensus.StateMachine) *AbdicateState {
	return &AbdicateState{
		sm: sm,
	}
}

func (state *AbdicateState) String() string {
	return fmt.Sprintf("AbdicateState %p", state)
}

// Event handle event.
func (state *AbdicateState) Event(e consensus.Event) (bool, consensus.State) {
	switch e.EventType() {
	case NewAbdicateVoteEvent:
		return true, nil
	}
	return false, nil
}

// Enter called when transiting to this state.
func (state *AbdicateState) Enter(data interface{}) {
	log.Debug("AbdicateState enter.")
	// if the block is on canonical chain, vote
}

// Leave called when leaving this state.
func (state *AbdicateState) Leave(data interface{}) {
	log.Debug("AbdicateState leave.")
}
