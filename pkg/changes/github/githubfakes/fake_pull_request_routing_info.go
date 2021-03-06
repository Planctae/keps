// Code generated by counterfeiter. DO NOT EDIT.
package githubfakes

import (
	sync "sync"

	auth "github.com/calebamiles/keps/pkg/changes/auth"
	github "github.com/calebamiles/keps/pkg/changes/github"
)

type FakePullRequestRoutingInfo struct {
	SourceBranchStub        func() string
	sourceBranchMutex       sync.RWMutex
	sourceBranchArgsForCall []struct {
	}
	sourceBranchReturns struct {
		result1 string
	}
	sourceBranchReturnsOnCall map[int]struct {
		result1 string
	}
	SourceRepositoryStub        func() string
	sourceRepositoryMutex       sync.RWMutex
	sourceRepositoryArgsForCall []struct {
	}
	sourceRepositoryReturns struct {
		result1 string
	}
	sourceRepositoryReturnsOnCall map[int]struct {
		result1 string
	}
	SourceRepositoryOwnerStub        func() string
	sourceRepositoryOwnerMutex       sync.RWMutex
	sourceRepositoryOwnerArgsForCall []struct {
	}
	sourceRepositoryOwnerReturns struct {
		result1 string
	}
	sourceRepositoryOwnerReturnsOnCall map[int]struct {
		result1 string
	}
	TargetBranchStub        func() string
	targetBranchMutex       sync.RWMutex
	targetBranchArgsForCall []struct {
	}
	targetBranchReturns struct {
		result1 string
	}
	targetBranchReturnsOnCall map[int]struct {
		result1 string
	}
	TargetRepositoryStub        func() string
	targetRepositoryMutex       sync.RWMutex
	targetRepositoryArgsForCall []struct {
	}
	targetRepositoryReturns struct {
		result1 string
	}
	targetRepositoryReturnsOnCall map[int]struct {
		result1 string
	}
	TargetRepositoryOwnerStub        func() string
	targetRepositoryOwnerMutex       sync.RWMutex
	targetRepositoryOwnerArgsForCall []struct {
	}
	targetRepositoryOwnerReturns struct {
		result1 string
	}
	targetRepositoryOwnerReturnsOnCall map[int]struct {
		result1 string
	}
	TokenStub        func() auth.TokenProvider
	tokenMutex       sync.RWMutex
	tokenArgsForCall []struct {
	}
	tokenReturns struct {
		result1 auth.TokenProvider
	}
	tokenReturnsOnCall map[int]struct {
		result1 auth.TokenProvider
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePullRequestRoutingInfo) SourceBranch() string {
	fake.sourceBranchMutex.Lock()
	ret, specificReturn := fake.sourceBranchReturnsOnCall[len(fake.sourceBranchArgsForCall)]
	fake.sourceBranchArgsForCall = append(fake.sourceBranchArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceBranch", []interface{}{})
	fake.sourceBranchMutex.Unlock()
	if fake.SourceBranchStub != nil {
		return fake.SourceBranchStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceBranchReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) SourceBranchCallCount() int {
	fake.sourceBranchMutex.RLock()
	defer fake.sourceBranchMutex.RUnlock()
	return len(fake.sourceBranchArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) SourceBranchCalls(stub func() string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = stub
}

func (fake *FakePullRequestRoutingInfo) SourceBranchReturns(result1 string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = nil
	fake.sourceBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) SourceBranchReturnsOnCall(i int, result1 string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = nil
	if fake.sourceBranchReturnsOnCall == nil {
		fake.sourceBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) SourceRepository() string {
	fake.sourceRepositoryMutex.Lock()
	ret, specificReturn := fake.sourceRepositoryReturnsOnCall[len(fake.sourceRepositoryArgsForCall)]
	fake.sourceRepositoryArgsForCall = append(fake.sourceRepositoryArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceRepository", []interface{}{})
	fake.sourceRepositoryMutex.Unlock()
	if fake.SourceRepositoryStub != nil {
		return fake.SourceRepositoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceRepositoryReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryCallCount() int {
	fake.sourceRepositoryMutex.RLock()
	defer fake.sourceRepositoryMutex.RUnlock()
	return len(fake.sourceRepositoryArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryCalls(stub func() string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = stub
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryReturns(result1 string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = nil
	fake.sourceRepositoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryReturnsOnCall(i int, result1 string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = nil
	if fake.sourceRepositoryReturnsOnCall == nil {
		fake.sourceRepositoryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceRepositoryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryOwner() string {
	fake.sourceRepositoryOwnerMutex.Lock()
	ret, specificReturn := fake.sourceRepositoryOwnerReturnsOnCall[len(fake.sourceRepositoryOwnerArgsForCall)]
	fake.sourceRepositoryOwnerArgsForCall = append(fake.sourceRepositoryOwnerArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceRepositoryOwner", []interface{}{})
	fake.sourceRepositoryOwnerMutex.Unlock()
	if fake.SourceRepositoryOwnerStub != nil {
		return fake.SourceRepositoryOwnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceRepositoryOwnerReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryOwnerCallCount() int {
	fake.sourceRepositoryOwnerMutex.RLock()
	defer fake.sourceRepositoryOwnerMutex.RUnlock()
	return len(fake.sourceRepositoryOwnerArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryOwnerCalls(stub func() string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = stub
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryOwnerReturns(result1 string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = nil
	fake.sourceRepositoryOwnerReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) SourceRepositoryOwnerReturnsOnCall(i int, result1 string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = nil
	if fake.sourceRepositoryOwnerReturnsOnCall == nil {
		fake.sourceRepositoryOwnerReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceRepositoryOwnerReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetBranch() string {
	fake.targetBranchMutex.Lock()
	ret, specificReturn := fake.targetBranchReturnsOnCall[len(fake.targetBranchArgsForCall)]
	fake.targetBranchArgsForCall = append(fake.targetBranchArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetBranch", []interface{}{})
	fake.targetBranchMutex.Unlock()
	if fake.TargetBranchStub != nil {
		return fake.TargetBranchStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetBranchReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) TargetBranchCallCount() int {
	fake.targetBranchMutex.RLock()
	defer fake.targetBranchMutex.RUnlock()
	return len(fake.targetBranchArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) TargetBranchCalls(stub func() string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = stub
}

func (fake *FakePullRequestRoutingInfo) TargetBranchReturns(result1 string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = nil
	fake.targetBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetBranchReturnsOnCall(i int, result1 string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = nil
	if fake.targetBranchReturnsOnCall == nil {
		fake.targetBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetRepository() string {
	fake.targetRepositoryMutex.Lock()
	ret, specificReturn := fake.targetRepositoryReturnsOnCall[len(fake.targetRepositoryArgsForCall)]
	fake.targetRepositoryArgsForCall = append(fake.targetRepositoryArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetRepository", []interface{}{})
	fake.targetRepositoryMutex.Unlock()
	if fake.TargetRepositoryStub != nil {
		return fake.TargetRepositoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetRepositoryReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryCallCount() int {
	fake.targetRepositoryMutex.RLock()
	defer fake.targetRepositoryMutex.RUnlock()
	return len(fake.targetRepositoryArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryCalls(stub func() string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = stub
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryReturns(result1 string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = nil
	fake.targetRepositoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryReturnsOnCall(i int, result1 string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = nil
	if fake.targetRepositoryReturnsOnCall == nil {
		fake.targetRepositoryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetRepositoryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryOwner() string {
	fake.targetRepositoryOwnerMutex.Lock()
	ret, specificReturn := fake.targetRepositoryOwnerReturnsOnCall[len(fake.targetRepositoryOwnerArgsForCall)]
	fake.targetRepositoryOwnerArgsForCall = append(fake.targetRepositoryOwnerArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetRepositoryOwner", []interface{}{})
	fake.targetRepositoryOwnerMutex.Unlock()
	if fake.TargetRepositoryOwnerStub != nil {
		return fake.TargetRepositoryOwnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetRepositoryOwnerReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryOwnerCallCount() int {
	fake.targetRepositoryOwnerMutex.RLock()
	defer fake.targetRepositoryOwnerMutex.RUnlock()
	return len(fake.targetRepositoryOwnerArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryOwnerCalls(stub func() string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = stub
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryOwnerReturns(result1 string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = nil
	fake.targetRepositoryOwnerReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TargetRepositoryOwnerReturnsOnCall(i int, result1 string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = nil
	if fake.targetRepositoryOwnerReturnsOnCall == nil {
		fake.targetRepositoryOwnerReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetRepositoryOwnerReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) Token() auth.TokenProvider {
	fake.tokenMutex.Lock()
	ret, specificReturn := fake.tokenReturnsOnCall[len(fake.tokenArgsForCall)]
	fake.tokenArgsForCall = append(fake.tokenArgsForCall, struct {
	}{})
	fake.recordInvocation("Token", []interface{}{})
	fake.tokenMutex.Unlock()
	if fake.TokenStub != nil {
		return fake.TokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tokenReturns
	return fakeReturns.result1
}

func (fake *FakePullRequestRoutingInfo) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakePullRequestRoutingInfo) TokenCalls(stub func() auth.TokenProvider) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = stub
}

func (fake *FakePullRequestRoutingInfo) TokenReturns(result1 auth.TokenProvider) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 auth.TokenProvider
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) TokenReturnsOnCall(i int, result1 auth.TokenProvider) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	if fake.tokenReturnsOnCall == nil {
		fake.tokenReturnsOnCall = make(map[int]struct {
			result1 auth.TokenProvider
		})
	}
	fake.tokenReturnsOnCall[i] = struct {
		result1 auth.TokenProvider
	}{result1}
}

func (fake *FakePullRequestRoutingInfo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sourceBranchMutex.RLock()
	defer fake.sourceBranchMutex.RUnlock()
	fake.sourceRepositoryMutex.RLock()
	defer fake.sourceRepositoryMutex.RUnlock()
	fake.sourceRepositoryOwnerMutex.RLock()
	defer fake.sourceRepositoryOwnerMutex.RUnlock()
	fake.targetBranchMutex.RLock()
	defer fake.targetBranchMutex.RUnlock()
	fake.targetRepositoryMutex.RLock()
	defer fake.targetRepositoryMutex.RUnlock()
	fake.targetRepositoryOwnerMutex.RLock()
	defer fake.targetRepositoryOwnerMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePullRequestRoutingInfo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ github.PullRequestRoutingInfo = new(FakePullRequestRoutingInfo)
