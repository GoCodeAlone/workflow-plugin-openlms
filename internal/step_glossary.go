package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type glossaryGetByCourseStep struct {
	name       string
	moduleName string
}

func newGlossaryGetByCourseStep(name string, config map[string]any) (*glossaryGetByCourseStep, error) {
	return &glossaryGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *glossaryGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToMap(ctx, "mod_glossary_get_glossaries_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	glossaries, _ := result["glossaries"].([]any)
	return &sdk.StepResult{Output: map[string]any{"glossaries": glossaries}}, nil
}

type glossaryGetEntriesStep struct {
	name       string
	moduleName string
}

func newGlossaryGetEntriesStep(name string, config map[string]any) (*glossaryGetEntriesStep, error) {
	return &glossaryGetEntriesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *glossaryGetEntriesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	glossaryID := resolveValue("id", current, config)
	if glossaryID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "id is required"}}, nil
	}
	params := map[string]string{
		"id":     glossaryID,
		"letter": resolveValue("letter", current, config),
	}
	result, err := client.callToMap(ctx, "mod_glossary_get_entries_by_letter", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	entries, _ := result["entries"].([]any)
	return &sdk.StepResult{Output: map[string]any{"entries": entries, "count": result["count"]}}, nil
}

type glossaryAddEntryStep struct {
	name       string
	moduleName string
}

func newGlossaryAddEntryStep(name string, config map[string]any) (*glossaryAddEntryStep, error) {
	return &glossaryAddEntryStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *glossaryAddEntryStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	glossaryID := resolveValue("glossaryid", current, config)
	concept := resolveValue("concept", current, config)
	definition := resolveValue("definition", current, config)
	if glossaryID == "" || concept == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "glossaryid and concept are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_glossary_add_entry", map[string]string{
		"glossaryid": glossaryID, "concept": concept, "definition": definition, "definitionformat": "1",
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type glossaryDeleteEntryStep struct {
	name       string
	moduleName string
}

func newGlossaryDeleteEntryStep(name string, config map[string]any) (*glossaryDeleteEntryStep, error) {
	return &glossaryDeleteEntryStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *glossaryDeleteEntryStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	entryID := resolveValue("entryid", current, config)
	if entryID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "entryid is required"}}, nil
	}
	_, err := client.call(ctx, "mod_glossary_delete_entry", map[string]string{"entryid": entryID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "entryid": entryID}}, nil
}
