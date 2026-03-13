package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type badgeGetUserBadgesStep struct {
	name       string
	moduleName string
}

func newBadgeGetUserBadgesStep(name string, config map[string]any) (*badgeGetUserBadgesStep, error) {
	return &badgeGetUserBadgesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *badgeGetUserBadgesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if userID := resolveValue("userid", current, config); userID != "" {
		params["userid"] = userID
	}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseid"] = courseID
	}
	result, err := client.callToMap(ctx, "core_badges_get_user_badges", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	badges, _ := result["badges"].([]any)
	return &sdk.StepResult{Output: map[string]any{"badges": badges, "warnings": result["warnings"]}}, nil
}
