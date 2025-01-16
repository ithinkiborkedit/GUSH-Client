package storage

import (
	"errors"
	"sync"

	dclient "github.com/ithinkiborkedit/GUSH-Client.git/domain/client"
)

type InMemoryPlayerRepo struct {
	mu     sync.RWMutex
	player *dclient.LocalPlayer
}

func NewInMemoryPlayerRepo() *InMemoryPlayerRepo {
	return &InMemoryPlayerRepo{}
}

func (r *InMemoryPlayerRepo) GetLocalPlayer() (*dclient.LocalPlayer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.player == nil {
		return nil, errors.New("no player set")
	}
	return r.player, nil
}

func (r *InMemoryPlayerRepo) SaveLocalPlayer(p *dclient.LocalPlayer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.player = p

	return nil
}

//__

type InMemoryWorldRepo struct {
	mu    sync.RWMutex
	world *dclient.World
}

func NewInMemoryWorldRepo() *InMemoryWorldRepo {
	return &InMemoryWorldRepo{}
}

func (r *InMemoryWorldRepo) GetWorld() (*dclient.World, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.world == nil {
		r.world = dclient.NewWorld()
	}
	return r.world, nil
}

func (r *InMemoryWorldRepo) SaveWorld(w *dclient.World) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.world = w
	return nil
}
