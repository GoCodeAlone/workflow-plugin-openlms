package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type h5pGetByCourseStep struct {
	name       string
	moduleName string
}

func newH5PGetByCourseStep(name string, config map[string]any) (*h5pGetByCourseStep, error) {
	return &h5pGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *h5pGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_h5pactivity_get_h5pactivities_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	activities, _ := result["h5pactivities"].([]any)
	return &sdk.StepResult{Output: map[string]any{"h5pactivities": activities}}, nil
}

type h5pGetAttemptsStep struct {
	name       string
	moduleName string
}

func newH5PGetAttemptsStep(name string, config map[string]any) (*h5pGetAttemptsStep, error) {
	return &h5pGetAttemptsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *h5pGetAttemptsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	h5pactivityid := resolveValue("h5pactivityid", current, config)
	if h5pactivityid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "h5pactivityid is required"}}, nil
	}
	params := map[string]string{"h5pactivityid": h5pactivityid}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userid"] = userID
	}
	result, err := client.callToMap(ctx, "mod_h5pactivity_get_attempts", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	attempts, _ := result["attempts"].([]any)
	return &sdk.StepResult{Output: map[string]any{"attempts": attempts}}, nil
}

type h5pGetResultsStep struct {
	name       string
	moduleName string
}

func newH5PGetResultsStep(name string, config map[string]any) (*h5pGetResultsStep, error) {
	return &h5pGetResultsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *h5pGetResultsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	h5pactivityid := resolveValue("h5pactivityid", current, config)
	attemptids := resolveValue("attemptid", current, config)
	if h5pactivityid == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "h5pactivityid is required"}}, nil
	}
	params := map[string]string{"h5pactivityid": h5pactivityid}
	if attemptids != "" {
		params["attemptids[0]"] = attemptids
	}
	result, err := client.callToMap(ctx, "mod_h5pactivity_get_results", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
