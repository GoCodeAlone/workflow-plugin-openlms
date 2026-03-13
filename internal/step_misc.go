package internal

import (
	"context"
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// Generic call function step

type callFunctionStep struct {
	name       string
	moduleName string
}

func newCallFunctionStep(name string, config map[string]any) (*callFunctionStep, error) {
	return &callFunctionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *callFunctionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	wsfunction := resolveValue("wsfunction", current, config)
	if wsfunction == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "wsfunction is required"}}, nil
	}
	params := map[string]string{}
	if p := resolveMap("params", current, config); p != nil {
		for k, v := range p {
			params[k] = fmt.Sprintf("%v", v)
		}
	}
	result, err := client.call(ctx, wsfunction, params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	if m, ok := result.(map[string]any); ok {
		return &sdk.StepResult{Output: m}, nil
	}
	if sl, ok := result.([]any); ok {
		return &sdk.StepResult{Output: map[string]any{"result": sl}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"result": result}}, nil
}
