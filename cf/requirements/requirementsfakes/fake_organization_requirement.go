// This file was generated by counterfeiter
package requirementsfakes

import (
	"sync"

	"github.com/cloudfoundry/cli/cf/models"
	"github.com/cloudfoundry/cli/cf/requirements"
)

type FakeOrganizationRequirement struct {
	ExecuteStub        func() error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct{}
	executeReturns     struct {
		result1 error
	}
	SetOrganizationNameStub        func(string)
	setOrganizationNameMutex       sync.RWMutex
	setOrganizationNameArgsForCall []struct {
		arg1 string
	}
	GetOrganizationStub        func() models.Organization
	getOrganizationMutex       sync.RWMutex
	getOrganizationArgsForCall []struct{}
	getOrganizationReturns     struct {
		result1 models.Organization
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOrganizationRequirement) Execute() error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct{}{})
	fake.recordInvocation("Execute", []interface{}{})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub()
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *FakeOrganizationRequirement) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *FakeOrganizationRequirement) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOrganizationRequirement) SetOrganizationName(arg1 string) {
	fake.setOrganizationNameMutex.Lock()
	fake.setOrganizationNameArgsForCall = append(fake.setOrganizationNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("SetOrganizationName", []interface{}{arg1})
	fake.setOrganizationNameMutex.Unlock()
	if fake.SetOrganizationNameStub != nil {
		fake.SetOrganizationNameStub(arg1)
	}
}

func (fake *FakeOrganizationRequirement) SetOrganizationNameCallCount() int {
	fake.setOrganizationNameMutex.RLock()
	defer fake.setOrganizationNameMutex.RUnlock()
	return len(fake.setOrganizationNameArgsForCall)
}

func (fake *FakeOrganizationRequirement) SetOrganizationNameArgsForCall(i int) string {
	fake.setOrganizationNameMutex.RLock()
	defer fake.setOrganizationNameMutex.RUnlock()
	return fake.setOrganizationNameArgsForCall[i].arg1
}

func (fake *FakeOrganizationRequirement) GetOrganization() models.Organization {
	fake.getOrganizationMutex.Lock()
	fake.getOrganizationArgsForCall = append(fake.getOrganizationArgsForCall, struct{}{})
	fake.recordInvocation("GetOrganization", []interface{}{})
	fake.getOrganizationMutex.Unlock()
	if fake.GetOrganizationStub != nil {
		return fake.GetOrganizationStub()
	} else {
		return fake.getOrganizationReturns.result1
	}
}

func (fake *FakeOrganizationRequirement) GetOrganizationCallCount() int {
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return len(fake.getOrganizationArgsForCall)
}

func (fake *FakeOrganizationRequirement) GetOrganizationReturns(result1 models.Organization) {
	fake.GetOrganizationStub = nil
	fake.getOrganizationReturns = struct {
		result1 models.Organization
	}{result1}
}

func (fake *FakeOrganizationRequirement) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	fake.setOrganizationNameMutex.RLock()
	defer fake.setOrganizationNameMutex.RUnlock()
	fake.getOrganizationMutex.RLock()
	defer fake.getOrganizationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeOrganizationRequirement) recordInvocation(key string, args []interface{}) {
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

var _ requirements.OrganizationRequirement = new(FakeOrganizationRequirement)
