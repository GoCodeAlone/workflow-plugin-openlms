package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type scormGetByCourseStep struct {
	name       string
	moduleName string
}

func newScormGetByCourseStep(name string, config map[string]any) (*scormGetByCourseStep, error) {
	return &scormGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_scorm_get_scorms_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	scorms, _ := result["scorms"].([]any)
	return &sdk.StepResult{Output: map[string]any{"scorms": scorms}}, nil
}

type scormGetAttemptCountStep struct {
	name       string
	moduleName string
}

func newScormGetAttemptCountStep(name string, config map[string]any) (*scormGetAttemptCountStep, error) {
	return &scormGetAttemptCountStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormGetAttemptCountStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	scormID := resolveValue("scormid", current, config)
	userID := resolveValue("userid", current, config)
	if scormID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "scormid and userid are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_scorm_get_scorm_attempt_count", map[string]string{
		"scormid": scormID, "userid": userID,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type scormGetScosStep struct {
	name       string
	moduleName string
}

func newScormGetScosStep(name string, config map[string]any) (*scormGetScosStep, error) {
	return &scormGetScosStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormGetScosStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	scormID := resolveValue("scormid", current, config)
	if scormID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "scormid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_scorm_get_scorm_scos", map[string]string{"scormid": scormID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	scos, _ := result["scos"].([]any)
	return &sdk.StepResult{Output: map[string]any{"scos": scos}}, nil
}

type scormGetUserDataStep struct {
	name       string
	moduleName string
}

func newScormGetUserDataStep(name string, config map[string]any) (*scormGetUserDataStep, error) {
	return &scormGetUserDataStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormGetUserDataStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	scormID := resolveValue("scormid", current, config)
	attempt := resolveValue("attempt", current, config)
	if scormID == "" || attempt == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "scormid and attempt are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_scorm_get_scorm_user_data", map[string]string{
		"scormid": scormID, "attempt": attempt,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	data, _ := result["data"].([]any)
	return &sdk.StepResult{Output: map[string]any{"data": data}}, nil
}

type scormInsertTracksStep struct {
	name       string
	moduleName string
}

func newScormInsertTracksStep(name string, config map[string]any) (*scormInsertTracksStep, error) {
	return &scormInsertTracksStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormInsertTracksStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	scoid := resolveValue("scoid", current, config)
	attempt := resolveValue("attempt", current, config)
	element := resolveValue("element", current, config)
	value := resolveValue("value", current, config)
	if scoid == "" || attempt == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "scoid and attempt are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_scorm_insert_scorm_tracks", map[string]string{
		"scoid": scoid, "attempt": attempt,
		"tracks[0][element]": element, "tracks[0][value]": value,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type scormLaunchScoStep struct {
	name       string
	moduleName string
}

func newScormLaunchScoStep(name string, config map[string]any) (*scormLaunchScoStep, error) {
	return &scormLaunchScoStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *scormLaunchScoStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	scormID := resolveValue("scormid", current, config)
	if scormID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "scormid is required"}}, nil
	}
	params := map[string]string{"scormid": scormID}
	if scoid := resolveValue("scoid", current, config); scoid != "" {
		params["scoid"] = scoid
	}
	result, err := client.callToMap(ctx, "mod_scorm_launch_sco", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}
