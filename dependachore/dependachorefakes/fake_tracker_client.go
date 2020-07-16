// Code generated by counterfeiter. DO NOT EDIT.
package dependachorefakes

import (
	"sync"

	"github.com/masters-of-cats/dependachore/dependachore"
	"github.com/masters-of-cats/dependachore/tracker"
)

type FakeTrackerClient struct {
	GetStub        func(int) (tracker.Story, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 int
	}
	getReturns struct {
		result1 tracker.Story
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 tracker.Story
		result2 error
	}
	MoveAndChorifyStub        func(int, int) error
	moveAndChorifyMutex       sync.RWMutex
	moveAndChorifyArgsForCall []struct {
		arg1 int
		arg2 int
	}
	moveAndChorifyReturns struct {
		result1 error
	}
	moveAndChorifyReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTrackerClient) Get(arg1 int) (tracker.Story, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeTrackerClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeTrackerClient) GetCalls(stub func(int) (tracker.Story, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeTrackerClient) GetArgsForCall(i int) int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeTrackerClient) GetReturns(result1 tracker.Story, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 tracker.Story
		result2 error
	}{result1, result2}
}

func (fake *FakeTrackerClient) GetReturnsOnCall(i int, result1 tracker.Story, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 tracker.Story
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 tracker.Story
		result2 error
	}{result1, result2}
}

func (fake *FakeTrackerClient) MoveAndChorify(arg1 int, arg2 int) error {
	fake.moveAndChorifyMutex.Lock()
	ret, specificReturn := fake.moveAndChorifyReturnsOnCall[len(fake.moveAndChorifyArgsForCall)]
	fake.moveAndChorifyArgsForCall = append(fake.moveAndChorifyArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("MoveAndChorify", []interface{}{arg1, arg2})
	fake.moveAndChorifyMutex.Unlock()
	if fake.MoveAndChorifyStub != nil {
		return fake.MoveAndChorifyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.moveAndChorifyReturns
	return fakeReturns.result1
}

func (fake *FakeTrackerClient) MoveAndChorifyCallCount() int {
	fake.moveAndChorifyMutex.RLock()
	defer fake.moveAndChorifyMutex.RUnlock()
	return len(fake.moveAndChorifyArgsForCall)
}

func (fake *FakeTrackerClient) MoveAndChorifyCalls(stub func(int, int) error) {
	fake.moveAndChorifyMutex.Lock()
	defer fake.moveAndChorifyMutex.Unlock()
	fake.MoveAndChorifyStub = stub
}

func (fake *FakeTrackerClient) MoveAndChorifyArgsForCall(i int) (int, int) {
	fake.moveAndChorifyMutex.RLock()
	defer fake.moveAndChorifyMutex.RUnlock()
	argsForCall := fake.moveAndChorifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeTrackerClient) MoveAndChorifyReturns(result1 error) {
	fake.moveAndChorifyMutex.Lock()
	defer fake.moveAndChorifyMutex.Unlock()
	fake.MoveAndChorifyStub = nil
	fake.moveAndChorifyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerClient) MoveAndChorifyReturnsOnCall(i int, result1 error) {
	fake.moveAndChorifyMutex.Lock()
	defer fake.moveAndChorifyMutex.Unlock()
	fake.MoveAndChorifyStub = nil
	if fake.moveAndChorifyReturnsOnCall == nil {
		fake.moveAndChorifyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.moveAndChorifyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeTrackerClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.moveAndChorifyMutex.RLock()
	defer fake.moveAndChorifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTrackerClient) recordInvocation(key string, args []interface{}) {
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

var _ dependachore.TrackerClient = new(FakeTrackerClient)
