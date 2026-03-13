package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type ltiGetByCourseStep struct {
	name       string
	moduleName string
}

func newLtiGetByCourseStep(name string, config map[string]any) (*ltiGetByCourseStep, error) {
	return &ltiGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *ltiGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_lti_get_ltis_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	ltis, _ := result["ltis"].([]any)
	return &sdk.StepResult{Output: map[string]any{"ltis": ltis}}, nil
}

type ltiGetToolLaunchDataStep struct {
	name       string
	moduleName string
}

func newLtiGetToolLaunchDataStep(name string, config map[string]any) (*ltiGetToolLaunchDataStep, error) {
	return &ltiGetToolLaunchDataStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *ltiGetToolLaunchDataStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	toolid := resolveValue("toolid", current, config)
	if toolid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "toolid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lti_get_tool_launch_data", map[string]string{"toolid": toolid})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type ltiGetToolTypesStep struct {
	name       string
	moduleName string
}

func newLtiGetToolTypesStep(name string, config map[string]any) (*ltiGetToolTypesStep, error) {
	return &ltiGetToolTypesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *ltiGetToolTypesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, _ map[string]any, _ map[string]any, _ map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	result, err := client.callToMap(ctx, "mod_lti_get_tool_types", map[string]string{})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	types, _ := result["types"].([]any)
	return &sdk.StepResult{Output: map[string]any{"types": types}}, nil
}
