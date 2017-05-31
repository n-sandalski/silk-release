// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type DatabaseHandler struct {
	AllStub        func() ([]controller.Lease, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []controller.Lease
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []controller.Lease
		result2 error
	}
	AllActiveStub        func(int) ([]controller.Lease, error)
	allActiveMutex       sync.RWMutex
	allActiveArgsForCall []struct {
		arg1 int
	}
	allActiveReturns struct {
		result1 []controller.Lease
		result2 error
	}
	allActiveReturnsOnCall map[int]struct {
		result1 []controller.Lease
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DatabaseHandler) All() ([]controller.Lease, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *DatabaseHandler) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *DatabaseHandler) AllReturns(result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AllReturnsOnCall(i int, result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []controller.Lease
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AllActive(arg1 int) ([]controller.Lease, error) {
	fake.allActiveMutex.Lock()
	ret, specificReturn := fake.allActiveReturnsOnCall[len(fake.allActiveArgsForCall)]
	fake.allActiveArgsForCall = append(fake.allActiveArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("AllActive", []interface{}{arg1})
	fake.allActiveMutex.Unlock()
	if fake.AllActiveStub != nil {
		return fake.AllActiveStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allActiveReturns.result1, fake.allActiveReturns.result2
}

func (fake *DatabaseHandler) AllActiveCallCount() int {
	fake.allActiveMutex.RLock()
	defer fake.allActiveMutex.RUnlock()
	return len(fake.allActiveArgsForCall)
}

func (fake *DatabaseHandler) AllActiveArgsForCall(i int) int {
	fake.allActiveMutex.RLock()
	defer fake.allActiveMutex.RUnlock()
	return fake.allActiveArgsForCall[i].arg1
}

func (fake *DatabaseHandler) AllActiveReturns(result1 []controller.Lease, result2 error) {
	fake.AllActiveStub = nil
	fake.allActiveReturns = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AllActiveReturnsOnCall(i int, result1 []controller.Lease, result2 error) {
	fake.AllActiveStub = nil
	if fake.allActiveReturnsOnCall == nil {
		fake.allActiveReturnsOnCall = make(map[int]struct {
			result1 []controller.Lease
			result2 error
		})
	}
	fake.allActiveReturnsOnCall[i] = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.allActiveMutex.RLock()
	defer fake.allActiveMutex.RUnlock()
	return fake.invocations
}

func (fake *DatabaseHandler) recordInvocation(key string, args []interface{}) {
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
