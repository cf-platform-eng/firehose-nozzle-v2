// Code generated by counterfeiter. DO NOT EDIT.
package nozzlefakes

import (
	"gateway/nozzle"
	"sync"
)

type FakeUAA struct {
	GetAuthTokenStub        func() (string, error)
	getAuthTokenMutex       sync.RWMutex
	getAuthTokenArgsForCall []struct {
	}
	getAuthTokenReturns struct {
		result1 string
		result2 error
	}
	getAuthTokenReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUAA) GetAuthToken() (string, error) {
	fake.getAuthTokenMutex.Lock()
	ret, specificReturn := fake.getAuthTokenReturnsOnCall[len(fake.getAuthTokenArgsForCall)]
	fake.getAuthTokenArgsForCall = append(fake.getAuthTokenArgsForCall, struct {
	}{})
	fake.recordInvocation("GetAuthToken", []interface{}{})
	fake.getAuthTokenMutex.Unlock()
	if fake.GetAuthTokenStub != nil {
		return fake.GetAuthTokenStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getAuthTokenReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUAA) GetAuthTokenCallCount() int {
	fake.getAuthTokenMutex.RLock()
	defer fake.getAuthTokenMutex.RUnlock()
	return len(fake.getAuthTokenArgsForCall)
}

func (fake *FakeUAA) GetAuthTokenCalls(stub func() (string, error)) {
	fake.getAuthTokenMutex.Lock()
	defer fake.getAuthTokenMutex.Unlock()
	fake.GetAuthTokenStub = stub
}

func (fake *FakeUAA) GetAuthTokenReturns(result1 string, result2 error) {
	fake.getAuthTokenMutex.Lock()
	defer fake.getAuthTokenMutex.Unlock()
	fake.GetAuthTokenStub = nil
	fake.getAuthTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) GetAuthTokenReturnsOnCall(i int, result1 string, result2 error) {
	fake.getAuthTokenMutex.Lock()
	defer fake.getAuthTokenMutex.Unlock()
	fake.GetAuthTokenStub = nil
	if fake.getAuthTokenReturnsOnCall == nil {
		fake.getAuthTokenReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getAuthTokenReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeUAA) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAuthTokenMutex.RLock()
	defer fake.getAuthTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUAA) recordInvocation(key string, args []interface{}) {
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

var _ nozzle.UAA = new(FakeUAA)