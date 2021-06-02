// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mattermost/mattermost-mattermod/server (interfaces: CircleCIService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	circleci "github.com/mattermost/go-circleci"
)

// MockCircleCIService is a mock of CircleCIService interface.
type MockCircleCIService struct {
	ctrl     *gomock.Controller
	recorder *MockCircleCIServiceMockRecorder
}

// MockCircleCIServiceMockRecorder is the mock recorder for MockCircleCIService.
type MockCircleCIServiceMockRecorder struct {
	mock *MockCircleCIService
}

// NewMockCircleCIService creates a new mock instance.
func NewMockCircleCIService(ctrl *gomock.Controller) *MockCircleCIService {
	mock := &MockCircleCIService{ctrl: ctrl}
	mock.recorder = &MockCircleCIServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCircleCIService) EXPECT() *MockCircleCIServiceMockRecorder {
	return m.recorder
}

// BuildByProjectWithContext mocks base method.
func (m *MockCircleCIService) BuildByProjectWithContext(arg0 context.Context, arg1 circleci.VcsType, arg2, arg3 string, arg4 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildByProjectWithContext", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuildByProjectWithContext indicates an expected call of BuildByProjectWithContext.
func (mr *MockCircleCIServiceMockRecorder) BuildByProjectWithContext(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildByProjectWithContext", reflect.TypeOf((*MockCircleCIService)(nil).BuildByProjectWithContext), arg0, arg1, arg2, arg3, arg4)
}

// CancelWorkflow mocks base method.
func (m *MockCircleCIService) CancelWorkflow(arg0 string) (*circleci.CancelWorkflow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelWorkflow", arg0)
	ret0, _ := ret[0].(*circleci.CancelWorkflow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelWorkflow indicates an expected call of CancelWorkflow.
func (mr *MockCircleCIServiceMockRecorder) CancelWorkflow(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelWorkflow", reflect.TypeOf((*MockCircleCIService)(nil).CancelWorkflow), arg0)
}

// GetPipelineByBranch mocks base method.
func (m *MockCircleCIService) GetPipelineByBranch(arg0 circleci.VcsType, arg1, arg2, arg3, arg4 string) (*circleci.Pipelines, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineByBranch", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*circleci.Pipelines)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineByBranch indicates an expected call of GetPipelineByBranch.
func (mr *MockCircleCIServiceMockRecorder) GetPipelineByBranch(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineByBranch", reflect.TypeOf((*MockCircleCIService)(nil).GetPipelineByBranch), arg0, arg1, arg2, arg3, arg4)
}

// GetPipelineWorkflow mocks base method.
func (m *MockCircleCIService) GetPipelineWorkflow(arg0, arg1 string) (*circleci.WorkflowList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineWorkflow", arg0, arg1)
	ret0, _ := ret[0].(*circleci.WorkflowList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineWorkflow indicates an expected call of GetPipelineWorkflow.
func (mr *MockCircleCIServiceMockRecorder) GetPipelineWorkflow(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineWorkflow", reflect.TypeOf((*MockCircleCIService)(nil).GetPipelineWorkflow), arg0, arg1)
}

// GetPipelineWorkflowWithContext mocks base method.
func (m *MockCircleCIService) GetPipelineWorkflowWithContext(arg0 context.Context, arg1, arg2 string) (*circleci.WorkflowList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPipelineWorkflowWithContext", arg0, arg1, arg2)
	ret0, _ := ret[0].(*circleci.WorkflowList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPipelineWorkflowWithContext indicates an expected call of GetPipelineWorkflowWithContext.
func (mr *MockCircleCIServiceMockRecorder) GetPipelineWorkflowWithContext(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPipelineWorkflowWithContext", reflect.TypeOf((*MockCircleCIService)(nil).GetPipelineWorkflowWithContext), arg0, arg1, arg2)
}

// ListBuildArtifactsWithContext mocks base method.
func (m *MockCircleCIService) ListBuildArtifactsWithContext(arg0 context.Context, arg1 circleci.VcsType, arg2, arg3 string, arg4 int) ([]*circleci.Artifact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBuildArtifactsWithContext", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*circleci.Artifact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBuildArtifactsWithContext indicates an expected call of ListBuildArtifactsWithContext.
func (mr *MockCircleCIServiceMockRecorder) ListBuildArtifactsWithContext(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBuildArtifactsWithContext", reflect.TypeOf((*MockCircleCIService)(nil).ListBuildArtifactsWithContext), arg0, arg1, arg2, arg3, arg4)
}

// ListRecentBuildsForProjectWithContext mocks base method.
func (m *MockCircleCIService) ListRecentBuildsForProjectWithContext(arg0 context.Context, arg1 circleci.VcsType, arg2, arg3, arg4, arg5 string, arg6, arg7 int) ([]*circleci.Build, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRecentBuildsForProjectWithContext", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].([]*circleci.Build)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRecentBuildsForProjectWithContext indicates an expected call of ListRecentBuildsForProjectWithContext.
func (mr *MockCircleCIServiceMockRecorder) ListRecentBuildsForProjectWithContext(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRecentBuildsForProjectWithContext", reflect.TypeOf((*MockCircleCIService)(nil).ListRecentBuildsForProjectWithContext), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// TriggerPipelineWithContext mocks base method.
func (m *MockCircleCIService) TriggerPipelineWithContext(arg0 context.Context, arg1 circleci.VcsType, arg2, arg3, arg4, arg5 string, arg6 map[string]interface{}) (*circleci.Pipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TriggerPipelineWithContext", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(*circleci.Pipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TriggerPipelineWithContext indicates an expected call of TriggerPipelineWithContext.
func (mr *MockCircleCIServiceMockRecorder) TriggerPipelineWithContext(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TriggerPipelineWithContext", reflect.TypeOf((*MockCircleCIService)(nil).TriggerPipelineWithContext), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
