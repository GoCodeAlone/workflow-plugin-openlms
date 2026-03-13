package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type roleAssignStep struct {
	name       string
	moduleName string
}

func newRoleAssignStep(name string, config map[string]any) (*roleAssignStep, error) {
	return &roleAssignStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleAssignStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	roleID := resolveValue("roleid", current, config)
	userID := resolveValue("userid", current, config)
	contextID := resolveValue("contextid", current, config)
	if roleID == "" || userID == "" || contextID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "roleid, userid, and contextid are required"}}, nil
	}
	params := map[string]string{
		"assignments[0][roleid]":    roleID,
		"assignments[0][userid]":    userID,
		"assignments[0][contextid]": contextID,
	}
	if contextlevel := resolveValue("contextlevel", current, config); contextlevel != "" {
		params["assignments[0][contextlevel]"] = contextlevel
		params["assignments[0][instanceid]"] = resolveValue("instanceid", current, config)
	}
	_, err := client.call(ctx, "core_role_assign_roles", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"assigned": true, "roleid": roleID, "userid": userID}}, nil
}

type roleUnassignStep struct {
	name       string
	moduleName string
}

func newRoleUnassignStep(name string, config map[string]any) (*roleUnassignStep, error) {
	return &roleUnassignStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *roleUnassignStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	roleID := resolveValue("roleid", current, config)
	userID := resolveValue("userid", current, config)
	contextID := resolveValue("contextid", current, config)
	if roleID == "" || userID == "" || contextID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "roleid, userid, and contextid are required"}}, nil
	}
	params := map[string]string{
		"unassignments[0][roleid]":    roleID,
		"unassignments[0][userid]":    userID,
		"unassignments[0][contextid]": contextID,
	}
	_, err := client.call(ctx, "core_role_unassign_roles", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"unassigned": true, "roleid": roleID, "userid": userID}}, nil
}
