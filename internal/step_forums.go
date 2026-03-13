package internal

import (
	"context"

	sdk "github.com/GoCodeAlone/workflow/plugin/external/sdk"
)

type forumGetByCourseStep struct {
	name       string
	moduleName string
}

func newForumGetByCourseStep(name string, config map[string]any) (*forumGetByCourseStep, error) {
	return &forumGetByCourseStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumGetByCourseStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	params := map[string]string{}
	if courseID := resolveValue("courseid", current, config); courseID != "" {
		params["courseids[0]"] = courseID
	}
	result, err := client.callToSlice(ctx, "mod_forum_get_forums_by_courses", params)
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"forums": result, "count": len(result)}}, nil
}

type forumGetDiscussionsStep struct {
	name       string
	moduleName string
}

func newForumGetDiscussionsStep(name string, config map[string]any) (*forumGetDiscussionsStep, error) {
	return &forumGetDiscussionsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumGetDiscussionsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	forumID := resolveValue("forumid", current, config)
	if forumID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "forumid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_forum_get_forum_discussions", map[string]string{"forumid": forumID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	discussions, _ := result["discussions"].([]any)
	return &sdk.StepResult{Output: map[string]any{"discussions": discussions, "count": len(discussions)}}, nil
}

type forumGetPostsStep struct {
	name       string
	moduleName string
}

func newForumGetPostsStep(name string, config map[string]any) (*forumGetPostsStep, error) {
	return &forumGetPostsStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumGetPostsStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	discussionID := resolveValue("discussionid", current, config)
	if discussionID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "discussionid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_forum_get_discussion_posts", map[string]string{"discussionid": discussionID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	posts, _ := result["posts"].([]any)
	return &sdk.StepResult{Output: map[string]any{"posts": posts, "count": len(posts)}}, nil
}

type forumAddDiscussionStep struct {
	name       string
	moduleName string
}

func newForumAddDiscussionStep(name string, config map[string]any) (*forumAddDiscussionStep, error) {
	return &forumAddDiscussionStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumAddDiscussionStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	forumID := resolveValue("forumid", current, config)
	subject := resolveValue("subject", current, config)
	message := resolveValue("message", current, config)
	if forumID == "" || subject == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "forumid and subject are required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_forum_add_discussion", map[string]string{
		"forumid": forumID,
		"subject": subject,
		"message": message,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type forumAddPostStep struct {
	name       string
	moduleName string
}

func newForumAddPostStep(name string, config map[string]any) (*forumAddPostStep, error) {
	return &forumAddPostStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumAddPostStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	postID := resolveValue("postid", current, config)
	subject := resolveValue("subject", current, config)
	message := resolveValue("message", current, config)
	if postID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "postid is required"}}, nil
	}
	result, err := client.callToMap(ctx, "mod_forum_add_discussion_post", map[string]string{
		"postid":  postID,
		"subject": subject,
		"message": message,
	})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: result}, nil
}

type forumDeletePostStep struct {
	name       string
	moduleName string
}

func newForumDeletePostStep(name string, config map[string]any) (*forumDeletePostStep, error) {
	return &forumDeletePostStep{name: name, moduleName: getModuleName(config)}, nil
}

func (s *forumDeletePostStep) Execute(ctx context.Context, _ map[string]any, _ map[string]map[string]any, current map[string]any, _ map[string]any, config map[string]any) (*sdk.StepResult, error) {
	client, ok := GetClient(s.moduleName)
	if !ok {
		return &sdk.StepResult{Output: map[string]any{"error": "openlms client not found: " + s.moduleName}}, nil
	}
	postID := resolveValue("postid", current, config)
	if postID == "" {
		return &sdk.StepResult{Output: map[string]any{"error": "postid is required"}}, nil
	}
	_, err := client.call(ctx, "mod_forum_delete_post", map[string]string{"postid": postID})
	if err != nil {
		return &sdk.StepResult{Output: map[string]any{"error": err.Error()}}, nil
	}
	return &sdk.StepResult{Output: map[string]any{"deleted": true, "postid": postID}}, nil
}
