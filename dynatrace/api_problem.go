package dynatrace

import (
	"errors"

	"github.com/antihax/optional"
	"github.com/dtcookie/dtapi/environment"
)

// ProblemAPI TODO: documentation
type ProblemAPI service

// GetEventByID TODO: documentation
func (api ProblemAPI) GetEventByID(eventID string) (environment.EventRestEntry, error) {
	result, _, err := api.client.EventApi.GetEventById(nil, eventID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// DeleteComment TODO: documentation
func (api ProblemAPI) DeleteComment(problemID int64, commentID int64) error {
	_, err := api.client.ProblemApi.DeleteComment(nil, problemID, commentID)
	if err != nil {
		return err
	}
	return nil
}

// GetComment TODO: documentation
func (api ProblemAPI) GetComment(problemID int64) (environment.ProblemCommentList, error) {
	result, _, err := api.client.ProblemApi.GetComment(nil, problemID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetDetails TODO: documentation
func (api ProblemAPI) GetDetails(problemID int64) (environment.ProblemDetailsResultWrapper, error) {
	result, _, err := api.client.ProblemApi.GetDetails(nil, problemID)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetFeed TODO: documentation
func (api ProblemAPI) GetFeed(opts *environment.GetFeedOpts) (environment.ProblemFeedResultWrapper, error) {
	result, _, err := api.client.ProblemApi.GetFeed(nil, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetProblemStatus TODO: documentation
func (api ProblemAPI) GetProblemStatus() (environment.ProblemStatusResultWrapper, error) {
	result, _, err := api.client.ProblemApi.GetProblemStatus(nil)
	if err != nil {
		return result, err
	}
	return result, nil
}

// PushComment TODO: documentation
func (api ProblemAPI) PushComment(problemID int64, opts *environment.PushCommentOpts) (environment.ProblemComment, error) {
	result, _, err := api.client.ProblemApi.PushComment(nil, problemID, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateCommentCall TODO: documentation
type UpdateCommentCall struct {
	api       ProblemAPI
	ProblemID int64
	CommentID int64
	Comment   environment.ProblemComment
}

func assertNotEmpty(fieldName string, fieldValue string) error {
	if fieldValue == "" {
		return errors.New("field '" + fieldName + "' is not optional")
	}
	return nil
}

// Execute TODO: documentation
func (call UpdateCommentCall) Execute() (environment.ProblemComment, error) {
	if err := assertNotEmpty("User", call.Comment.User); err != nil {
		return call.Comment, err
	}

	return call.api.UpdateComment(call.ProblemID, call.CommentID, &environment.UpdateCommentOpts{
		ProblemComment: optional.NewInterface(call.Comment),
	})
}

// UpdateCommentREST TODO: documentation
func (api ProblemAPI) UpdateCommentREST(problemID int64, commentID int64, opts *environment.UpdateCommentOpts) UpdateCommentCall {
	return UpdateCommentCall{}
}

// UpdateComment TODO: documentation
func (api ProblemAPI) UpdateComment(problemID int64, commentID int64, opts *environment.UpdateCommentOpts) (environment.ProblemComment, error) {
	result, _, err := api.client.ProblemApi.UpdateComment(nil, problemID, commentID, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}
