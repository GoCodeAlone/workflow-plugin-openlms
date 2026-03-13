package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type messageSendStep struct {
	name       string
	moduleName string
}

func newMessageSendStep(name string, config map[string]any) (*messageSendStep, error) {
	return &messageSendStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageSendStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	touserid := resolveValue("touserid", current, config)
	text := resolveValue("text", current, config)
	if touserid == "" || text == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "touserid and text are required"}}, nil
	}
	params := map[string]string{
		"messages[0][touserid]":   touserid,
		"messages[0][text]":       text,
		"messages[0][textformat]": "0",
	}
	result, err := client.callToSlice(ctx, "core_message_send_instant_messages", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	var msg map[string]any
	if len(result) > 0 {
		msg, _ = result[0].(map[string]any)
	}
	return &sdk.StepResult{Output: map[string]any{"message": msg, "messages": result}}, nil
}

type messageGetMessagesStep struct {
	name       string
	moduleName string
}

func newMessageGetMessagesStep(name string, config map[string]any) (*messageGetMessagesStep, error) {
	return &messageGetMessagesStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageGetMessagesStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	params := map[string]string{"useridto": userID, "type": "conversations"}
	if msgType := resolveValue("type", current, config); msgType != "" {
		params["type"] = msgType
	}
	result, err := client.callToMap(ctx, "core_message_get_messages", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	messages, _ := result["messages"].([]any)
	return &sdk.StepResult{Output: map[string]any{"messages": messages, "count": result["count"]}}, nil
}

type messageGetConversationsStep struct {
	name       string
	moduleName string
}

func newMessageGetConversationsStep(name string, config map[string]any) (*messageGetConversationsStep, error) {
	return &messageGetConversationsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageGetConversationsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_message_get_conversations", map[string]string{"userid": userID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	conversations, _ := result["conversations"].([]any)
	return &sdk.StepResult{Output: map[string]any{"conversations": conversations}}, nil
}

type messageGetUnreadCountStep struct {
	name       string
	moduleName string
}

func newMessageGetUnreadCountStep(name string, config map[string]any) (*messageGetUnreadCountStep, error) {
	return &messageGetUnreadCountStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageGetUnreadCountStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	if userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "core_message_get_unread_conversations_count", map[string]string{"useridto": userID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type messageMarkReadStep struct {
	name       string
	moduleName string
}

func newMessageMarkReadStep(name string, config map[string]any) (*messageMarkReadStep, error) {
	return &messageMarkReadStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageMarkReadStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	messageID := resolveValue("messageid", current, config)
	userID := resolveValue("userid", current, config)
	if messageID == "" || userID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "messageid and userid are required"}}, nil
	}
	result, err := client.call(ctx, "core_message_mark_message_read", map[string]string{"messageid": messageID, "timeread": "0"})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	if m, ok := result.(map[string]any); ok {
		return &sdk.StepResult{Output: m}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"marked": true}}, nil
}

type messageBlockUserStep struct {
	name       string
	moduleName string
}

func newMessageBlockUserStep(name string, config map[string]any) (*messageBlockUserStep, error) {
	return &messageBlockUserStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageBlockUserStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	blockedUserID := resolveValue("blockeduserid", current, config)
	if userID == "" || blockedUserID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid and blockeduserid are required"}}, nil
	}
	_, err := client.call(ctx, "core_message_block_user", map[string]string{"userid": userID, "blockeduserid": blockedUserID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"blocked": true}}, nil
}

type messageUnblockUserStep struct {
	name       string
	moduleName string
}

func newMessageUnblockUserStep(name string, config map[string]any) (*messageUnblockUserStep, error) {
	return &messageUnblockUserStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *messageUnblockUserStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	userID := resolveValue("userid", current, config)
	unblockedUserID := resolveValue("unblockeduserid", current, config)
	if userID == "" || unblockedUserID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "userid and unblockeduserid are required"}}, nil
	}
	_, err := client.call(ctx, "core_message_unblock_user", map[string]string{"userid": userID, "unblockeduserid": unblockedUserID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"unblocked": true}}, nil
}
