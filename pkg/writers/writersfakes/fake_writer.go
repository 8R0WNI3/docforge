// Code generated by counterfeiter. DO NOT EDIT.
package writersfakes

import (
	"sync"

	"github.com/gardener/docforge/pkg/api"
	"github.com/gardener/docforge/pkg/writers"
)

type FakeWriter struct {
	WriteStub        func(string, string, []byte, *api.Node) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 *api.Node
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWriter) Write(arg1 string, arg2 string, arg3 []byte, arg4 *api.Node) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 *api.Node
	}{arg1, arg2, arg3Copy, arg4})
	stub := fake.WriteStub
	fakeReturns := fake.writeReturns
	fake.recordInvocation("Write", []interface{}{arg1, arg2, arg3Copy, arg4})
	fake.writeMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWriter) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeWriter) WriteCalls(stub func(string, string, []byte, *api.Node) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeWriter) WriteArgsForCall(i int) (string, string, []byte, *api.Node) {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeWriter) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWriter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWriter) recordInvocation(key string, args []interface{}) {
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

var _ writers.Writer = new(FakeWriter)
