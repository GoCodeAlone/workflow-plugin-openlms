package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type searchGetResultsStep struct {
	name       string
	moduleName string
}

func newSearchGetResultsStep(name string, config map[string]any) (*searchGetResultsStep, error) {
	return &searchGetResultsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *searchGetResultsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	query := resolveValue("query", current, config)
	if query == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "query is required"}}, nil
	}
	params := map[string]string{"q": query}
	if areaids := resolveValue("areaids", current, config); areaids != "" {
		params["areaids[0]"] = areaids
	}
	if courseids := resolveValue("courseids", current, config); courseids != "" {
		params["courseids[0]"] = courseids
	}
	result, err := client.callToMap(ctx, "core_search_get_results", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	results, _ := result["results"].([]any)
	return &sdk.StepResult{Output: map[string]any{"results": results, "totalcount": result["totalcount"]}}, nil
}
