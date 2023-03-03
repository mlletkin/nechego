package game

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Universe holds worlds.
type Universe struct {
	// worlds is a map of active (loaded) worlds accessed by TGID.
	worlds map[int64]*World

	// dir is a location of the directory where worlds should be saved.
	dir string

	// initWorld is called on the first world's load.
	initWorld func(*World)

	mu sync.Mutex
}

// NewUniverse returns a new Universe with no worlds.
func NewUniverse(dir string, initWorld func(*World)) *Universe {
	return &Universe{
		dir:       dir,
		worlds:    map[int64]*World{},
		initWorld: initWorld,
	}
}

// worldPath returns the location of a save file of the world by the
// specified ID.
func (u *Universe) worldPath(id int64) string {
	return filepath.Join(u.dir, fmt.Sprintf("world%d.json", id))
}

// ForEachWorld applies action to each world in the universe.
func (u *Universe) ForEachWorld(action func(*World)) {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, w := range u.worlds {
		w.Lock()
		action(w)
		w.Unlock()
	}
}

// World returns the world by the given ID from the universe. If the
// world is not active, loads the save from the world's file. If there
// is no save file found, creates a new world.
func (u *Universe) World(id int64) (*World, error) {
	u.mu.Lock()
	defer u.mu.Unlock()

	w, ok := u.worlds[id]
	if !ok {
		// Invariant: the world in not initialized.
		// This case holds only once for each world.
		var err error
		w, err = LoadWorld(u.worldPath(id))
		if errors.Is(err, os.ErrNotExist) {
			w = NewWorld(id)
		} else if err != nil {
			return nil, err
		}
		w.migrate()
		u.initWorld(w)
		u.worlds[id] = w
	}
	return w, nil
}

// SaveAll saves all active worlds.
func (u *Universe) SaveAll() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	for _, w := range u.worlds {
		if err := w.Save(u.worldPath(w.TGID)); err != nil {
			return err
		}
	}
	return nil
}