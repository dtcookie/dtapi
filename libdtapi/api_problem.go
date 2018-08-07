package dtapi

import (
	"errors"

	"github.com/antihax/optional"
	dtapi "github.com/dtcookie/dtapi/libdtapienv"
)

// ProblemAPI TODO: documentation
type ProblemAPI envService

// GetEventByID TODO: documentation
func (api ProblemAPI) GetEventByID(eventID string) (dtapi.EventRestEntry, error) {
	result, _, err := api.client.EventApi.GetEventById(nil, eventID)
	return result, err
}

// DeleteComment TODO: documentation
func (api ProblemAPI) DeleteComment(problemID int64, commentID int64) error {
	_, err := api.client.ProblemApi.DeleteComment(nil, problemID, commentID)
	return err
}

// GetComments TODO: documentation
func (api ProblemAPI) GetComments(problemID int64) ([]dtapi.ProblemComment, error) {
	result, _, err := api.client.ProblemApi.GetComment(nil, problemID)
	return result.Comments, err
}

// GetDetails TODO: documentation
func (api ProblemAPI) GetDetails(problemID int64) (dtapi.Problem, error) {
	result, _, err := api.client.ProblemApi.GetDetails(nil, problemID)
	return result.Result, err
}

// GetFeed TODO: documentation
func (api ProblemAPI) GetFeed(opts *dtapi.GetFeedOpts) (dtapi.ProblemFeedQueryResult, error) {
	result, _, err := api.client.ProblemApi.GetFeed(nil, opts)
	return result.Result, err
}

// GetProblemStatus TODO: documentation
func (api ProblemAPI) GetProblemStatus() (dtapi.GlobalProblemStatus, error) {
	result, _, err := api.client.ProblemApi.GetProblemStatus(nil)
	return result.Result, err
}

// PushComment TODO: documentation
func (api ProblemAPI) PushComment(problemID int64, comment dtapi.ProblemComment) (dtapi.ProblemComment, error) {
	result, _, err := api.client.ProblemApi.PushComment(nil, problemID, &dtapi.PushCommentOpts{
		ProblemComment: optional.NewInterface(comment),
	})
	return result, err
}

// UpdateCommentCall TODO: documentation
type UpdateCommentCall struct {
	api       ProblemAPI
	ProblemID int64
	CommentID int64
	Comment   dtapi.ProblemComment
}

func assertNotEmpty(fieldName string, fieldValue string) error {
	if fieldValue == "" {
		return errors.New("field '" + fieldName + "' is not optional")
	}
	return nil
}

// Execute TODO: documentation
func (call UpdateCommentCall) Execute() (dtapi.ProblemComment, error) {
	if err := assertNotEmpty("User", call.Comment.User); err != nil {
		return call.Comment, err
	}

	return call.api.UpdateComment(call.ProblemID, call.CommentID, &dtapi.UpdateCommentOpts{
		ProblemComment: optional.NewInterface(call.Comment),
	})
}

// UpdateCommentREST TODO: documentation
func (api ProblemAPI) UpdateCommentREST(problemID int64, commentID int64, opts *dtapi.UpdateCommentOpts) UpdateCommentCall {
	return UpdateCommentCall{}
}

// UpdateComment TODO: documentation
func (api ProblemAPI) UpdateComment(problemID int64, commentID int64, opts *dtapi.UpdateCommentOpts) (dtapi.ProblemComment, error) {
	result, _, err := api.client.ProblemApi.UpdateComment(nil, problemID, commentID, opts)
	return result, err
}
