// This file was generated by counterfeiter
package volumefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/baggageclaim/volume"
)

type FakeStrategy struct {
	MaterializeStub        func(lager.Logger, string, volume.Filesystem) (volume.FilesystemInitVolume, error)
	materializeMutex       sync.RWMutex
	materializeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 volume.Filesystem
	}
	materializeReturns struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStrategy) Materialize(arg1 lager.Logger, arg2 string, arg3 volume.Filesystem) (volume.FilesystemInitVolume, error) {
	fake.materializeMutex.Lock()
	fake.materializeArgsForCall = append(fake.materializeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 volume.Filesystem
	}{arg1, arg2, arg3})
	fake.recordInvocation("Materialize", []interface{}{arg1, arg2, arg3})
	fake.materializeMutex.Unlock()
	if fake.MaterializeStub != nil {
		return fake.MaterializeStub(arg1, arg2, arg3)
	} else {
		return fake.materializeReturns.result1, fake.materializeReturns.result2
	}
}

func (fake *FakeStrategy) MaterializeCallCount() int {
	fake.materializeMutex.RLock()
	defer fake.materializeMutex.RUnlock()
	return len(fake.materializeArgsForCall)
}

func (fake *FakeStrategy) MaterializeArgsForCall(i int) (lager.Logger, string, volume.Filesystem) {
	fake.materializeMutex.RLock()
	defer fake.materializeMutex.RUnlock()
	return fake.materializeArgsForCall[i].arg1, fake.materializeArgsForCall[i].arg2, fake.materializeArgsForCall[i].arg3
}

func (fake *FakeStrategy) MaterializeReturns(result1 volume.FilesystemInitVolume, result2 error) {
	fake.MaterializeStub = nil
	fake.materializeReturns = struct {
		result1 volume.FilesystemInitVolume
		result2 error
	}{result1, result2}
}

func (fake *FakeStrategy) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.materializeMutex.RLock()
	defer fake.materializeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStrategy) recordInvocation(key string, args []interface{}) {
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

var _ volume.Strategy = new(FakeStrategy)
