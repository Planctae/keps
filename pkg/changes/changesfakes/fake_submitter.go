// Code generated by counterfeiter. DO NOT EDIT.
package changesfakes

import (
	sync "sync"

	changes "github.com/calebamiles/keps/pkg/changes"
)

type FakeSubmitter struct {
	SubmitChangesStub        func() (string, error)
	submitChangesMutex       sync.RWMutex
	submitChangesArgsForCall []struct {
	}
	submitChangesReturns struct {
		result1 string
		result2 error
	}
	submitChangesReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	SubmitterNameStub        func() string
	submitterNameMutex       sync.RWMutex
	submitterNameArgsForCall []struct {
	}
	submitterNameReturns struct {
		result1 string
	}
	submitterNameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubmitter) SubmitChanges() (string, error) {
	fake.submitChangesMutex.Lock()
	ret, specificReturn := fake.submitChangesReturnsOnCall[len(fake.submitChangesArgsForCall)]
	fake.submitChangesArgsForCall = append(fake.submitChangesArgsForCall, struct {
	}{})
	fake.recordInvocation("SubmitChanges", []interface{}{})
	fake.submitChangesMutex.Unlock()
	if fake.SubmitChangesStub != nil {
		return fake.SubmitChangesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.submitChangesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSubmitter) SubmitChangesCallCount() int {
	fake.submitChangesMutex.RLock()
	defer fake.submitChangesMutex.RUnlock()
	return len(fake.submitChangesArgsForCall)
}

func (fake *FakeSubmitter) SubmitChangesCalls(stub func() (string, error)) {
	fake.submitChangesMutex.Lock()
	defer fake.submitChangesMutex.Unlock()
	fake.SubmitChangesStub = stub
}

func (fake *FakeSubmitter) SubmitChangesReturns(result1 string, result2 error) {
	fake.submitChangesMutex.Lock()
	defer fake.submitChangesMutex.Unlock()
	fake.SubmitChangesStub = nil
	fake.submitChangesReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSubmitter) SubmitChangesReturnsOnCall(i int, result1 string, result2 error) {
	fake.submitChangesMutex.Lock()
	defer fake.submitChangesMutex.Unlock()
	fake.SubmitChangesStub = nil
	if fake.submitChangesReturnsOnCall == nil {
		fake.submitChangesReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.submitChangesReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeSubmitter) SubmitterName() string {
	fake.submitterNameMutex.Lock()
	ret, specificReturn := fake.submitterNameReturnsOnCall[len(fake.submitterNameArgsForCall)]
	fake.submitterNameArgsForCall = append(fake.submitterNameArgsForCall, struct {
	}{})
	fake.recordInvocation("SubmitterName", []interface{}{})
	fake.submitterNameMutex.Unlock()
	if fake.SubmitterNameStub != nil {
		return fake.SubmitterNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.submitterNameReturns
	return fakeReturns.result1
}

func (fake *FakeSubmitter) SubmitterNameCallCount() int {
	fake.submitterNameMutex.RLock()
	defer fake.submitterNameMutex.RUnlock()
	return len(fake.submitterNameArgsForCall)
}

func (fake *FakeSubmitter) SubmitterNameCalls(stub func() string) {
	fake.submitterNameMutex.Lock()
	defer fake.submitterNameMutex.Unlock()
	fake.SubmitterNameStub = stub
}

func (fake *FakeSubmitter) SubmitterNameReturns(result1 string) {
	fake.submitterNameMutex.Lock()
	defer fake.submitterNameMutex.Unlock()
	fake.SubmitterNameStub = nil
	fake.submitterNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSubmitter) SubmitterNameReturnsOnCall(i int, result1 string) {
	fake.submitterNameMutex.Lock()
	defer fake.submitterNameMutex.Unlock()
	fake.SubmitterNameStub = nil
	if fake.submitterNameReturnsOnCall == nil {
		fake.submitterNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.submitterNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeSubmitter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.submitChangesMutex.RLock()
	defer fake.submitChangesMutex.RUnlock()
	fake.submitterNameMutex.RLock()
	defer fake.submitterNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubmitter) recordInvocation(key string, args []interface{}) {
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

var _ changes.Submitter = new(FakeSubmitter)
