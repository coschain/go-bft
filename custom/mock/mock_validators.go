// Code generated by MockGen. DO NOT EDIT.
// Source: ../Ivalidator.go

// Package mock is a generated GoMock package.
package mock

import (
	custom "github.com/coschain/gobft/custom"
	message "github.com/coschain/gobft/message"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICommittee is a mock of ICommittee interface
type MockICommittee struct {
	ctrl     *gomock.Controller
	recorder *MockICommitteeMockRecorder
}

// MockICommitteeMockRecorder is the mock recorder for MockICommittee
type MockICommitteeMockRecorder struct {
	mock *MockICommittee
}

// NewMockICommittee creates a new mock instance
func NewMockICommittee(ctrl *gomock.Controller) *MockICommittee {
	mock := &MockICommittee{ctrl: ctrl}
	mock.recorder = &MockICommitteeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICommittee) EXPECT() *MockICommitteeMockRecorder {
	return m.recorder
}

// GetValidator mocks base method
func (m *MockICommittee) GetValidator(key message.PubKey) custom.IPubValidator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidator", key)
	ret0, _ := ret[0].(custom.IPubValidator)
	return ret0
}

// GetValidator indicates an expected call of GetValidator
func (mr *MockICommitteeMockRecorder) GetValidator(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidator", reflect.TypeOf((*MockICommittee)(nil).GetValidator), key)
}

// IsValidator mocks base method
func (m *MockICommittee) IsValidator(key message.PubKey) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator
func (mr *MockICommitteeMockRecorder) IsValidator(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockICommittee)(nil).IsValidator), key)
}

// TotalVotingPower mocks base method
func (m *MockICommittee) TotalVotingPower() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalVotingPower")
	ret0, _ := ret[0].(int64)
	return ret0
}

// TotalVotingPower indicates an expected call of TotalVotingPower
func (mr *MockICommitteeMockRecorder) TotalVotingPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalVotingPower", reflect.TypeOf((*MockICommittee)(nil).TotalVotingPower))
}

// GetCurrentProposer mocks base method
func (m *MockICommittee) GetCurrentProposer(round int) message.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentProposer", round)
	ret0, _ := ret[0].(message.PubKey)
	return ret0
}

// GetCurrentProposer indicates an expected call of GetCurrentProposer
func (mr *MockICommitteeMockRecorder) GetCurrentProposer(round interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentProposer", reflect.TypeOf((*MockICommittee)(nil).GetCurrentProposer), round)
}

// DecidesProposal mocks base method
func (m *MockICommittee) DecidesProposal() message.ProposedData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecidesProposal")
	ret0, _ := ret[0].(message.ProposedData)
	return ret0
}

// DecidesProposal indicates an expected call of DecidesProposal
func (mr *MockICommitteeMockRecorder) DecidesProposal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecidesProposal", reflect.TypeOf((*MockICommittee)(nil).DecidesProposal))
}

// ValidateProposal mocks base method
func (m *MockICommittee) ValidateProposal(data message.ProposedData) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateProposal", data)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ValidateProposal indicates an expected call of ValidateProposal
func (mr *MockICommitteeMockRecorder) ValidateProposal(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProposal", reflect.TypeOf((*MockICommittee)(nil).ValidateProposal), data)
}

// Commit mocks base method
func (m *MockICommittee) Commit(commitRecords *message.Commit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", commitRecords)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockICommitteeMockRecorder) Commit(commitRecords interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockICommittee)(nil).Commit), commitRecords)
}

// GetAppState mocks base method
func (m *MockICommittee) GetAppState() *message.AppState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppState")
	ret0, _ := ret[0].(*message.AppState)
	return ret0
}

// GetAppState indicates an expected call of GetAppState
func (mr *MockICommitteeMockRecorder) GetAppState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppState", reflect.TypeOf((*MockICommittee)(nil).GetAppState))
}

// BroadCast mocks base method
func (m *MockICommittee) BroadCast(msg message.ConsensusMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BroadCast", msg)
	ret0, _ := ret[0].(error)
	return ret0
}

// BroadCast indicates an expected call of BroadCast
func (mr *MockICommitteeMockRecorder) BroadCast(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadCast", reflect.TypeOf((*MockICommittee)(nil).BroadCast), msg)
}

// MockIPubValidator is a mock of IPubValidator interface
type MockIPubValidator struct {
	ctrl     *gomock.Controller
	recorder *MockIPubValidatorMockRecorder
}

// MockIPubValidatorMockRecorder is the mock recorder for MockIPubValidator
type MockIPubValidatorMockRecorder struct {
	mock *MockIPubValidator
}

// NewMockIPubValidator creates a new mock instance
func NewMockIPubValidator(ctrl *gomock.Controller) *MockIPubValidator {
	mock := &MockIPubValidator{ctrl: ctrl}
	mock.recorder = &MockIPubValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPubValidator) EXPECT() *MockIPubValidatorMockRecorder {
	return m.recorder
}

// VerifySig mocks base method
func (m *MockIPubValidator) VerifySig(digest, signature []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifySig", digest, signature)
	ret0, _ := ret[0].(bool)
	return ret0
}

// VerifySig indicates an expected call of VerifySig
func (mr *MockIPubValidatorMockRecorder) VerifySig(digest, signature interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifySig", reflect.TypeOf((*MockIPubValidator)(nil).VerifySig), digest, signature)
}

// GetPubKey mocks base method
func (m *MockIPubValidator) GetPubKey() message.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey")
	ret0, _ := ret[0].(message.PubKey)
	return ret0
}

// GetPubKey indicates an expected call of GetPubKey
func (mr *MockIPubValidatorMockRecorder) GetPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockIPubValidator)(nil).GetPubKey))
}

// GetVotingPower mocks base method
func (m *MockIPubValidator) GetVotingPower() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotingPower")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetVotingPower indicates an expected call of GetVotingPower
func (mr *MockIPubValidatorMockRecorder) GetVotingPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotingPower", reflect.TypeOf((*MockIPubValidator)(nil).GetVotingPower))
}

// SetVotingPower mocks base method
func (m *MockIPubValidator) SetVotingPower(arg0 int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVotingPower", arg0)
}

// SetVotingPower indicates an expected call of SetVotingPower
func (mr *MockIPubValidatorMockRecorder) SetVotingPower(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVotingPower", reflect.TypeOf((*MockIPubValidator)(nil).SetVotingPower), arg0)
}

// MockIPrivValidator is a mock of IPrivValidator interface
type MockIPrivValidator struct {
	ctrl     *gomock.Controller
	recorder *MockIPrivValidatorMockRecorder
}

// MockIPrivValidatorMockRecorder is the mock recorder for MockIPrivValidator
type MockIPrivValidatorMockRecorder struct {
	mock *MockIPrivValidator
}

// NewMockIPrivValidator creates a new mock instance
func NewMockIPrivValidator(ctrl *gomock.Controller) *MockIPrivValidator {
	mock := &MockIPrivValidator{ctrl: ctrl}
	mock.recorder = &MockIPrivValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPrivValidator) EXPECT() *MockIPrivValidatorMockRecorder {
	return m.recorder
}

// GetPubKey mocks base method
func (m *MockIPrivValidator) GetPubKey() message.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey")
	ret0, _ := ret[0].(message.PubKey)
	return ret0
}

// GetPubKey indicates an expected call of GetPubKey
func (mr *MockIPrivValidatorMockRecorder) GetPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockIPrivValidator)(nil).GetPubKey))
}

// Sign mocks base method
func (m *MockIPrivValidator) Sign(digest []byte) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", digest)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Sign indicates an expected call of Sign
func (mr *MockIPrivValidatorMockRecorder) Sign(digest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockIPrivValidator)(nil).Sign), digest)
}
