package internal

import (
	"context"
	"fmt"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

// userCreateStep implements step.openlms_user_create → core_user_create_users
type userCreateStep struct {
	name       string
	moduleName string
}

func newUserCreateStep(name string, config map[string]any) (*userCreateStep, error) {
	return &userCreateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userCreateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	username := resolveValue("username", current, config)
	password := resolveValue("password", current, config)
	firstname := resolveValue("firstname", current, config)
	lastname := resolveValue("lastname", current, config)
	email := resolveValue("email", current, config)
	if username == "" || email == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "username and email are required"}}, nil
	}
	params := map[string]string{
		"users[0][username]":  username,
		"users[0][password]":  password,
		"users[0][firstname]": firstname,
		"users[0][lastname]":  lastname,
		"users[0][email]":     email,
	}
	if auth := resolveValue("auth", current, config); auth != "" {
		params["users[0][auth]"] = auth
	}
	result, err := client.callToSlice(ctx, "core_user_create_users", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var user map[string]any
	if len(result) > 0 {
		if m, ok := result[0].(map[string]any); ok {
			user = m
		}
	}
	return &sdk.StepResult{Output: map[string]any{"user": user, "users": result}}, nil
}

// userUpdateStep implements step.openlms_user_update → core_user_update_users
type userUpdateStep struct {
	name       string
	moduleName string
}

func newUserUpdateStep(name string, config map[string]any) (*userUpdateStep, error) {
	return &userUpdateStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userUpdateStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	params := map[string]string{"users[0][id]": userID}
	for _, field := range []string{"username", "firstname", "lastname", "email", "password", "auth"} {
		if v := resolveValue(field, current, config); v != "" {
			params[fmt.Sprintf("users[0][%s]", field)] = v
		}
	}
	_, err := client.call(ctx, "core_user_update_users", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"updated": true, "userid": userID}}, nil
}

// userDeleteStep implements step.openlms_user_delete → core_user_delete_users
type userDeleteStep struct {
	name       string
	moduleName string
}

func newUserDeleteStep(name string, config map[string]any) (*userDeleteStep, error) {
	return &userDeleteStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userDeleteStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	_, err := client.call(ctx, "core_user_delete_users", map[string]string{"userids[0]": userID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "userid": userID}}, nil
}

// userGetStep implements step.openlms_user_get → core_user_get_users
type userGetStep struct {
	name       string
	moduleName string
}

func newUserGetStep(name string, config map[string]any) (*userGetStep, error) {
	return &userGetStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userGetStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	// Search by criteria
	key := resolveValue("key", current, config)
	value := resolveValue("value", current, config)
	if key == "" {
		key = "id"
	}
	if value == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "value is required"}}, nil
	}
	params := map[string]string{
		"criteria[0][key]":   key,
		"criteria[0][value]": value,
	}
	result, err := client.callToMap(ctx, "core_user_get_users", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	users, _ := result["users"].([]any)
	var user map[string]any
	if len(users) > 0 {
		user, _ = users[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"user": user, "users": users, "total": result["totalrecords"]}}, nil
}

// userGetByFieldStep implements step.openlms_user_get_by_field → core_user_get_users_by_field
type userGetByFieldStep struct {
	name       string
	moduleName string
}

func newUserGetByFieldStep(name string, config map[string]any) (*userGetByFieldStep, error) {
	return &userGetByFieldStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userGetByFieldStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	field := resolveValue("field", current, config)
	value := resolveValue("value", current, config)
	if field == "" || value == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "field and value are required"}}, nil
	}
	params := map[string]string{
		"field":      field,
		"values[0]":  value,
	}
	result, err := client.callToSlice(ctx, "core_user_get_users_by_field", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var user map[string]any
	if len(result) > 0 {
		user, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"user": user, "users": result}}, nil
}

// userSearchStep implements step.openlms_user_search → core_user_search_identity
type userSearchStep struct {
	name       string
	moduleName string
}

func newUserSearchStep(name string, config map[string]any) (*userSearchStep, error) {
	return &userSearchStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *userSearchStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	query := resolveValue("query", current, config)
	if query == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "query is required"}}, nil
	}
	params := map[string]string{"query": query}
	result, err := client.callToMap(ctx, "core_user_search_identity", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	users, _ := result["list"].([]any)
	return &sdk.StepResult{Output: map[string]any{"users": users, "total": result["maxusersperpage"]}}, nil
}
