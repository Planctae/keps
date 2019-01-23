// Code generated by counterfeiter. DO NOT EDIT.
package kepsfakes

import (
	sync "sync"
	time "time"

	keps "github.com/calebamiles/keps/pkg/keps"
	check "github.com/calebamiles/keps/pkg/keps/check"
	states "github.com/calebamiles/keps/pkg/keps/states"
)

type FakeInstance struct {
	AddApproversStub        func(...string)
	addApproversMutex       sync.RWMutex
	addApproversArgsForCall []struct {
		arg1 []string
	}
	AddChecksStub        func(...check.That)
	addChecksMutex       sync.RWMutex
	addChecksArgsForCall []struct {
		arg1 []check.That
	}
	AddReviewersStub        func(...string)
	addReviewersMutex       sync.RWMutex
	addReviewersArgsForCall []struct {
		arg1 []string
	}
	AuthorsStub        func() []string
	authorsMutex       sync.RWMutex
	authorsArgsForCall []struct {
	}
	authorsReturns struct {
		result1 []string
	}
	authorsReturnsOnCall map[int]struct {
		result1 []string
	}
	CheckStub        func() error
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
	}
	checkReturns struct {
		result1 error
	}
	checkReturnsOnCall map[int]struct {
		result1 error
	}
	ContentDirStub        func() string
	contentDirMutex       sync.RWMutex
	contentDirArgsForCall []struct {
	}
	contentDirReturns struct {
		result1 string
	}
	contentDirReturnsOnCall map[int]struct {
		result1 string
	}
	CreatedStub        func() time.Time
	createdMutex       sync.RWMutex
	createdArgsForCall []struct {
	}
	createdReturns struct {
		result1 time.Time
	}
	createdReturnsOnCall map[int]struct {
		result1 time.Time
	}
	LastUpdatedStub        func() time.Time
	lastUpdatedMutex       sync.RWMutex
	lastUpdatedArgsForCall []struct {
	}
	lastUpdatedReturns struct {
		result1 time.Time
	}
	lastUpdatedReturnsOnCall map[int]struct {
		result1 time.Time
	}
	OwningSIGStub        func() string
	owningSIGMutex       sync.RWMutex
	owningSIGArgsForCall []struct {
	}
	owningSIGReturns struct {
		result1 string
	}
	owningSIGReturnsOnCall map[int]struct {
		result1 string
	}
	PersistStub        func() error
	persistMutex       sync.RWMutex
	persistArgsForCall []struct {
	}
	persistReturns struct {
		result1 error
	}
	persistReturnsOnCall map[int]struct {
		result1 error
	}
	SectionsStub        func() []string
	sectionsMutex       sync.RWMutex
	sectionsArgsForCall []struct {
	}
	sectionsReturns struct {
		result1 []string
	}
	sectionsReturnsOnCall map[int]struct {
		result1 []string
	}
	SetStateStub        func(states.Name) error
	setStateMutex       sync.RWMutex
	setStateArgsForCall []struct {
		arg1 states.Name
	}
	setStateReturns struct {
		result1 error
	}
	setStateReturnsOnCall map[int]struct {
		result1 error
	}
	ShortIDStub        func() int
	shortIDMutex       sync.RWMutex
	shortIDArgsForCall []struct {
	}
	shortIDReturns struct {
		result1 int
	}
	shortIDReturnsOnCall map[int]struct {
		result1 int
	}
	StateStub        func() states.Name
	stateMutex       sync.RWMutex
	stateArgsForCall []struct {
	}
	stateReturns struct {
		result1 states.Name
	}
	stateReturnsOnCall map[int]struct {
		result1 states.Name
	}
	TitleStub        func() string
	titleMutex       sync.RWMutex
	titleArgsForCall []struct {
	}
	titleReturns struct {
		result1 string
	}
	titleReturnsOnCall map[int]struct {
		result1 string
	}
	UniqueIDStub        func() string
	uniqueIDMutex       sync.RWMutex
	uniqueIDArgsForCall []struct {
	}
	uniqueIDReturns struct {
		result1 string
	}
	uniqueIDReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInstance) AddApprovers(arg1 ...string) {
	fake.addApproversMutex.Lock()
	fake.addApproversArgsForCall = append(fake.addApproversArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("AddApprovers", []interface{}{arg1})
	fake.addApproversMutex.Unlock()
	if fake.AddApproversStub != nil {
		fake.AddApproversStub(arg1...)
	}
}

func (fake *FakeInstance) AddApproversCallCount() int {
	fake.addApproversMutex.RLock()
	defer fake.addApproversMutex.RUnlock()
	return len(fake.addApproversArgsForCall)
}

func (fake *FakeInstance) AddApproversCalls(stub func(...string)) {
	fake.addApproversMutex.Lock()
	defer fake.addApproversMutex.Unlock()
	fake.AddApproversStub = stub
}

func (fake *FakeInstance) AddApproversArgsForCall(i int) []string {
	fake.addApproversMutex.RLock()
	defer fake.addApproversMutex.RUnlock()
	argsForCall := fake.addApproversArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstance) AddChecks(arg1 ...check.That) {
	fake.addChecksMutex.Lock()
	fake.addChecksArgsForCall = append(fake.addChecksArgsForCall, struct {
		arg1 []check.That
	}{arg1})
	fake.recordInvocation("AddChecks", []interface{}{arg1})
	fake.addChecksMutex.Unlock()
	if fake.AddChecksStub != nil {
		fake.AddChecksStub(arg1...)
	}
}

func (fake *FakeInstance) AddChecksCallCount() int {
	fake.addChecksMutex.RLock()
	defer fake.addChecksMutex.RUnlock()
	return len(fake.addChecksArgsForCall)
}

func (fake *FakeInstance) AddChecksCalls(stub func(...check.That)) {
	fake.addChecksMutex.Lock()
	defer fake.addChecksMutex.Unlock()
	fake.AddChecksStub = stub
}

func (fake *FakeInstance) AddChecksArgsForCall(i int) []check.That {
	fake.addChecksMutex.RLock()
	defer fake.addChecksMutex.RUnlock()
	argsForCall := fake.addChecksArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstance) AddReviewers(arg1 ...string) {
	fake.addReviewersMutex.Lock()
	fake.addReviewersArgsForCall = append(fake.addReviewersArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("AddReviewers", []interface{}{arg1})
	fake.addReviewersMutex.Unlock()
	if fake.AddReviewersStub != nil {
		fake.AddReviewersStub(arg1...)
	}
}

func (fake *FakeInstance) AddReviewersCallCount() int {
	fake.addReviewersMutex.RLock()
	defer fake.addReviewersMutex.RUnlock()
	return len(fake.addReviewersArgsForCall)
}

func (fake *FakeInstance) AddReviewersCalls(stub func(...string)) {
	fake.addReviewersMutex.Lock()
	defer fake.addReviewersMutex.Unlock()
	fake.AddReviewersStub = stub
}

func (fake *FakeInstance) AddReviewersArgsForCall(i int) []string {
	fake.addReviewersMutex.RLock()
	defer fake.addReviewersMutex.RUnlock()
	argsForCall := fake.addReviewersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstance) Authors() []string {
	fake.authorsMutex.Lock()
	ret, specificReturn := fake.authorsReturnsOnCall[len(fake.authorsArgsForCall)]
	fake.authorsArgsForCall = append(fake.authorsArgsForCall, struct {
	}{})
	fake.recordInvocation("Authors", []interface{}{})
	fake.authorsMutex.Unlock()
	if fake.AuthorsStub != nil {
		return fake.AuthorsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.authorsReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) AuthorsCallCount() int {
	fake.authorsMutex.RLock()
	defer fake.authorsMutex.RUnlock()
	return len(fake.authorsArgsForCall)
}

func (fake *FakeInstance) AuthorsCalls(stub func() []string) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = stub
}

func (fake *FakeInstance) AuthorsReturns(result1 []string) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = nil
	fake.authorsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) AuthorsReturnsOnCall(i int, result1 []string) {
	fake.authorsMutex.Lock()
	defer fake.authorsMutex.Unlock()
	fake.AuthorsStub = nil
	if fake.authorsReturnsOnCall == nil {
		fake.authorsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.authorsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) Check() error {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
	}{})
	fake.recordInvocation("Check", []interface{}{})
	fake.checkMutex.Unlock()
	if fake.CheckStub != nil {
		return fake.CheckStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.checkReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeInstance) CheckCalls(stub func() error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeInstance) CheckReturns(result1 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) CheckReturnsOnCall(i int, result1 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) ContentDir() string {
	fake.contentDirMutex.Lock()
	ret, specificReturn := fake.contentDirReturnsOnCall[len(fake.contentDirArgsForCall)]
	fake.contentDirArgsForCall = append(fake.contentDirArgsForCall, struct {
	}{})
	fake.recordInvocation("ContentDir", []interface{}{})
	fake.contentDirMutex.Unlock()
	if fake.ContentDirStub != nil {
		return fake.ContentDirStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.contentDirReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) ContentDirCallCount() int {
	fake.contentDirMutex.RLock()
	defer fake.contentDirMutex.RUnlock()
	return len(fake.contentDirArgsForCall)
}

func (fake *FakeInstance) ContentDirCalls(stub func() string) {
	fake.contentDirMutex.Lock()
	defer fake.contentDirMutex.Unlock()
	fake.ContentDirStub = stub
}

func (fake *FakeInstance) ContentDirReturns(result1 string) {
	fake.contentDirMutex.Lock()
	defer fake.contentDirMutex.Unlock()
	fake.ContentDirStub = nil
	fake.contentDirReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) ContentDirReturnsOnCall(i int, result1 string) {
	fake.contentDirMutex.Lock()
	defer fake.contentDirMutex.Unlock()
	fake.ContentDirStub = nil
	if fake.contentDirReturnsOnCall == nil {
		fake.contentDirReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.contentDirReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) Created() time.Time {
	fake.createdMutex.Lock()
	ret, specificReturn := fake.createdReturnsOnCall[len(fake.createdArgsForCall)]
	fake.createdArgsForCall = append(fake.createdArgsForCall, struct {
	}{})
	fake.recordInvocation("Created", []interface{}{})
	fake.createdMutex.Unlock()
	if fake.CreatedStub != nil {
		return fake.CreatedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createdReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) CreatedCallCount() int {
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	return len(fake.createdArgsForCall)
}

func (fake *FakeInstance) CreatedCalls(stub func() time.Time) {
	fake.createdMutex.Lock()
	defer fake.createdMutex.Unlock()
	fake.CreatedStub = stub
}

func (fake *FakeInstance) CreatedReturns(result1 time.Time) {
	fake.createdMutex.Lock()
	defer fake.createdMutex.Unlock()
	fake.CreatedStub = nil
	fake.createdReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeInstance) CreatedReturnsOnCall(i int, result1 time.Time) {
	fake.createdMutex.Lock()
	defer fake.createdMutex.Unlock()
	fake.CreatedStub = nil
	if fake.createdReturnsOnCall == nil {
		fake.createdReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.createdReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeInstance) LastUpdated() time.Time {
	fake.lastUpdatedMutex.Lock()
	ret, specificReturn := fake.lastUpdatedReturnsOnCall[len(fake.lastUpdatedArgsForCall)]
	fake.lastUpdatedArgsForCall = append(fake.lastUpdatedArgsForCall, struct {
	}{})
	fake.recordInvocation("LastUpdated", []interface{}{})
	fake.lastUpdatedMutex.Unlock()
	if fake.LastUpdatedStub != nil {
		return fake.LastUpdatedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lastUpdatedReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) LastUpdatedCallCount() int {
	fake.lastUpdatedMutex.RLock()
	defer fake.lastUpdatedMutex.RUnlock()
	return len(fake.lastUpdatedArgsForCall)
}

func (fake *FakeInstance) LastUpdatedCalls(stub func() time.Time) {
	fake.lastUpdatedMutex.Lock()
	defer fake.lastUpdatedMutex.Unlock()
	fake.LastUpdatedStub = stub
}

func (fake *FakeInstance) LastUpdatedReturns(result1 time.Time) {
	fake.lastUpdatedMutex.Lock()
	defer fake.lastUpdatedMutex.Unlock()
	fake.LastUpdatedStub = nil
	fake.lastUpdatedReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeInstance) LastUpdatedReturnsOnCall(i int, result1 time.Time) {
	fake.lastUpdatedMutex.Lock()
	defer fake.lastUpdatedMutex.Unlock()
	fake.LastUpdatedStub = nil
	if fake.lastUpdatedReturnsOnCall == nil {
		fake.lastUpdatedReturnsOnCall = make(map[int]struct {
			result1 time.Time
		})
	}
	fake.lastUpdatedReturnsOnCall[i] = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeInstance) OwningSIG() string {
	fake.owningSIGMutex.Lock()
	ret, specificReturn := fake.owningSIGReturnsOnCall[len(fake.owningSIGArgsForCall)]
	fake.owningSIGArgsForCall = append(fake.owningSIGArgsForCall, struct {
	}{})
	fake.recordInvocation("OwningSIG", []interface{}{})
	fake.owningSIGMutex.Unlock()
	if fake.OwningSIGStub != nil {
		return fake.OwningSIGStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.owningSIGReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) OwningSIGCallCount() int {
	fake.owningSIGMutex.RLock()
	defer fake.owningSIGMutex.RUnlock()
	return len(fake.owningSIGArgsForCall)
}

func (fake *FakeInstance) OwningSIGCalls(stub func() string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = stub
}

func (fake *FakeInstance) OwningSIGReturns(result1 string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = nil
	fake.owningSIGReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) OwningSIGReturnsOnCall(i int, result1 string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = nil
	if fake.owningSIGReturnsOnCall == nil {
		fake.owningSIGReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.owningSIGReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) Persist() error {
	fake.persistMutex.Lock()
	ret, specificReturn := fake.persistReturnsOnCall[len(fake.persistArgsForCall)]
	fake.persistArgsForCall = append(fake.persistArgsForCall, struct {
	}{})
	fake.recordInvocation("Persist", []interface{}{})
	fake.persistMutex.Unlock()
	if fake.PersistStub != nil {
		return fake.PersistStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.persistReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) PersistCallCount() int {
	fake.persistMutex.RLock()
	defer fake.persistMutex.RUnlock()
	return len(fake.persistArgsForCall)
}

func (fake *FakeInstance) PersistCalls(stub func() error) {
	fake.persistMutex.Lock()
	defer fake.persistMutex.Unlock()
	fake.PersistStub = stub
}

func (fake *FakeInstance) PersistReturns(result1 error) {
	fake.persistMutex.Lock()
	defer fake.persistMutex.Unlock()
	fake.PersistStub = nil
	fake.persistReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) PersistReturnsOnCall(i int, result1 error) {
	fake.persistMutex.Lock()
	defer fake.persistMutex.Unlock()
	fake.PersistStub = nil
	if fake.persistReturnsOnCall == nil {
		fake.persistReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.persistReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) Sections() []string {
	fake.sectionsMutex.Lock()
	ret, specificReturn := fake.sectionsReturnsOnCall[len(fake.sectionsArgsForCall)]
	fake.sectionsArgsForCall = append(fake.sectionsArgsForCall, struct {
	}{})
	fake.recordInvocation("Sections", []interface{}{})
	fake.sectionsMutex.Unlock()
	if fake.SectionsStub != nil {
		return fake.SectionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sectionsReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) SectionsCallCount() int {
	fake.sectionsMutex.RLock()
	defer fake.sectionsMutex.RUnlock()
	return len(fake.sectionsArgsForCall)
}

func (fake *FakeInstance) SectionsCalls(stub func() []string) {
	fake.sectionsMutex.Lock()
	defer fake.sectionsMutex.Unlock()
	fake.SectionsStub = stub
}

func (fake *FakeInstance) SectionsReturns(result1 []string) {
	fake.sectionsMutex.Lock()
	defer fake.sectionsMutex.Unlock()
	fake.SectionsStub = nil
	fake.sectionsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) SectionsReturnsOnCall(i int, result1 []string) {
	fake.sectionsMutex.Lock()
	defer fake.sectionsMutex.Unlock()
	fake.SectionsStub = nil
	if fake.sectionsReturnsOnCall == nil {
		fake.sectionsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.sectionsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeInstance) SetState(arg1 states.Name) error {
	fake.setStateMutex.Lock()
	ret, specificReturn := fake.setStateReturnsOnCall[len(fake.setStateArgsForCall)]
	fake.setStateArgsForCall = append(fake.setStateArgsForCall, struct {
		arg1 states.Name
	}{arg1})
	fake.recordInvocation("SetState", []interface{}{arg1})
	fake.setStateMutex.Unlock()
	if fake.SetStateStub != nil {
		return fake.SetStateStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setStateReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) SetStateCallCount() int {
	fake.setStateMutex.RLock()
	defer fake.setStateMutex.RUnlock()
	return len(fake.setStateArgsForCall)
}

func (fake *FakeInstance) SetStateCalls(stub func(states.Name) error) {
	fake.setStateMutex.Lock()
	defer fake.setStateMutex.Unlock()
	fake.SetStateStub = stub
}

func (fake *FakeInstance) SetStateArgsForCall(i int) states.Name {
	fake.setStateMutex.RLock()
	defer fake.setStateMutex.RUnlock()
	argsForCall := fake.setStateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInstance) SetStateReturns(result1 error) {
	fake.setStateMutex.Lock()
	defer fake.setStateMutex.Unlock()
	fake.SetStateStub = nil
	fake.setStateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) SetStateReturnsOnCall(i int, result1 error) {
	fake.setStateMutex.Lock()
	defer fake.setStateMutex.Unlock()
	fake.SetStateStub = nil
	if fake.setStateReturnsOnCall == nil {
		fake.setStateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setStateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeInstance) ShortID() int {
	fake.shortIDMutex.Lock()
	ret, specificReturn := fake.shortIDReturnsOnCall[len(fake.shortIDArgsForCall)]
	fake.shortIDArgsForCall = append(fake.shortIDArgsForCall, struct {
	}{})
	fake.recordInvocation("ShortID", []interface{}{})
	fake.shortIDMutex.Unlock()
	if fake.ShortIDStub != nil {
		return fake.ShortIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.shortIDReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) ShortIDCallCount() int {
	fake.shortIDMutex.RLock()
	defer fake.shortIDMutex.RUnlock()
	return len(fake.shortIDArgsForCall)
}

func (fake *FakeInstance) ShortIDCalls(stub func() int) {
	fake.shortIDMutex.Lock()
	defer fake.shortIDMutex.Unlock()
	fake.ShortIDStub = stub
}

func (fake *FakeInstance) ShortIDReturns(result1 int) {
	fake.shortIDMutex.Lock()
	defer fake.shortIDMutex.Unlock()
	fake.ShortIDStub = nil
	fake.shortIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeInstance) ShortIDReturnsOnCall(i int, result1 int) {
	fake.shortIDMutex.Lock()
	defer fake.shortIDMutex.Unlock()
	fake.ShortIDStub = nil
	if fake.shortIDReturnsOnCall == nil {
		fake.shortIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.shortIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeInstance) State() states.Name {
	fake.stateMutex.Lock()
	ret, specificReturn := fake.stateReturnsOnCall[len(fake.stateArgsForCall)]
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct {
	}{})
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stateReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeInstance) StateCalls(stub func() states.Name) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = stub
}

func (fake *FakeInstance) StateReturns(result1 states.Name) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 states.Name
	}{result1}
}

func (fake *FakeInstance) StateReturnsOnCall(i int, result1 states.Name) {
	fake.stateMutex.Lock()
	defer fake.stateMutex.Unlock()
	fake.StateStub = nil
	if fake.stateReturnsOnCall == nil {
		fake.stateReturnsOnCall = make(map[int]struct {
			result1 states.Name
		})
	}
	fake.stateReturnsOnCall[i] = struct {
		result1 states.Name
	}{result1}
}

func (fake *FakeInstance) Title() string {
	fake.titleMutex.Lock()
	ret, specificReturn := fake.titleReturnsOnCall[len(fake.titleArgsForCall)]
	fake.titleArgsForCall = append(fake.titleArgsForCall, struct {
	}{})
	fake.recordInvocation("Title", []interface{}{})
	fake.titleMutex.Unlock()
	if fake.TitleStub != nil {
		return fake.TitleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.titleReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) TitleCallCount() int {
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	return len(fake.titleArgsForCall)
}

func (fake *FakeInstance) TitleCalls(stub func() string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = stub
}

func (fake *FakeInstance) TitleReturns(result1 string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = nil
	fake.titleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) TitleReturnsOnCall(i int, result1 string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = nil
	if fake.titleReturnsOnCall == nil {
		fake.titleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.titleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) UniqueID() string {
	fake.uniqueIDMutex.Lock()
	ret, specificReturn := fake.uniqueIDReturnsOnCall[len(fake.uniqueIDArgsForCall)]
	fake.uniqueIDArgsForCall = append(fake.uniqueIDArgsForCall, struct {
	}{})
	fake.recordInvocation("UniqueID", []interface{}{})
	fake.uniqueIDMutex.Unlock()
	if fake.UniqueIDStub != nil {
		return fake.UniqueIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.uniqueIDReturns
	return fakeReturns.result1
}

func (fake *FakeInstance) UniqueIDCallCount() int {
	fake.uniqueIDMutex.RLock()
	defer fake.uniqueIDMutex.RUnlock()
	return len(fake.uniqueIDArgsForCall)
}

func (fake *FakeInstance) UniqueIDCalls(stub func() string) {
	fake.uniqueIDMutex.Lock()
	defer fake.uniqueIDMutex.Unlock()
	fake.UniqueIDStub = stub
}

func (fake *FakeInstance) UniqueIDReturns(result1 string) {
	fake.uniqueIDMutex.Lock()
	defer fake.uniqueIDMutex.Unlock()
	fake.UniqueIDStub = nil
	fake.uniqueIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) UniqueIDReturnsOnCall(i int, result1 string) {
	fake.uniqueIDMutex.Lock()
	defer fake.uniqueIDMutex.Unlock()
	fake.UniqueIDStub = nil
	if fake.uniqueIDReturnsOnCall == nil {
		fake.uniqueIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.uniqueIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addApproversMutex.RLock()
	defer fake.addApproversMutex.RUnlock()
	fake.addChecksMutex.RLock()
	defer fake.addChecksMutex.RUnlock()
	fake.addReviewersMutex.RLock()
	defer fake.addReviewersMutex.RUnlock()
	fake.authorsMutex.RLock()
	defer fake.authorsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	fake.contentDirMutex.RLock()
	defer fake.contentDirMutex.RUnlock()
	fake.createdMutex.RLock()
	defer fake.createdMutex.RUnlock()
	fake.lastUpdatedMutex.RLock()
	defer fake.lastUpdatedMutex.RUnlock()
	fake.owningSIGMutex.RLock()
	defer fake.owningSIGMutex.RUnlock()
	fake.persistMutex.RLock()
	defer fake.persistMutex.RUnlock()
	fake.sectionsMutex.RLock()
	defer fake.sectionsMutex.RUnlock()
	fake.setStateMutex.RLock()
	defer fake.setStateMutex.RUnlock()
	fake.shortIDMutex.RLock()
	defer fake.shortIDMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	fake.uniqueIDMutex.RLock()
	defer fake.uniqueIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInstance) recordInvocation(key string, args []interface{}) {
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

var _ keps.Instance = new(FakeInstance)
