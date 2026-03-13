package internal

import (
	"sync"
)

var (
	clientMu       sync.RWMutex
	clientRegistry = make(map[string]*moodleClient)
)

// RegisterClient adds a Moodle client to the global registry under the given name.
func RegisterClient(name string, c *moodleClient) {
	clientMu.Lock()
	defer clientMu.Unlock()
	clientRegistry[name] = c
}

// GetClient looks up a Moodle client by name.
func GetClient(name string) (*moodleClient, bool) {
	clientMu.RLock()
	defer clientMu.RUnlock()
	c, ok := clientRegistry[name]
	return c, ok
}

// UnregisterClient removes a client from the registry.
func UnregisterClient(name string) {
	clientMu.Lock()
	defer clientMu.Unlock()
	delete(clientRegistry, name)
}
