package internal

import (
	"context"
	"fmt"
)

// openLMSModule creates a Moodle client and registers it.
type openLMSModule struct {
	name   string
	config map[string]any
}

func newOpenLMSModule(name string, config map[string]any) (*openLMSModule, error) {
	return &openLMSModule{name: name, config: config}, nil
}

// Init creates the Moodle REST client and registers it in the global registry.
func (m *openLMSModule) Init() error {
	siteURL, _ := m.config["siteUrl"].(string)
	token, _ := m.config["token"].(string)

	if siteURL == "" {
		return fmt.Errorf("openlms.provider %q: siteUrl is required", m.name)
	}
	if token == "" {
		return fmt.Errorf("openlms.provider %q: token is required", m.name)
	}

	restful := false
	if v, ok := m.config["restful"].(bool); ok {
		restful = v
	} else if v, ok := m.config["restful"].(string); ok {
		restful = v == "true" || v == "1" || v == "yes"
	}

	client := newMoodleClient(siteURL, token, restful)
	RegisterClient(m.name, client)
	return nil
}

// Start is a no-op for this module.
func (m *openLMSModule) Start(_ context.Context) error { return nil }

// Stop unregisters the Moodle client.
func (m *openLMSModule) Stop(_ context.Context) error {
	UnregisterClient(m.name)
	return nil
}
