// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/jenkins-x/jx/v2/pkg/cloud/amazon/eks (interfaces: EKSer)

package eks_test

import (
	"reflect"
	"time"

	eks "github.com/aws/aws-sdk-go/service/eks"
	cluster "github.com/jenkins-x/jx/v2/pkg/cluster"
	pegomock "github.com/petergtz/pegomock"
)

type MockEKSer struct {
	fail func(message string, callerSkip ...int)
}

func NewMockEKSer(options ...pegomock.Option) *MockEKSer {
	mock := &MockEKSer{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockEKSer) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockEKSer) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockEKSer) AddTagsToCluster(_param0 string, _param1 map[string]*string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AddTagsToCluster", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockEKSer) CleanUpObsoleteEksClusterStack(_param0 string, _param1 string, _param2 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CleanUpObsoleteEksClusterStack", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockEKSer) DescribeCluster(_param0 string) (*cluster.Cluster, string, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("DescribeCluster", params, []reflect.Type{reflect.TypeOf((**cluster.Cluster)(nil)).Elem(), reflect.TypeOf((*string)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *cluster.Cluster
	var ret1 string
	var ret2 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*cluster.Cluster)
		}
		if result[1] != nil {
			ret1 = result[1].(string)
		}
		if result[2] != nil {
			ret2 = result[2].(error)
		}
	}
	return ret0, ret1, ret2
}

func (mock *MockEKSer) EksClusterExists(_param0 string, _param1 string, _param2 string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("EksClusterExists", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEKSer) EksClusterObsoleteStackExists(_param0 string, _param1 string, _param2 string) (bool, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("EksClusterObsoleteStackExists", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 bool
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEKSer) GetClusterAsEKSCluster(_param0 string) (*eks.Cluster, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{_param0}
	result := pegomock.GetGenericMockFrom(mock).Invoke("GetClusterAsEKSCluster", params, []reflect.Type{reflect.TypeOf((**eks.Cluster)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *eks.Cluster
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*eks.Cluster)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEKSer) ListClusters() ([]*cluster.Cluster, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockEKSer().")
	}
	params := []pegomock.Param{}
	result := pegomock.GetGenericMockFrom(mock).Invoke("ListClusters", params, []reflect.Type{reflect.TypeOf((*[]*cluster.Cluster)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 []*cluster.Cluster
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].([]*cluster.Cluster)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockEKSer) VerifyWasCalledOnce() *VerifierMockEKSer {
	return &VerifierMockEKSer{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockEKSer) VerifyWasCalled(invocationCountMatcher pegomock.Matcher) *VerifierMockEKSer {
	return &VerifierMockEKSer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockEKSer) VerifyWasCalledInOrder(invocationCountMatcher pegomock.Matcher, inOrderContext *pegomock.InOrderContext) *VerifierMockEKSer {
	return &VerifierMockEKSer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockEKSer) VerifyWasCalledEventually(invocationCountMatcher pegomock.Matcher, timeout time.Duration) *VerifierMockEKSer {
	return &VerifierMockEKSer{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockEKSer struct {
	mock                   *MockEKSer
	invocationCountMatcher pegomock.Matcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockEKSer) AddTagsToCluster(_param0 string, _param1 map[string]*string) *MockEKSer_AddTagsToCluster_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AddTagsToCluster", params, verifier.timeout)
	return &MockEKSer_AddTagsToCluster_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_AddTagsToCluster_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_AddTagsToCluster_OngoingVerification) GetCapturedArguments() (string, map[string]*string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockEKSer_AddTagsToCluster_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []map[string]*string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]map[string]*string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(map[string]*string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) CleanUpObsoleteEksClusterStack(_param0 string, _param1 string, _param2 string) *MockEKSer_CleanUpObsoleteEksClusterStack_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CleanUpObsoleteEksClusterStack", params, verifier.timeout)
	return &MockEKSer_CleanUpObsoleteEksClusterStack_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_CleanUpObsoleteEksClusterStack_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_CleanUpObsoleteEksClusterStack_OngoingVerification) GetCapturedArguments() (string, string, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockEKSer_CleanUpObsoleteEksClusterStack_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
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
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) DescribeCluster(_param0 string) *MockEKSer_DescribeCluster_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "DescribeCluster", params, verifier.timeout)
	return &MockEKSer_DescribeCluster_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_DescribeCluster_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_DescribeCluster_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockEKSer_DescribeCluster_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) EksClusterExists(_param0 string, _param1 string, _param2 string) *MockEKSer_EksClusterExists_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "EksClusterExists", params, verifier.timeout)
	return &MockEKSer_EksClusterExists_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_EksClusterExists_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_EksClusterExists_OngoingVerification) GetCapturedArguments() (string, string, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockEKSer_EksClusterExists_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
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
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) EksClusterObsoleteStackExists(_param0 string, _param1 string, _param2 string) *MockEKSer_EksClusterObsoleteStackExists_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "EksClusterObsoleteStackExists", params, verifier.timeout)
	return &MockEKSer_EksClusterObsoleteStackExists_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_EksClusterObsoleteStackExists_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_EksClusterObsoleteStackExists_OngoingVerification) GetCapturedArguments() (string, string, string) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockEKSer_EksClusterObsoleteStackExists_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string, _param2 []string) {
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
		_param2 = make([]string, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) GetClusterAsEKSCluster(_param0 string) *MockEKSer_GetClusterAsEKSCluster_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "GetClusterAsEKSCluster", params, verifier.timeout)
	return &MockEKSer_GetClusterAsEKSCluster_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_GetClusterAsEKSCluster_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_GetClusterAsEKSCluster_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockEKSer_GetClusterAsEKSCluster_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockEKSer) ListClusters() *MockEKSer_ListClusters_OngoingVerification {
	params := []pegomock.Param{}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "ListClusters", params, verifier.timeout)
	return &MockEKSer_ListClusters_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockEKSer_ListClusters_OngoingVerification struct {
	mock              *MockEKSer
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockEKSer_ListClusters_OngoingVerification) GetCapturedArguments() {
}

func (c *MockEKSer_ListClusters_OngoingVerification) GetAllCapturedArguments() {
}
