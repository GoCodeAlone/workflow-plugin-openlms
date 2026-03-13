package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type siteGetInfoStep struct {
	name       string
	moduleName string
}

func newSiteGetInfoStep(name string, config map[string]any) (*siteGetInfoStep, error) {
	return &siteGetInfoStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *siteGetInfoStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, _ map[string]any, _ map[string]any, _ map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	result, err := client.callToMap(ctx, "core_webservice_get_site_info", map[string]string{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type webserviceGetSiteInfoStep struct {
	name       string
	moduleName string
}

func newWebserviceGetSiteInfoStep(name string, config map[string]any) (*webserviceGetSiteInfoStep, error) {
	return &webserviceGetSiteInfoStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *webserviceGetSiteInfoStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, _ map[string]any, _ map[string]any, _ map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	result, err := client.callToMap(ctx, "core_webservice_get_site_info", map[string]string{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
