package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type fileGetFilesStep struct {
	name       string
	moduleName string
}

func newFileGetFilesStep(name string, config map[string]any) (*fileGetFilesStep, error) {
	return &fileGetFilesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fileGetFilesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	contextID := resolveValue("contextid", current, config)
	component := resolveValue("component", current, config)
	filearea := resolveValue("filearea", current, config)
	itemID := resolveValue("itemid", current, config)
	if contextID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "contextid is required"}}, nil
	}
	params := map[string]string{
		"contextid": contextID,
		"component": component,
		"filearea":  filearea,
		"itemid":    itemID,
		"filepath":  resolveValue("filepath", current, config),
		"filename":  resolveValue("filename", current, config),
	}
	result, err := client.callToMap(ctx, "core_files_get_files", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	files, _ := result["files"].([]any)
	return &sdk.StepResult{Output: map[string]any{"files": files, "warnings": result["warnings"]}}, nil
}

type fileUploadStep struct {
	name       string
	moduleName string
}

func newFileUploadStep(name string, config map[string]any) (*fileUploadStep, error) {
	return &fileUploadStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *fileUploadStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	filename := resolveValue("filename", current, config)
	filecontent := resolveValue("filecontent", current, config)
	if filename == "" || filecontent == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "filename and filecontent are required"}}, nil
	}
	params := map[string]string{
		"files[0][filename]":       filename,
		"files[0][filecontent]":    filecontent,
		"files[0][contextid]":      resolveValue("contextid", current, config),
		"files[0][component]":      resolveValue("component", current, config),
		"files[0][filearea]":       resolveValue("filearea", current, config),
		"files[0][itemid]":         resolveValue("itemid", current, config),
	}
	result, err := client.callToSlice(ctx, "core_files_upload", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var file map[string]any
	if len(result) > 0 {
		file, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"file": file, "files": result}}, nil
}
