// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/v2/pkg/util (interfaces: Commander)

package util_test

import (
	"reflect"
	"time"

	backoff "github.com/cenkalti/backoff"
	pegomock "github.com/petergtz/pegomock"
)

type MockCommander struct {
	fail func(message string, callerSkip ...int)
}

func NewMockCommander(options ...pegomock.Option) *MockCommander {
	mock := &MockCommander{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockCommander) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockCommander) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockCommander) CurrentArgs() []string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentArgs", params, []reflect.Type{reflect.TypeOf((*[]string)(nil)).Elem()})
	var ret0 []string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentDir() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentDir", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentEnv() map[string]string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentEnv", params, []reflect.Type{reflect.TypeOf((*map[string]string)(nil)).Elem()})
	var ret0 map[string]string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(map[string]string)
		}
	}
	return ret0
}

func (mock *MockCommander) CurrentName() string {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CurrentName", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem()})
	var ret0 string
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
	}
	return ret0
}

func (mock *MockCommander) DidError() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidError", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommander) DidFail() bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DidFail", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockCommander) Error() error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Error", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockCommander) Run() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Run", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommander) RunWithoutRetry() (string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("RunWithoutRetry", params, []reflect.Type{reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 string
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(string)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockCommander) SetArgs(_param0 []string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetArgs", params, []reflect.Type{})
}

func (mock *MockCommander) SetDir(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetDir", params, []reflect.Type{})
}

func (mock *MockCommander) SetEnv(_param0 map[string]string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetEnv", params, []reflect.Type{})
}

func (mock *MockCommander) SetEnvVariable(_param0 string, _param1 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0, _param1}
	pegomock.GetGenericMockFrom(mock).Invoke("SetEnvVariable", params, []reflect.Type{})
}

func (mock *MockCommander) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetExponentialBackOff", params, []reflect.Type{})
}

func (mock *MockCommander) SetName(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetName", params, []reflect.Type{})
}

func (mock *MockCommander) SetTimeout(_param0 time.Duration) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockCommander().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("SetTimeout", params, []reflect.Type{})
}

func (mock *MockCommander) VerifyWasCalledOnce() *VerifierMockCommander {
	return &VerifierMockCommander{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockCommander) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockCommander {
	return &VerifierMockCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockCommander) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockCommander {
	return &VerifierMockCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockCommander) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockCommander {
	return &VerifierMockCommander{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockCommander struct {
	mock                   *MockCommander
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockCommander) CurrentArgs() *MockCommander_CurrentArgs_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentArgs", params, verifier.timeout)
	return &MockCommander_CurrentArgs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_CurrentArgs_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_CurrentArgs_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_CurrentArgs_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) CurrentDir() *MockCommander_CurrentDir_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentDir", params, verifier.timeout)
	return &MockCommander_CurrentDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_CurrentDir_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_CurrentDir_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_CurrentDir_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) CurrentEnv() *MockCommander_CurrentEnv_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentEnv", params, verifier.timeout)
	return &MockCommander_CurrentEnv_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_CurrentEnv_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_CurrentEnv_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_CurrentEnv_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) CurrentName() *MockCommander_CurrentName_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CurrentName", params, verifier.timeout)
	return &MockCommander_CurrentName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_CurrentName_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_CurrentName_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_CurrentName_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) DidError() *MockCommander_DidError_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidError", params, verifier.timeout)
	return &MockCommander_DidError_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_DidError_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_DidError_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_DidError_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) DidFail() *MockCommander_DidFail_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DidFail", params, verifier.timeout)
	return &MockCommander_DidFail_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_DidFail_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_DidFail_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_DidFail_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) Error() *MockCommander_Error_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Error", params, verifier.timeout)
	return &MockCommander_Error_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_Error_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_Error_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_Error_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) Run() *MockCommander_Run_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Run", params, verifier.timeout)
	return &MockCommander_Run_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_Run_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_Run_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_Run_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) RunWithoutRetry() *MockCommander_RunWithoutRetry_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RunWithoutRetry", params, verifier.timeout)
	return &MockCommander_RunWithoutRetry_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_RunWithoutRetry_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_RunWithoutRetry_OngoingVerification) GetCapturedArguments() {
}

func (c *MockCommander_RunWithoutRetry_OngoingVerification) GetAllCapturedArguments() {
}

func (verifier *VerifierMockCommander) SetArgs(_param0 []string) *MockCommander_SetArgs_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetArgs", params, verifier.timeout)
	return &MockCommander_SetArgs_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetArgs_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetArgs_OngoingVerification) GetCapturedArguments() []string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetArgs_OngoingVerification) GetAllCapturedArguments() (_param0 [][]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([][]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.([]string)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetDir(_param0 string) *MockCommander_SetDir_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetDir", params, verifier.timeout)
	return &MockCommander_SetDir_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetDir_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetDir_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetDir_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetEnv(_param0 map[string]string) *MockCommander_SetEnv_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetEnv", params, verifier.timeout)
	return &MockCommander_SetEnv_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetEnv_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetEnv_OngoingVerification) GetCapturedArguments() map[string]string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetEnv_OngoingVerification) GetAllCapturedArguments() (_param0 []map[string]string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]map[string]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(map[string]string)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetEnvVariable(_param0 string, _param1 string) *MockCommander_SetEnvVariable_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetEnvVariable", params, verifier.timeout)
	return &MockCommander_SetEnvVariable_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetEnvVariable_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetEnvVariable_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockCommander_SetEnvVariable_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetExponentialBackOff(_param0 *backoff.ExponentialBackOff) *MockCommander_SetExponentialBackOff_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetExponentialBackOff", params, verifier.timeout)
	return &MockCommander_SetExponentialBackOff_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetExponentialBackOff_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetExponentialBackOff_OngoingVerification) GetCapturedArguments() *backoff.ExponentialBackOff {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetExponentialBackOff_OngoingVerification) GetAllCapturedArguments() (_param0 []*backoff.ExponentialBackOff) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*backoff.ExponentialBackOff, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*backoff.ExponentialBackOff)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetName(_param0 string) *MockCommander_SetName_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetName", params, verifier.timeout)
	return &MockCommander_SetName_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetName_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetName_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetName_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockCommander) SetTimeout(_param0 time.Duration) *MockCommander_SetTimeout_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetTimeout", params, verifier.timeout)
	return &MockCommander_SetTimeout_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockCommander_SetTimeout_OngoingVerification struct {
	mock              *MockCommander
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockCommander_SetTimeout_OngoingVerification) GetCapturedArguments() time.Duration {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockCommander_SetTimeout_OngoingVerification) GetAllCapturedArguments() (_param0 []time.Duration) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]time.Duration, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(time.Duration)
		}
	}
	return
}
