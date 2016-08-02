// This file was generated by counterfeiter
package modelsfakes

import (
	"net"
	"sync"

	"github.com/cloudfoundry-incubator/switchboard/models"
)

type FakeBridges struct {
	CreateStub        func(clientConn, backendConn net.Conn) models.Bridge
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		clientConn  net.Conn
		backendConn net.Conn
	}
	createReturns struct {
		result1 models.Bridge
	}
	RemoveStub        func(bridge models.Bridge) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		bridge models.Bridge
	}
	removeReturns struct {
		result1 error
	}
	RemoveAndCloseAllStub        func()
	removeAndCloseAllMutex       sync.RWMutex
	removeAndCloseAllArgsForCall []struct{}
	SizeStub                     func() uint
	sizeMutex                    sync.RWMutex
	sizeArgsForCall              []struct{}
	sizeReturns                  struct {
		result1 uint
	}
	ContainsStub        func(bridge models.Bridge) bool
	containsMutex       sync.RWMutex
	containsArgsForCall []struct {
		bridge models.Bridge
	}
	containsReturns struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBridges) Create(clientConn net.Conn, backendConn net.Conn) models.Bridge {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		clientConn  net.Conn
		backendConn net.Conn
	}{clientConn, backendConn})
	fake.recordInvocation("Create", []interface{}{clientConn, backendConn})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(clientConn, backendConn)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeBridges) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBridges) CreateArgsForCall(i int) (net.Conn, net.Conn) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].clientConn, fake.createArgsForCall[i].backendConn
}

func (fake *FakeBridges) CreateReturns(result1 models.Bridge) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 models.Bridge
	}{result1}
}

func (fake *FakeBridges) Remove(bridge models.Bridge) error {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		bridge models.Bridge
	}{bridge})
	fake.recordInvocation("Remove", []interface{}{bridge})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(bridge)
	} else {
		return fake.removeReturns.result1
	}
}

func (fake *FakeBridges) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeBridges) RemoveArgsForCall(i int) models.Bridge {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return fake.removeArgsForCall[i].bridge
}

func (fake *FakeBridges) RemoveReturns(result1 error) {
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBridges) RemoveAndCloseAll() {
	fake.removeAndCloseAllMutex.Lock()
	fake.removeAndCloseAllArgsForCall = append(fake.removeAndCloseAllArgsForCall, struct{}{})
	fake.recordInvocation("RemoveAndCloseAll", []interface{}{})
	fake.removeAndCloseAllMutex.Unlock()
	if fake.RemoveAndCloseAllStub != nil {
		fake.RemoveAndCloseAllStub()
	}
}

func (fake *FakeBridges) RemoveAndCloseAllCallCount() int {
	fake.removeAndCloseAllMutex.RLock()
	defer fake.removeAndCloseAllMutex.RUnlock()
	return len(fake.removeAndCloseAllArgsForCall)
}

func (fake *FakeBridges) Size() uint {
	fake.sizeMutex.Lock()
	fake.sizeArgsForCall = append(fake.sizeArgsForCall, struct{}{})
	fake.recordInvocation("Size", []interface{}{})
	fake.sizeMutex.Unlock()
	if fake.SizeStub != nil {
		return fake.SizeStub()
	} else {
		return fake.sizeReturns.result1
	}
}

func (fake *FakeBridges) SizeCallCount() int {
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	return len(fake.sizeArgsForCall)
}

func (fake *FakeBridges) SizeReturns(result1 uint) {
	fake.SizeStub = nil
	fake.sizeReturns = struct {
		result1 uint
	}{result1}
}

func (fake *FakeBridges) Contains(bridge models.Bridge) bool {
	fake.containsMutex.Lock()
	fake.containsArgsForCall = append(fake.containsArgsForCall, struct {
		bridge models.Bridge
	}{bridge})
	fake.recordInvocation("Contains", []interface{}{bridge})
	fake.containsMutex.Unlock()
	if fake.ContainsStub != nil {
		return fake.ContainsStub(bridge)
	} else {
		return fake.containsReturns.result1
	}
}

func (fake *FakeBridges) ContainsCallCount() int {
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	return len(fake.containsArgsForCall)
}

func (fake *FakeBridges) ContainsArgsForCall(i int) models.Bridge {
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	return fake.containsArgsForCall[i].bridge
}

func (fake *FakeBridges) ContainsReturns(result1 bool) {
	fake.ContainsStub = nil
	fake.containsReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeBridges) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.removeAndCloseAllMutex.RLock()
	defer fake.removeAndCloseAllMutex.RUnlock()
	fake.sizeMutex.RLock()
	defer fake.sizeMutex.RUnlock()
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBridges) recordInvocation(key string, args []interface{}) {
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

var _ models.Bridges = new(FakeBridges)