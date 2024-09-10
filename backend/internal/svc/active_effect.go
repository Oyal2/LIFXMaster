package svc

import (
	"context"
	"sync"
)

type ActiveEffects struct {
	sync.RWMutex
	EffectCancel map[uint64]context.CancelFunc
}

func NewActiveEffects() *ActiveEffects {
	return &ActiveEffects{
		EffectCancel: make(map[uint64]context.CancelFunc),
	}
}

func (ae *ActiveEffects) CancelEffect(id uint64) {
	ae.Lock()
	defer ae.Unlock()

	cancel, ok := ae.EffectCancel[id]
	if ok {
		cancel()
		delete(ae.EffectCancel, id)
	}
}

func (ae *ActiveEffects) Add(id uint64, cancel context.CancelFunc) {
	ae.Lock()
	defer ae.Unlock()
	ae.EffectCancel[id] = cancel
}
