// This file was generated by counterfeiter
package fake_translator

import (
	"os"
	"os/exec"
	"sync"

	"github.com/concourse/baggageclaim/uidgid"
)

type FakeTranslator struct {
	CacheKeyStub        func() string
	cacheKeyMutex       sync.RWMutex
	cacheKeyArgsForCall []struct{}
	cacheKeyReturns     struct {
		result1 string
	}
	TranslatePathStub        func(path string, info os.FileInfo, err error) error
	translatePathMutex       sync.RWMutex
	translatePathArgsForCall []struct {
		path string
		info os.FileInfo
		err  error
	}
	translatePathReturns struct {
		result1 error
	}
	TranslateCommandStub        func(*exec.Cmd)
	translateCommandMutex       sync.RWMutex
	translateCommandArgsForCall []struct {
		arg1 *exec.Cmd
	}
}

func (fake *FakeTranslator) CacheKey() string {
	fake.cacheKeyMutex.Lock()
	fake.cacheKeyArgsForCall = append(fake.cacheKeyArgsForCall, struct{}{})
	fake.cacheKeyMutex.Unlock()
	if fake.CacheKeyStub != nil {
		return fake.CacheKeyStub()
	} else {
		return fake.cacheKeyReturns.result1
	}
}

func (fake *FakeTranslator) CacheKeyCallCount() int {
	fake.cacheKeyMutex.RLock()
	defer fake.cacheKeyMutex.RUnlock()
	return len(fake.cacheKeyArgsForCall)
}

func (fake *FakeTranslator) CacheKeyReturns(result1 string) {
	fake.CacheKeyStub = nil
	fake.cacheKeyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeTranslator) TranslatePath(path string, info os.FileInfo, err error) error {
	fake.translatePathMutex.Lock()
	fake.translatePathArgsForCall = append(fake.translatePathArgsForCall, struct {
		path string
		info os.FileInfo
		err  error
	}{path, info, err})
	fake.translatePathMutex.Unlock()
	if fake.TranslatePathStub != nil {
		return fake.TranslatePathStub(path, info, err)
	} else {
		return fake.translatePathReturns.result1
	}
}

func (fake *FakeTranslator) TranslatePathCallCount() int {
	fake.translatePathMutex.RLock()
	defer fake.translatePathMutex.RUnlock()
	return len(fake.translatePathArgsForCall)
}

func (fake *FakeTranslator) TranslatePathArgsForCall(i int) (string, os.FileInfo, error) {
	fake.translatePathMutex.RLock()
	defer fake.translatePathMutex.RUnlock()
	return fake.translatePathArgsForCall[i].path, fake.translatePathArgsForCall[i].info, fake.translatePathArgsForCall[i].err
}

func (fake *FakeTranslator) TranslatePathReturns(result1 error) {
	fake.TranslatePathStub = nil
	fake.translatePathReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTranslator) TranslateCommand(arg1 *exec.Cmd) {
	fake.translateCommandMutex.Lock()
	fake.translateCommandArgsForCall = append(fake.translateCommandArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	fake.translateCommandMutex.Unlock()
	if fake.TranslateCommandStub != nil {
		fake.TranslateCommandStub(arg1)
	}
}

func (fake *FakeTranslator) TranslateCommandCallCount() int {
	fake.translateCommandMutex.RLock()
	defer fake.translateCommandMutex.RUnlock()
	return len(fake.translateCommandArgsForCall)
}

func (fake *FakeTranslator) TranslateCommandArgsForCall(i int) *exec.Cmd {
	fake.translateCommandMutex.RLock()
	defer fake.translateCommandMutex.RUnlock()
	return fake.translateCommandArgsForCall[i].arg1
}

var _ uidgid.Translator = new(FakeTranslator)