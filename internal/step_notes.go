package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type noteCreateStep struct {
	name       string
	moduleName string
}

func newNoteCreateStep(name string, config map[string]any) (*noteCreateStep, error) {
	return &noteCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *noteCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	text := resolveValue("text", current, config)
	publishState := resolveValue("publishstate", current, config)
	if userID == "" || text == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid and text are required"}}, nil
	}
	if publishState == "" {
		publishState = "personal"
	}
	params := map[string]string{
		"notes[0][userid]":       userID,
		"notes[0][publishstate]": publishState,
		"notes[0][courseid]":     resolveValue("courseid", current, config),
		"notes[0][text]":         text,
		"notes[0][format]":       "1",
	}
	result, err := client.callToSlice(ctx, "core_notes_create_notes", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var note map[string]any
	if len(result) > 0 {
		note, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"note": note, "notes": result}}, nil
}

type noteGetStep struct {
	name       string
	moduleName string
}

func newNoteGetStep(name string, config map[string]any) (*noteGetStep, error) {
	return &noteGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *noteGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	courseID := resolveValue("courseid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	params := map[string]string{
		"notes[0][userid]":   userID,
		"notes[0][courseid]": courseID,
	}
	result, err := client.callToMap(ctx, "core_notes_get_course_notes", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type noteDeleteStep struct {
	name       string
	moduleName string
}

func newNoteDeleteStep(name string, config map[string]any) (*noteDeleteStep, error) {
	return &noteDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *noteDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	noteID := resolveValue("noteid", current, config)
	if noteID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "noteid is required"}}, nil
	}
	_, err := client.call(ctx, "core_notes_delete_notes", map[string]string{"notes[0]": noteID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "noteid": noteID}}, nil
}
