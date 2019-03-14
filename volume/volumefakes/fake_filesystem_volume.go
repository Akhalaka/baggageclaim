// Code generated by counterfeiter. DO NOT EDIT.
package volumefakes

import (
	sync "sync"

	volume "github.com/concourse/baggageclaim/volume"
)

type FakeFilesystemVolume struct {
	DataPathStub        func() string
	dataPathMutex       sync.RWMutex
	dataPathArgsForCall []struct {
	}
	dataPathReturns struct {
		result1 string
	}
	dataPathReturnsOnCall map[int]struct {
		result1 string
	}
	DestroyStub        func() error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct {
	}
	handleReturns struct {
		result1 string
	}
	handleReturnsOnCall map[int]struct {
		result1 string
	}
	LoadPrivilegedStub        func() (bool, error)
	loadPrivilegedMutex       sync.RWMutex
	loadPrivilegedArgsForCall []struct {
	}
	loadPrivilegedReturns struct {
		result1 bool
		result2 error
	}
	loadPrivilegedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	LoadPropertiesStub        func() (volume.Properties, error)
	loadPropertiesMutex       sync.RWMutex
	loadPropertiesArgsForCall []struct {
	}
	loadPropertiesReturns struct {
		result1 volume.Properties
		result2 error
	}
	loadPropertiesReturnsOnCall map[int]struct {
		result1 volume.Properties
		result2 error
	}
	ParentStub        func() (volume.FilesystemLiveVolume, bool, error)
	parentMutex       sync.RWMutex
	parentArgsForCall []struct {
	}
	parentReturns struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}
	parentReturnsOnCall map[int]struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}
	StorePrivilegedStub        func(bool) error
	storePrivilegedMutex       sync.RWMutex
	storePrivilegedArgsForCall []struct {
		arg1 bool
	}
	storePrivilegedReturns struct {
		result1 error
	}
	storePrivilegedReturnsOnCall map[int]struct {
		result1 error
	}
	StorePropertiesStub        func(volume.Properties) error
	storePropertiesMutex       sync.RWMutex
	storePropertiesArgsForCall []struct {
		arg1 volume.Properties
	}
	storePropertiesReturns struct {
		result1 error
	}
	storePropertiesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFilesystemVolume) DataPath() string {
	fake.dataPathMutex.Lock()
	ret, specificReturn := fake.dataPathReturnsOnCall[len(fake.dataPathArgsForCall)]
	fake.dataPathArgsForCall = append(fake.dataPathArgsForCall, struct {
	}{})
	fake.recordInvocation("DataPath", []interface{}{})
	fake.dataPathMutex.Unlock()
	if fake.DataPathStub != nil {
		return fake.DataPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dataPathReturns
	return fakeReturns.result1
}

func (fake *FakeFilesystemVolume) DataPathCallCount() int {
	fake.dataPathMutex.RLock()
	defer fake.dataPathMutex.RUnlock()
	return len(fake.dataPathArgsForCall)
}

func (fake *FakeFilesystemVolume) DataPathCalls(stub func() string) {
	fake.dataPathMutex.Lock()
	defer fake.dataPathMutex.Unlock()
	fake.DataPathStub = stub
}

func (fake *FakeFilesystemVolume) DataPathReturns(result1 string) {
	fake.dataPathMutex.Lock()
	defer fake.dataPathMutex.Unlock()
	fake.DataPathStub = nil
	fake.dataPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemVolume) DataPathReturnsOnCall(i int, result1 string) {
	fake.dataPathMutex.Lock()
	defer fake.dataPathMutex.Unlock()
	fake.DataPathStub = nil
	if fake.dataPathReturnsOnCall == nil {
		fake.dataPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.dataPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemVolume) Destroy() error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
	}{})
	fake.recordInvocation("Destroy", []interface{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeFilesystemVolume) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeFilesystemVolume) DestroyCalls(stub func() error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeFilesystemVolume) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) Handle() string {
	fake.handleMutex.Lock()
	ret, specificReturn := fake.handleReturnsOnCall[len(fake.handleArgsForCall)]
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct {
	}{})
	fake.recordInvocation("Handle", []interface{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleReturns
	return fakeReturns.result1
}

func (fake *FakeFilesystemVolume) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeFilesystemVolume) HandleCalls(stub func() string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = stub
}

func (fake *FakeFilesystemVolume) HandleReturns(result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemVolume) HandleReturnsOnCall(i int, result1 string) {
	fake.handleMutex.Lock()
	defer fake.handleMutex.Unlock()
	fake.HandleStub = nil
	if fake.handleReturnsOnCall == nil {
		fake.handleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.handleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeFilesystemVolume) LoadPrivileged() (bool, error) {
	fake.loadPrivilegedMutex.Lock()
	ret, specificReturn := fake.loadPrivilegedReturnsOnCall[len(fake.loadPrivilegedArgsForCall)]
	fake.loadPrivilegedArgsForCall = append(fake.loadPrivilegedArgsForCall, struct {
	}{})
	fake.recordInvocation("LoadPrivileged", []interface{}{})
	fake.loadPrivilegedMutex.Unlock()
	if fake.LoadPrivilegedStub != nil {
		return fake.LoadPrivilegedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadPrivilegedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFilesystemVolume) LoadPrivilegedCallCount() int {
	fake.loadPrivilegedMutex.RLock()
	defer fake.loadPrivilegedMutex.RUnlock()
	return len(fake.loadPrivilegedArgsForCall)
}

func (fake *FakeFilesystemVolume) LoadPrivilegedCalls(stub func() (bool, error)) {
	fake.loadPrivilegedMutex.Lock()
	defer fake.loadPrivilegedMutex.Unlock()
	fake.LoadPrivilegedStub = stub
}

func (fake *FakeFilesystemVolume) LoadPrivilegedReturns(result1 bool, result2 error) {
	fake.loadPrivilegedMutex.Lock()
	defer fake.loadPrivilegedMutex.Unlock()
	fake.LoadPrivilegedStub = nil
	fake.loadPrivilegedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemVolume) LoadPrivilegedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.loadPrivilegedMutex.Lock()
	defer fake.loadPrivilegedMutex.Unlock()
	fake.LoadPrivilegedStub = nil
	if fake.loadPrivilegedReturnsOnCall == nil {
		fake.loadPrivilegedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.loadPrivilegedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemVolume) LoadProperties() (volume.Properties, error) {
	fake.loadPropertiesMutex.Lock()
	ret, specificReturn := fake.loadPropertiesReturnsOnCall[len(fake.loadPropertiesArgsForCall)]
	fake.loadPropertiesArgsForCall = append(fake.loadPropertiesArgsForCall, struct {
	}{})
	fake.recordInvocation("LoadProperties", []interface{}{})
	fake.loadPropertiesMutex.Unlock()
	if fake.LoadPropertiesStub != nil {
		return fake.LoadPropertiesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.loadPropertiesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFilesystemVolume) LoadPropertiesCallCount() int {
	fake.loadPropertiesMutex.RLock()
	defer fake.loadPropertiesMutex.RUnlock()
	return len(fake.loadPropertiesArgsForCall)
}

func (fake *FakeFilesystemVolume) LoadPropertiesCalls(stub func() (volume.Properties, error)) {
	fake.loadPropertiesMutex.Lock()
	defer fake.loadPropertiesMutex.Unlock()
	fake.LoadPropertiesStub = stub
}

func (fake *FakeFilesystemVolume) LoadPropertiesReturns(result1 volume.Properties, result2 error) {
	fake.loadPropertiesMutex.Lock()
	defer fake.loadPropertiesMutex.Unlock()
	fake.LoadPropertiesStub = nil
	fake.loadPropertiesReturns = struct {
		result1 volume.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemVolume) LoadPropertiesReturnsOnCall(i int, result1 volume.Properties, result2 error) {
	fake.loadPropertiesMutex.Lock()
	defer fake.loadPropertiesMutex.Unlock()
	fake.LoadPropertiesStub = nil
	if fake.loadPropertiesReturnsOnCall == nil {
		fake.loadPropertiesReturnsOnCall = make(map[int]struct {
			result1 volume.Properties
			result2 error
		})
	}
	fake.loadPropertiesReturnsOnCall[i] = struct {
		result1 volume.Properties
		result2 error
	}{result1, result2}
}

func (fake *FakeFilesystemVolume) Parent() (volume.FilesystemLiveVolume, bool, error) {
	fake.parentMutex.Lock()
	ret, specificReturn := fake.parentReturnsOnCall[len(fake.parentArgsForCall)]
	fake.parentArgsForCall = append(fake.parentArgsForCall, struct {
	}{})
	fake.recordInvocation("Parent", []interface{}{})
	fake.parentMutex.Unlock()
	if fake.ParentStub != nil {
		return fake.ParentStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.parentReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeFilesystemVolume) ParentCallCount() int {
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	return len(fake.parentArgsForCall)
}

func (fake *FakeFilesystemVolume) ParentCalls(stub func() (volume.FilesystemLiveVolume, bool, error)) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = stub
}

func (fake *FakeFilesystemVolume) ParentReturns(result1 volume.FilesystemLiveVolume, result2 bool, result3 error) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = nil
	fake.parentReturns = struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemVolume) ParentReturnsOnCall(i int, result1 volume.FilesystemLiveVolume, result2 bool, result3 error) {
	fake.parentMutex.Lock()
	defer fake.parentMutex.Unlock()
	fake.ParentStub = nil
	if fake.parentReturnsOnCall == nil {
		fake.parentReturnsOnCall = make(map[int]struct {
			result1 volume.FilesystemLiveVolume
			result2 bool
			result3 error
		})
	}
	fake.parentReturnsOnCall[i] = struct {
		result1 volume.FilesystemLiveVolume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFilesystemVolume) StorePrivileged(arg1 bool) error {
	fake.storePrivilegedMutex.Lock()
	ret, specificReturn := fake.storePrivilegedReturnsOnCall[len(fake.storePrivilegedArgsForCall)]
	fake.storePrivilegedArgsForCall = append(fake.storePrivilegedArgsForCall, struct {
		arg1 bool
	}{arg1})
	fake.recordInvocation("StorePrivileged", []interface{}{arg1})
	fake.storePrivilegedMutex.Unlock()
	if fake.StorePrivilegedStub != nil {
		return fake.StorePrivilegedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.storePrivilegedReturns
	return fakeReturns.result1
}

func (fake *FakeFilesystemVolume) StorePrivilegedCallCount() int {
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	return len(fake.storePrivilegedArgsForCall)
}

func (fake *FakeFilesystemVolume) StorePrivilegedCalls(stub func(bool) error) {
	fake.storePrivilegedMutex.Lock()
	defer fake.storePrivilegedMutex.Unlock()
	fake.StorePrivilegedStub = stub
}

func (fake *FakeFilesystemVolume) StorePrivilegedArgsForCall(i int) bool {
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	argsForCall := fake.storePrivilegedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFilesystemVolume) StorePrivilegedReturns(result1 error) {
	fake.storePrivilegedMutex.Lock()
	defer fake.storePrivilegedMutex.Unlock()
	fake.StorePrivilegedStub = nil
	fake.storePrivilegedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) StorePrivilegedReturnsOnCall(i int, result1 error) {
	fake.storePrivilegedMutex.Lock()
	defer fake.storePrivilegedMutex.Unlock()
	fake.StorePrivilegedStub = nil
	if fake.storePrivilegedReturnsOnCall == nil {
		fake.storePrivilegedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storePrivilegedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) StoreProperties(arg1 volume.Properties) error {
	fake.storePropertiesMutex.Lock()
	ret, specificReturn := fake.storePropertiesReturnsOnCall[len(fake.storePropertiesArgsForCall)]
	fake.storePropertiesArgsForCall = append(fake.storePropertiesArgsForCall, struct {
		arg1 volume.Properties
	}{arg1})
	fake.recordInvocation("StoreProperties", []interface{}{arg1})
	fake.storePropertiesMutex.Unlock()
	if fake.StorePropertiesStub != nil {
		return fake.StorePropertiesStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.storePropertiesReturns
	return fakeReturns.result1
}

func (fake *FakeFilesystemVolume) StorePropertiesCallCount() int {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	return len(fake.storePropertiesArgsForCall)
}

func (fake *FakeFilesystemVolume) StorePropertiesCalls(stub func(volume.Properties) error) {
	fake.storePropertiesMutex.Lock()
	defer fake.storePropertiesMutex.Unlock()
	fake.StorePropertiesStub = stub
}

func (fake *FakeFilesystemVolume) StorePropertiesArgsForCall(i int) volume.Properties {
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	argsForCall := fake.storePropertiesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeFilesystemVolume) StorePropertiesReturns(result1 error) {
	fake.storePropertiesMutex.Lock()
	defer fake.storePropertiesMutex.Unlock()
	fake.StorePropertiesStub = nil
	fake.storePropertiesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) StorePropertiesReturnsOnCall(i int, result1 error) {
	fake.storePropertiesMutex.Lock()
	defer fake.storePropertiesMutex.Unlock()
	fake.StorePropertiesStub = nil
	if fake.storePropertiesReturnsOnCall == nil {
		fake.storePropertiesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.storePropertiesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeFilesystemVolume) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dataPathMutex.RLock()
	defer fake.dataPathMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	fake.loadPrivilegedMutex.RLock()
	defer fake.loadPrivilegedMutex.RUnlock()
	fake.loadPropertiesMutex.RLock()
	defer fake.loadPropertiesMutex.RUnlock()
	fake.parentMutex.RLock()
	defer fake.parentMutex.RUnlock()
	fake.storePrivilegedMutex.RLock()
	defer fake.storePrivilegedMutex.RUnlock()
	fake.storePropertiesMutex.RLock()
	defer fake.storePropertiesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFilesystemVolume) recordInvocation(key string, args []interface{}) {
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

var _ volume.FilesystemVolume = new(FakeFilesystemVolume)
