package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type tagGetTagsStep struct {
	name       string
	moduleName string
}

func newTagGetTagsStep(name string, config map[string]any) (*tagGetTagsStep, error) {
	return &tagGetTagsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tagGetTagsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if tagID := resolveValue("tagid", current, config); tagID != "" {
		params["tags[0][id]"] = tagID
	}
	if tagName := resolveValue("tagname", current, config); tagName != "" {
		params["tags[0][rawname]"] = tagName
	}
	result, err := client.callToMap(ctx, "core_tag_get_tags", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	tags, _ := result["tags"].([]any)
	return &sdk.StepResult{Output: map[string]any{"tags": tags, "warnings": result["warnings"]}}, nil
}

type tagUpdateStep struct {
	name       string
	moduleName string
}

func newTagUpdateStep(name string, config map[string]any) (*tagUpdateStep, error) {
	return &tagUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *tagUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	tagID := resolveValue("tagid", current, config)
	if tagID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "tagid is required"}}, nil
	}
	params := map[string]string{"tags[0][id]": tagID}
	if rawname := resolveValue("rawname", current, config); rawname != "" {
		params["tags[0][rawname]"] = rawname
	}
	if desc := resolveValue("description", current, config); desc != "" {
		params["tags[0][description]"] = desc
	}
	result, err := client.call(ctx, "core_tag_update_tags", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	if b, ok := result.(bool); ok {
		return &sdk.StepResult{Output: map[string]any{"updated": b}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"updated": true, "tagid": tagID}}, nil
}
