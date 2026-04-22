// Package internal implements the workflow-plugin-openlms plugin.
package internal

import (
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// Version is set at build time via -ldflags
// "-X github.com/GoCodeAlone/workflow-plugin-openlms/internal.Version=X.Y.Z".
// Default is a bare semver so plugin loaders that validate semver accept
// unreleased dev builds; goreleaser overrides with the real release tag.
var Version = "0.0.0"

// openLMSPlugin implements sdk.PluginProvider, sdk.ModuleProvider, and sdk.StepProvider.
type openLMSPlugin struct{}

// NewOpenLMSPlugin returns a new openLMSPlugin instance.
func NewOpenLMSPlugin() sdk.PluginProvider {
	return &openLMSPlugin{}
}

// Manifest returns plugin metadata.
func (p *openLMSPlugin) Manifest() sdk.PluginManifest {
	return sdk.PluginManifest{
		Name:        "workflow-plugin-openlms",
		Version:     Version,
		Author:      "GoCodeAlone",
		Description: "OpenLMS/Moodle LMS plugin (~120 step types across all Moodle Web Services APIs)",
	}
}

// ModuleTypes returns the module type names this plugin provides.
func (p *openLMSPlugin) ModuleTypes() []string {
	return []string{"openlms.provider"}
}

// CreateModule creates a module instance of the given type.
func (p *openLMSPlugin) CreateModule(typeName, name string, config map[string]any) (sdk.ModuleInstance, error) {
	switch typeName {
	case "openlms.provider":
		m, err := newOpenLMSModule(name, config)
		if err != nil {
			return nil, err
		}
		return m, nil
	default:
		return nil, fmt.Errorf("openlms plugin: unknown module type %q", typeName)
	}
}

// StepTypes returns the step type names this plugin provides.
func (p *openLMSPlugin) StepTypes() []string {
	return allStepTypes()
}

// CreateStep creates a step instance of the given type.
func (p *openLMSPlugin) CreateStep(typeName, name string, config map[string]any) (sdk.StepInstance, error) {
	return createStep(typeName, name, config)
}
