package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type reportBuilderListStep struct {
	name       string
	moduleName string
}

func newReportBuilderListStep(name string, config map[string]any) (*reportBuilderListStep, error) {
	return &reportBuilderListStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *reportBuilderListStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, _ map[string]any, _ map[string]any, _ map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	result, err := client.callToSlice(ctx, "core_reportbuilder_list_reports", map[string]string{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"reports": result, "count": len(result)}}, nil
}

type reportBuilderGetStep struct {
	name       string
	moduleName string
}

func newReportBuilderGetStep(name string, config map[string]any) (*reportBuilderGetStep, error) {
	return &reportBuilderGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *reportBuilderGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	reportID := resolveValue("reportid", current, config)
	if reportID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "reportid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_reportbuilder_retrieve_report", map[string]string{"reportid": reportID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type reportBuilderRetrieveStep struct {
	name       string
	moduleName string
}

func newReportBuilderRetrieveStep(name string, config map[string]any) (*reportBuilderRetrieveStep, error) {
	return &reportBuilderRetrieveStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *reportBuilderRetrieveStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	reportID := resolveValue("reportid", current, config)
	if reportID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "reportid is required"}}, nil
	}
	params := map[string]string{"reportid": reportID}
	if page := resolveValue("page", current, config); page != "" {
		params["page"] = page
	}
	if perpage := resolveValue("perpage", current, config); perpage != "" {
		params["perpage"] = perpage
	}
	result, err := client.callToMap(ctx, "core_reportbuilder_retrieve_report", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
