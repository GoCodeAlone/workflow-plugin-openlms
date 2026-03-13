package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type xapiStatementPostStep struct {
	name       string
	moduleName string
}

func newXAPIStatementPostStep(name string, config map[string]any) (*xapiStatementPostStep, error) {
	return &xapiStatementPostStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *xapiStatementPostStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	component := resolveValue("component", current, config)
	requestjson := resolveValue("requestjson", current, config)
	if component == "" || requestjson == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "component and requestjson are required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_xapi_statement_post", map[string]string{
		"component": component, "requestjson": requestjson,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type xapiGetStateStep struct {
	name       string
	moduleName string
}

func newXAPIGetStateStep(name string, config map[string]any) (*xapiGetStateStep, error) {
	return &xapiGetStateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *xapiGetStateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	component := resolveValue("component", current, config)
	activityID := resolveValue("activityid", current, config)
	agent := resolveValue("agent", current, config)
	stateID := resolveValue("stateid", current, config)
	if component == "" || activityID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "component and activityid are required"}}, nil
	}
	params := map[string]string{
		"component":  component,
		"activityId": activityID,
		"stateId":    stateID,
	}
	if agent != "" {
		params["agent"] = agent
	}
	result, err := client.callToMap(ctx, "core_xapi_get_state", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type xapiPostStateStep struct {
	name       string
	moduleName string
}

func newXAPIPostStateStep(name string, config map[string]any) (*xapiPostStateStep, error) {
	return &xapiPostStateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *xapiPostStateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	component := resolveValue("component", current, config)
	activityID := resolveValue("activityid", current, config)
	stateID := resolveValue("stateid", current, config)
	statedata := resolveValue("statedata", current, config)
	if component == "" || activityID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "component and activityid are required"}}, nil
	}
	params := map[string]string{
		"component":  component,
		"activityId": activityID,
		"stateId":    stateID,
		"stateData":  statedata,
	}
	if agent := resolveValue("agent", current, config); agent != "" {
		params["agent"] = agent
	}
	_, err := client.call(ctx, "core_xapi_post_state", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"posted": true}}, nil
}
