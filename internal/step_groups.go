package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type groupCreateStep struct {
	name       string
	moduleName string
}

func newGroupCreateStep(name string, config map[string]any) (*groupCreateStep, error) {
	return &groupCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseid := resolveValue("courseid", current, config)
	groupname := resolveValue("name", current, config)
	if courseid == "" || groupname == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid and name are required"}}, nil
	}
	params := map[string]string{
		"groups[0][courseid]": courseid,
		"groups[0][name]":     groupname,
	}
	if desc := resolveValue("description", current, config); desc != "" {
		params["groups[0][description]"] = desc
	}
	result, err := client.callToSlice(ctx, "core_group_create_groups", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var group map[string]any
	if len(result) > 0 {
		group, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"group": group, "groups": result}}, nil
}

type groupDeleteStep struct {
	name       string
	moduleName string
}

func newGroupDeleteStep(name string, config map[string]any) (*groupDeleteStep, error) {
	return &groupDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	groupID := resolveValue("groupid", current, config)
	if groupID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "groupid is required"}}, nil
	}
	_, err := client.call(ctx, "core_group_delete_groups", map[string]string{"groupids[0]": groupID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "groupid": groupID}}, nil
}

type groupGetCourseGroupsStep struct {
	name       string
	moduleName string
}

func newGroupGetCourseGroupsStep(name string, config map[string]any) (*groupGetCourseGroupsStep, error) {
	return &groupGetCourseGroupsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupGetCourseGroupsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	courseID := resolveValue("courseid", current, config)
	if courseID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "courseid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_group_get_course_groups", map[string]string{"courseid": courseID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"groups": result, "count": len(result)}}, nil
}

type groupGetMembersStep struct {
	name       string
	moduleName string
}

func newGroupGetMembersStep(name string, config map[string]any) (*groupGetMembersStep, error) {
	return &groupGetMembersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupGetMembersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	groupID := resolveValue("groupid", current, config)
	if groupID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "groupid is required"}}, nil
	}
	result, err := client.callToSlice(ctx, "core_group_get_group_members", map[string]string{"groupids[0]": groupID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"groups": result}}, nil
}

type groupAddMembersStep struct {
	name       string
	moduleName string
}

func newGroupAddMembersStep(name string, config map[string]any) (*groupAddMembersStep, error) {
	return &groupAddMembersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupAddMembersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	groupID := resolveValue("groupid", current, config)
	userID := resolveValue("userid", current, config)
	if groupID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "groupid and userid are required"}}, nil
	}
	params := map[string]string{
		"members[0][groupid]": groupID,
		"members[0][userid]":  userID,
	}
	_, err := client.call(ctx, "core_group_add_group_members", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"added": true, "groupid": groupID, "userid": userID}}, nil
}

type groupDeleteMembersStep struct {
	name       string
	moduleName string
}

func newGroupDeleteMembersStep(name string, config map[string]any) (*groupDeleteMembersStep, error) {
	return &groupDeleteMembersStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *groupDeleteMembersStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	groupID := resolveValue("groupid", current, config)
	userID := resolveValue("userid", current, config)
	if groupID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "groupid and userid are required"}}, nil
	}
	params := map[string]string{
		"members[0][groupid]": groupID,
		"members[0][userid]":  userID,
	}
	_, err := client.call(ctx, "core_group_delete_group_members", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "groupid": groupID, "userid": userID}}, nil
}
