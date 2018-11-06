// Code generated by counterfeiter. DO NOT EDIT.
package settingsfakes

import (
	sync "sync"

	settings "github.com/calebamiles/keps/pkg/settings"
)

type FakeRuntime struct {
	ContentRootStub        func() string
	contentRootMutex       sync.RWMutex
	contentRootArgsForCall []struct {
	}
	contentRootReturns struct {
		result1 string
	}
	contentRootReturnsOnCall map[int]struct {
		result1 string
	}
	PrincipalStub        func() string
	principalMutex       sync.RWMutex
	principalArgsForCall []struct {
	}
	principalReturns struct {
		result1 string
	}
	principalReturnsOnCall map[int]struct {
		result1 string
	}
	TargetDirStub        func() string
	targetDirMutex       sync.RWMutex
	targetDirArgsForCall []struct {
	}
	targetDirReturns struct {
		result1 string
	}
	targetDirReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuntime) ContentRoot() string {
	fake.contentRootMutex.Lock()
	ret, specificReturn := fake.contentRootReturnsOnCall[len(fake.contentRootArgsForCall)]
	fake.contentRootArgsForCall = append(fake.contentRootArgsForCall, struct {
	}{})
	fake.recordInvocation("ContentRoot", []interface{}{})
	fake.contentRootMutex.Unlock()
	if fake.ContentRootStub != nil {
		return fake.ContentRootStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contentRootReturns
	return fakeReturns.result1
}

func (fake *FakeRuntime) ContentRootCallCount() int {
	fake.contentRootMutex.RLock()
	defer fake.contentRootMutex.RUnlock()
	return len(fake.contentRootArgsForCall)
}

func (fake *FakeRuntime) ContentRootReturns(result1 string) {
	fake.ContentRootStub = nil
	fake.contentRootReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) ContentRootReturnsOnCall(i int, result1 string) {
	fake.ContentRootStub = nil
	if fake.contentRootReturnsOnCall == nil {
		fake.contentRootReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.contentRootReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) Principal() string {
	fake.principalMutex.Lock()
	ret, specificReturn := fake.principalReturnsOnCall[len(fake.principalArgsForCall)]
	fake.principalArgsForCall = append(fake.principalArgsForCall, struct {
	}{})
	fake.recordInvocation("Principal", []interface{}{})
	fake.principalMutex.Unlock()
	if fake.PrincipalStub != nil {
		return fake.PrincipalStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.principalReturns
	return fakeReturns.result1
}

func (fake *FakeRuntime) PrincipalCallCount() int {
	fake.principalMutex.RLock()
	defer fake.principalMutex.RUnlock()
	return len(fake.principalArgsForCall)
}

func (fake *FakeRuntime) PrincipalReturns(result1 string) {
	fake.PrincipalStub = nil
	fake.principalReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) PrincipalReturnsOnCall(i int, result1 string) {
	fake.PrincipalStub = nil
	if fake.principalReturnsOnCall == nil {
		fake.principalReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.principalReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) TargetDir() string {
	fake.targetDirMutex.Lock()
	ret, specificReturn := fake.targetDirReturnsOnCall[len(fake.targetDirArgsForCall)]
	fake.targetDirArgsForCall = append(fake.targetDirArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetDir", []interface{}{})
	fake.targetDirMutex.Unlock()
	if fake.TargetDirStub != nil {
		return fake.TargetDirStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetDirReturns
	return fakeReturns.result1
}

func (fake *FakeRuntime) TargetDirCallCount() int {
	fake.targetDirMutex.RLock()
	defer fake.targetDirMutex.RUnlock()
	return len(fake.targetDirArgsForCall)
}

func (fake *FakeRuntime) TargetDirReturns(result1 string) {
	fake.TargetDirStub = nil
	fake.targetDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) TargetDirReturnsOnCall(i int, result1 string) {
	fake.TargetDirStub = nil
	if fake.targetDirReturnsOnCall == nil {
		fake.targetDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.contentRootMutex.RLock()
	defer fake.contentRootMutex.RUnlock()
	fake.principalMutex.RLock()
	defer fake.principalMutex.RUnlock()
	fake.targetDirMutex.RLock()
	defer fake.targetDirMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRuntime) recordInvocation(key string, args []interface{}) {
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

var _ settings.Runtime = new(FakeRuntime)