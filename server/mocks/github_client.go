// Code generated by MockGen. DO NOT EDIT.
// Source: server/github_client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	github "github.com/google/go-github/v39/github"
)

// MockChecksService is a mock of ChecksService interface.
type MockChecksService struct {
	ctrl     *gomock.Controller
	recorder *MockChecksServiceMockRecorder
}

// MockChecksServiceMockRecorder is the mock recorder for MockChecksService.
type MockChecksServiceMockRecorder struct {
	mock *MockChecksService
}

// NewMockChecksService creates a new mock instance.
func NewMockChecksService(ctrl *gomock.Controller) *MockChecksService {
	mock := &MockChecksService{ctrl: ctrl}
	mock.recorder = &MockChecksServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChecksService) EXPECT() *MockChecksServiceMockRecorder {
	return m.recorder
}

// ListCheckRunsForRef mocks base method.
func (m *MockChecksService) ListCheckRunsForRef(ctx context.Context, owner, repo, ref string, opts *github.ListCheckRunsOptions) (*github.ListCheckRunsResults, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCheckRunsForRef", ctx, owner, repo, ref, opts)
	ret0, _ := ret[0].(*github.ListCheckRunsResults)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCheckRunsForRef indicates an expected call of ListCheckRunsForRef.
func (mr *MockChecksServiceMockRecorder) ListCheckRunsForRef(ctx, owner, repo, ref, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCheckRunsForRef", reflect.TypeOf((*MockChecksService)(nil).ListCheckRunsForRef), ctx, owner, repo, ref, opts)
}

// MockIssuesService is a mock of IssuesService interface.
type MockIssuesService struct {
	ctrl     *gomock.Controller
	recorder *MockIssuesServiceMockRecorder
}

// MockIssuesServiceMockRecorder is the mock recorder for MockIssuesService.
type MockIssuesServiceMockRecorder struct {
	mock *MockIssuesService
}

// NewMockIssuesService creates a new mock instance.
func NewMockIssuesService(ctrl *gomock.Controller) *MockIssuesService {
	mock := &MockIssuesService{ctrl: ctrl}
	mock.recorder = &MockIssuesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIssuesService) EXPECT() *MockIssuesServiceMockRecorder {
	return m.recorder
}

// AddAssignees mocks base method.
func (m *MockIssuesService) AddAssignees(ctx context.Context, owner, repo string, number int, assignees []string) (*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddAssignees", ctx, owner, repo, number, assignees)
	ret0, _ := ret[0].(*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddAssignees indicates an expected call of AddAssignees.
func (mr *MockIssuesServiceMockRecorder) AddAssignees(ctx, owner, repo, number, assignees interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAssignees", reflect.TypeOf((*MockIssuesService)(nil).AddAssignees), ctx, owner, repo, number, assignees)
}

// AddLabelsToIssue mocks base method.
func (m *MockIssuesService) AddLabelsToIssue(ctx context.Context, owner, repo string, number int, labels []string) ([]*github.Label, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLabelsToIssue", ctx, owner, repo, number, labels)
	ret0, _ := ret[0].([]*github.Label)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// AddLabelsToIssue indicates an expected call of AddLabelsToIssue.
func (mr *MockIssuesServiceMockRecorder) AddLabelsToIssue(ctx, owner, repo, number, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLabelsToIssue", reflect.TypeOf((*MockIssuesService)(nil).AddLabelsToIssue), ctx, owner, repo, number, labels)
}

// CreateComment mocks base method.
func (m *MockIssuesService) CreateComment(ctx context.Context, owner, repo string, number int, comment *github.IssueComment) (*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, owner, repo, number, comment)
	ret0, _ := ret[0].(*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockIssuesServiceMockRecorder) CreateComment(ctx, owner, repo, number, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockIssuesService)(nil).CreateComment), ctx, owner, repo, number, comment)
}

// DeleteComment mocks base method.
func (m *MockIssuesService) DeleteComment(ctx context.Context, owner, repo string, commentID int64) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, owner, repo, commentID)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockIssuesServiceMockRecorder) DeleteComment(ctx, owner, repo, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockIssuesService)(nil).DeleteComment), ctx, owner, repo, commentID)
}

// Edit mocks base method.
func (m *MockIssuesService) Edit(ctx context.Context, owner, repo string, number int, issue *github.IssueRequest) (*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Edit", ctx, owner, repo, number, issue)
	ret0, _ := ret[0].(*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Edit indicates an expected call of Edit.
func (mr *MockIssuesServiceMockRecorder) Edit(ctx, owner, repo, number, issue interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Edit", reflect.TypeOf((*MockIssuesService)(nil).Edit), ctx, owner, repo, number, issue)
}

// Get mocks base method.
func (m *MockIssuesService) Get(ctx context.Context, owner, repo string, number int) (*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, owner, repo, number)
	ret0, _ := ret[0].(*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockIssuesServiceMockRecorder) Get(ctx, owner, repo, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIssuesService)(nil).Get), ctx, owner, repo, number)
}

// ListByRepo mocks base method.
func (m *MockIssuesService) ListByRepo(ctx context.Context, owner, repo string, opts *github.IssueListByRepoOptions) ([]*github.Issue, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByRepo", ctx, owner, repo, opts)
	ret0, _ := ret[0].([]*github.Issue)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListByRepo indicates an expected call of ListByRepo.
func (mr *MockIssuesServiceMockRecorder) ListByRepo(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByRepo", reflect.TypeOf((*MockIssuesService)(nil).ListByRepo), ctx, owner, repo, opts)
}

// ListComments mocks base method.
func (m *MockIssuesService) ListComments(ctx context.Context, owner, repo string, number int, opts *github.IssueListCommentsOptions) ([]*github.IssueComment, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListComments", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].([]*github.IssueComment)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListComments indicates an expected call of ListComments.
func (mr *MockIssuesServiceMockRecorder) ListComments(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListComments", reflect.TypeOf((*MockIssuesService)(nil).ListComments), ctx, owner, repo, number, opts)
}

// ListLabelsByIssue mocks base method.
func (m *MockIssuesService) ListLabelsByIssue(ctx context.Context, owner, repo string, number int, opt *github.ListOptions) ([]*github.Label, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLabelsByIssue", ctx, owner, repo, number, opt)
	ret0, _ := ret[0].([]*github.Label)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListLabelsByIssue indicates an expected call of ListLabelsByIssue.
func (mr *MockIssuesServiceMockRecorder) ListLabelsByIssue(ctx, owner, repo, number, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLabelsByIssue", reflect.TypeOf((*MockIssuesService)(nil).ListLabelsByIssue), ctx, owner, repo, number, opt)
}

// RemoveLabelForIssue mocks base method.
func (m *MockIssuesService) RemoveLabelForIssue(ctx context.Context, owner, repo string, number int, label string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLabelForIssue", ctx, owner, repo, number, label)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLabelForIssue indicates an expected call of RemoveLabelForIssue.
func (mr *MockIssuesServiceMockRecorder) RemoveLabelForIssue(ctx, owner, repo, number, label interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLabelForIssue", reflect.TypeOf((*MockIssuesService)(nil).RemoveLabelForIssue), ctx, owner, repo, number, label)
}

// MockGitService is a mock of GitService interface.
type MockGitService struct {
	ctrl     *gomock.Controller
	recorder *MockGitServiceMockRecorder
}

// MockGitServiceMockRecorder is the mock recorder for MockGitService.
type MockGitServiceMockRecorder struct {
	mock *MockGitService
}

// NewMockGitService creates a new mock instance.
func NewMockGitService(ctrl *gomock.Controller) *MockGitService {
	mock := &MockGitService{ctrl: ctrl}
	mock.recorder = &MockGitServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitService) EXPECT() *MockGitServiceMockRecorder {
	return m.recorder
}

// CreateRef mocks base method.
func (m *MockGitService) CreateRef(ctx context.Context, owner, repo string, ref *github.Reference) (*github.Reference, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRef", ctx, owner, repo, ref)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateRef indicates an expected call of CreateRef.
func (mr *MockGitServiceMockRecorder) CreateRef(ctx, owner, repo, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRef", reflect.TypeOf((*MockGitService)(nil).CreateRef), ctx, owner, repo, ref)
}

// DeleteRef mocks base method.
func (m *MockGitService) DeleteRef(ctx context.Context, owner, repo, ref string) (*github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRef", ctx, owner, repo, ref)
	ret0, _ := ret[0].(*github.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRef indicates an expected call of DeleteRef.
func (mr *MockGitServiceMockRecorder) DeleteRef(ctx, owner, repo, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRef", reflect.TypeOf((*MockGitService)(nil).DeleteRef), ctx, owner, repo, ref)
}

// GetRef mocks base method.
func (m *MockGitService) GetRef(ctx context.Context, owner, repo, ref string) (*github.Reference, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRef", ctx, owner, repo, ref)
	ret0, _ := ret[0].(*github.Reference)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetRef indicates an expected call of GetRef.
func (mr *MockGitServiceMockRecorder) GetRef(ctx, owner, repo, ref interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRef", reflect.TypeOf((*MockGitService)(nil).GetRef), ctx, owner, repo, ref)
}

// MockOrganizationsService is a mock of OrganizationsService interface.
type MockOrganizationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationsServiceMockRecorder
}

// MockOrganizationsServiceMockRecorder is the mock recorder for MockOrganizationsService.
type MockOrganizationsServiceMockRecorder struct {
	mock *MockOrganizationsService
}

// NewMockOrganizationsService creates a new mock instance.
func NewMockOrganizationsService(ctrl *gomock.Controller) *MockOrganizationsService {
	mock := &MockOrganizationsService{ctrl: ctrl}
	mock.recorder = &MockOrganizationsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrganizationsService) EXPECT() *MockOrganizationsServiceMockRecorder {
	return m.recorder
}

// GetOrgMembership mocks base method.
func (m *MockOrganizationsService) GetOrgMembership(ctx context.Context, user, org string) (*github.Membership, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgMembership", ctx, user, org)
	ret0, _ := ret[0].(*github.Membership)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetOrgMembership indicates an expected call of GetOrgMembership.
func (mr *MockOrganizationsServiceMockRecorder) GetOrgMembership(ctx, user, org interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgMembership", reflect.TypeOf((*MockOrganizationsService)(nil).GetOrgMembership), ctx, user, org)
}

// IsMember mocks base method.
func (m *MockOrganizationsService) IsMember(ctx context.Context, org, user string) (bool, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMember", ctx, org, user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// IsMember indicates an expected call of IsMember.
func (mr *MockOrganizationsServiceMockRecorder) IsMember(ctx, org, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMember", reflect.TypeOf((*MockOrganizationsService)(nil).IsMember), ctx, org, user)
}

// ListMembers mocks base method.
func (m *MockOrganizationsService) ListMembers(ctx context.Context, org string, opts *github.ListMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMembers", ctx, org, opts)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMembers indicates an expected call of ListMembers.
func (mr *MockOrganizationsServiceMockRecorder) ListMembers(ctx, org, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMembers", reflect.TypeOf((*MockOrganizationsService)(nil).ListMembers), ctx, org, opts)
}

// MockPullRequestsService is a mock of PullRequestsService interface.
type MockPullRequestsService struct {
	ctrl     *gomock.Controller
	recorder *MockPullRequestsServiceMockRecorder
}

// MockPullRequestsServiceMockRecorder is the mock recorder for MockPullRequestsService.
type MockPullRequestsServiceMockRecorder struct {
	mock *MockPullRequestsService
}

// NewMockPullRequestsService creates a new mock instance.
func NewMockPullRequestsService(ctrl *gomock.Controller) *MockPullRequestsService {
	mock := &MockPullRequestsService{ctrl: ctrl}
	mock.recorder = &MockPullRequestsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPullRequestsService) EXPECT() *MockPullRequestsServiceMockRecorder {
	return m.recorder
}

// CreateReview mocks base method.
func (m *MockPullRequestsService) CreateReview(ctx context.Context, owner, repo string, number int, review *github.PullRequestReviewRequest) (*github.PullRequestReview, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateReview", ctx, owner, repo, number, review)
	ret0, _ := ret[0].(*github.PullRequestReview)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateReview indicates an expected call of CreateReview.
func (mr *MockPullRequestsServiceMockRecorder) CreateReview(ctx, owner, repo, number, review interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateReview", reflect.TypeOf((*MockPullRequestsService)(nil).CreateReview), ctx, owner, repo, number, review)
}

// Get mocks base method.
func (m *MockPullRequestsService) Get(ctx context.Context, owner, repo string, number int) (*github.PullRequest, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, owner, repo, number)
	ret0, _ := ret[0].(*github.PullRequest)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockPullRequestsServiceMockRecorder) Get(ctx, owner, repo, number interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPullRequestsService)(nil).Get), ctx, owner, repo, number)
}

// List mocks base method.
func (m *MockPullRequestsService) List(ctx context.Context, owner, repo string, opts *github.PullRequestListOptions) ([]*github.PullRequest, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, owner, repo, opts)
	ret0, _ := ret[0].([]*github.PullRequest)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List.
func (mr *MockPullRequestsServiceMockRecorder) List(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPullRequestsService)(nil).List), ctx, owner, repo, opts)
}

// ListFiles mocks base method.
func (m *MockPullRequestsService) ListFiles(ctx context.Context, owner, repo string, number int, opts *github.ListOptions) ([]*github.CommitFile, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].([]*github.CommitFile)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockPullRequestsServiceMockRecorder) ListFiles(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockPullRequestsService)(nil).ListFiles), ctx, owner, repo, number, opts)
}

// ListReviewers mocks base method.
func (m *MockPullRequestsService) ListReviewers(ctx context.Context, owner, repo string, number int, opts *github.ListOptions) (*github.Reviewers, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReviewers", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].(*github.Reviewers)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListReviewers indicates an expected call of ListReviewers.
func (mr *MockPullRequestsServiceMockRecorder) ListReviewers(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReviewers", reflect.TypeOf((*MockPullRequestsService)(nil).ListReviewers), ctx, owner, repo, number, opts)
}

// ListReviews mocks base method.
func (m *MockPullRequestsService) ListReviews(ctx context.Context, owner, repo string, number int, opts *github.ListOptions) ([]*github.PullRequestReview, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListReviews", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].([]*github.PullRequestReview)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListReviews indicates an expected call of ListReviews.
func (mr *MockPullRequestsServiceMockRecorder) ListReviews(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListReviews", reflect.TypeOf((*MockPullRequestsService)(nil).ListReviews), ctx, owner, repo, number, opts)
}

// Merge mocks base method.
func (m *MockPullRequestsService) Merge(ctx context.Context, owner, repo string, number int, commitMessage string, options *github.PullRequestOptions) (*github.PullRequestMergeResult, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Merge", ctx, owner, repo, number, commitMessage, options)
	ret0, _ := ret[0].(*github.PullRequestMergeResult)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Merge indicates an expected call of Merge.
func (mr *MockPullRequestsServiceMockRecorder) Merge(ctx, owner, repo, number, commitMessage, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Merge", reflect.TypeOf((*MockPullRequestsService)(nil).Merge), ctx, owner, repo, number, commitMessage, options)
}

// RequestReviewers mocks base method.
func (m *MockPullRequestsService) RequestReviewers(ctx context.Context, owner, repo string, number int, reviewers github.ReviewersRequest) (*github.PullRequest, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestReviewers", ctx, owner, repo, number, reviewers)
	ret0, _ := ret[0].(*github.PullRequest)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RequestReviewers indicates an expected call of RequestReviewers.
func (mr *MockPullRequestsServiceMockRecorder) RequestReviewers(ctx, owner, repo, number, reviewers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestReviewers", reflect.TypeOf((*MockPullRequestsService)(nil).RequestReviewers), ctx, owner, repo, number, reviewers)
}

// UpdateBranch mocks base method.
func (m *MockPullRequestsService) UpdateBranch(ctx context.Context, owner, repo string, number int, opts *github.PullRequestBranchUpdateOptions) (*github.PullRequestBranchUpdateResponse, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBranch", ctx, owner, repo, number, opts)
	ret0, _ := ret[0].(*github.PullRequestBranchUpdateResponse)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// UpdateBranch indicates an expected call of UpdateBranch.
func (mr *MockPullRequestsServiceMockRecorder) UpdateBranch(ctx, owner, repo, number, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBranch", reflect.TypeOf((*MockPullRequestsService)(nil).UpdateBranch), ctx, owner, repo, number, opts)
}

// MockRepositoriesService is a mock of RepositoriesService interface.
type MockRepositoriesService struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoriesServiceMockRecorder
}

// MockRepositoriesServiceMockRecorder is the mock recorder for MockRepositoriesService.
type MockRepositoriesServiceMockRecorder struct {
	mock *MockRepositoriesService
}

// NewMockRepositoriesService creates a new mock instance.
func NewMockRepositoriesService(ctrl *gomock.Controller) *MockRepositoriesService {
	mock := &MockRepositoriesService{ctrl: ctrl}
	mock.recorder = &MockRepositoriesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoriesService) EXPECT() *MockRepositoriesServiceMockRecorder {
	return m.recorder
}

// CreateStatus mocks base method.
func (m *MockRepositoriesService) CreateStatus(ctx context.Context, owner, repo, ref string, status *github.RepoStatus) (*github.RepoStatus, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStatus", ctx, owner, repo, ref, status)
	ret0, _ := ret[0].(*github.RepoStatus)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateStatus indicates an expected call of CreateStatus.
func (mr *MockRepositoriesServiceMockRecorder) CreateStatus(ctx, owner, repo, ref, status interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStatus", reflect.TypeOf((*MockRepositoriesService)(nil).CreateStatus), ctx, owner, repo, ref, status)
}

// Get mocks base method.
func (m *MockRepositoriesService) Get(ctx context.Context, owner, repo string) (*github.Repository, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, owner, repo)
	ret0, _ := ret[0].(*github.Repository)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Get indicates an expected call of Get.
func (mr *MockRepositoriesServiceMockRecorder) Get(ctx, owner, repo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRepositoriesService)(nil).Get), ctx, owner, repo)
}

// GetBranch mocks base method.
func (m *MockRepositoriesService) GetBranch(ctx context.Context, owner, repo, branch string, followRedirects bool) (*github.Branch, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBranch", ctx, owner, repo, branch, followRedirects)
	ret0, _ := ret[0].(*github.Branch)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetBranch indicates an expected call of GetBranch.
func (mr *MockRepositoriesServiceMockRecorder) GetBranch(ctx, owner, repo, branch, followRedirects interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBranch", reflect.TypeOf((*MockRepositoriesService)(nil).GetBranch), ctx, owner, repo, branch, followRedirects)
}

// GetCombinedStatus mocks base method.
func (m *MockRepositoriesService) GetCombinedStatus(ctx context.Context, owner, repo, ref string, opts *github.ListOptions) (*github.CombinedStatus, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCombinedStatus", ctx, owner, repo, ref, opts)
	ret0, _ := ret[0].(*github.CombinedStatus)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCombinedStatus indicates an expected call of GetCombinedStatus.
func (mr *MockRepositoriesServiceMockRecorder) GetCombinedStatus(ctx, owner, repo, ref, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCombinedStatus", reflect.TypeOf((*MockRepositoriesService)(nil).GetCombinedStatus), ctx, owner, repo, ref, opts)
}

// ListStatuses mocks base method.
func (m *MockRepositoriesService) ListStatuses(ctx context.Context, owner, repo, ref string, opts *github.ListOptions) ([]*github.RepoStatus, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStatuses", ctx, owner, repo, ref, opts)
	ret0, _ := ret[0].([]*github.RepoStatus)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListStatuses indicates an expected call of ListStatuses.
func (mr *MockRepositoriesServiceMockRecorder) ListStatuses(ctx, owner, repo, ref, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStatuses", reflect.TypeOf((*MockRepositoriesService)(nil).ListStatuses), ctx, owner, repo, ref, opts)
}

// ListTeams mocks base method.
func (m *MockRepositoriesService) ListTeams(ctx context.Context, owner, repo string, opts *github.ListOptions) ([]*github.Team, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeams", ctx, owner, repo, opts)
	ret0, _ := ret[0].([]*github.Team)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTeams indicates an expected call of ListTeams.
func (mr *MockRepositoriesServiceMockRecorder) ListTeams(ctx, owner, repo, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeams", reflect.TypeOf((*MockRepositoriesService)(nil).ListTeams), ctx, owner, repo, opts)
}

// MockTeamsService is a mock of TeamsService interface.
type MockTeamsService struct {
	ctrl     *gomock.Controller
	recorder *MockTeamsServiceMockRecorder
}

// MockTeamsServiceMockRecorder is the mock recorder for MockTeamsService.
type MockTeamsServiceMockRecorder struct {
	mock *MockTeamsService
}

// NewMockTeamsService creates a new mock instance.
func NewMockTeamsService(ctrl *gomock.Controller) *MockTeamsService {
	mock := &MockTeamsService{ctrl: ctrl}
	mock.recorder = &MockTeamsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamsService) EXPECT() *MockTeamsServiceMockRecorder {
	return m.recorder
}

// ListTeamMembersBySlug mocks base method.
func (m *MockTeamsService) ListTeamMembersBySlug(ctx context.Context, org, slug string, opts *github.TeamListTeamMembersOptions) ([]*github.User, *github.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTeamMembersBySlug", ctx, org, slug, opts)
	ret0, _ := ret[0].([]*github.User)
	ret1, _ := ret[1].(*github.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListTeamMembersBySlug indicates an expected call of ListTeamMembersBySlug.
func (mr *MockTeamsServiceMockRecorder) ListTeamMembersBySlug(ctx, org, slug, opts interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTeamMembersBySlug", reflect.TypeOf((*MockTeamsService)(nil).ListTeamMembersBySlug), ctx, org, slug, opts)
}
