package webrtc

import (
	"container/list"
	"sync"
	"sync/atomic"
)

type operation func()

type operations struct {
	mu     sync.Mutex
	busyCh chan struct{}
	ops    *list.List

	updateNegotiationNeededFlagOnEmptyChain *atomic.Bool
	onNegotiationNeeded                     func()
	isClosed                                bool
}

func newOperations(
	updateNegotiationNeededFlagOnEmptyChain *atomic.Bool,
	onNegotiationNeeded func(),
) *operations {
	_ = "STUB: not implemented"
	return nil
}

func (o *operations) Enqueue(op operation) { _ = "STUB: not implemented"; return }

func (o *operations) tryEnqueue(op operation) bool { _ = "STUB: not implemented"; return false }

func (o *operations) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (o *operations) Done() { _ = "STUB: not implemented"; return }

func (o *operations) GracefulClose() { _ = "STUB: not implemented"; return }

func (o *operations) pop() func() { _ = "STUB: not implemented"; return nil }

func (o *operations) start() { _ = "STUB: not implemented"; return }
