// This file was generated by counterfeiter
package fakes

import (
	"net/url"
	"sync"

	"github.com/cloudfoundry/go-ccapi/v3/client"
)

type FakeClient struct {
	TokensUpdatedStub        func() bool
	tokensUpdatedMutex       sync.RWMutex
	tokensUpdatedArgsForCall []struct{}
	tokensUpdatedReturns     struct {
		result1 bool
	}
	GetUpdatedTokensStub        func() (string, string)
	getUpdatedTokensMutex       sync.RWMutex
	getUpdatedTokensArgsForCall []struct{}
	getUpdatedTokensReturns     struct {
		result1 string
		result2 string
	}
	GetApplicationsStub        func(queryParams url.Values) ([]byte, error)
	getApplicationsMutex       sync.RWMutex
	getApplicationsArgsForCall []struct {
		queryParams url.Values
	}
	getApplicationsReturns struct {
		result1 []byte
		result2 error
	}
	GetResourceStub        func(path string) ([]byte, error)
	getResourceMutex       sync.RWMutex
	getResourceArgsForCall []struct {
		path string
	}
	getResourceReturns struct {
		result1 []byte
		result2 error
	}
	GetResourcesStub        func(path string, limit int) ([]byte, error)
	getResourcesMutex       sync.RWMutex
	getResourcesArgsForCall []struct {
		path  string
		limit int
	}
	getResourcesReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeClient) TokensUpdated() bool {
	fake.tokensUpdatedMutex.Lock()
	fake.tokensUpdatedArgsForCall = append(fake.tokensUpdatedArgsForCall, struct{}{})
	fake.tokensUpdatedMutex.Unlock()
	if fake.TokensUpdatedStub != nil {
		return fake.TokensUpdatedStub()
	} else {
		return fake.tokensUpdatedReturns.result1
	}
}

func (fake *FakeClient) TokensUpdatedCallCount() int {
	fake.tokensUpdatedMutex.RLock()
	defer fake.tokensUpdatedMutex.RUnlock()
	return len(fake.tokensUpdatedArgsForCall)
}

func (fake *FakeClient) TokensUpdatedReturns(result1 bool) {
	fake.TokensUpdatedStub = nil
	fake.tokensUpdatedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeClient) GetUpdatedTokens() (string, string) {
	fake.getUpdatedTokensMutex.Lock()
	fake.getUpdatedTokensArgsForCall = append(fake.getUpdatedTokensArgsForCall, struct{}{})
	fake.getUpdatedTokensMutex.Unlock()
	if fake.GetUpdatedTokensStub != nil {
		return fake.GetUpdatedTokensStub()
	} else {
		return fake.getUpdatedTokensReturns.result1, fake.getUpdatedTokensReturns.result2
	}
}

func (fake *FakeClient) GetUpdatedTokensCallCount() int {
	fake.getUpdatedTokensMutex.RLock()
	defer fake.getUpdatedTokensMutex.RUnlock()
	return len(fake.getUpdatedTokensArgsForCall)
}

func (fake *FakeClient) GetUpdatedTokensReturns(result1 string, result2 string) {
	fake.GetUpdatedTokensStub = nil
	fake.getUpdatedTokensReturns = struct {
		result1 string
		result2 string
	}{result1, result2}
}

func (fake *FakeClient) GetApplications(queryParams url.Values) ([]byte, error) {
	fake.getApplicationsMutex.Lock()
	fake.getApplicationsArgsForCall = append(fake.getApplicationsArgsForCall, struct {
		queryParams url.Values
	}{queryParams})
	fake.getApplicationsMutex.Unlock()
	if fake.GetApplicationsStub != nil {
		return fake.GetApplicationsStub(queryParams)
	} else {
		return fake.getApplicationsReturns.result1, fake.getApplicationsReturns.result2
	}
}

func (fake *FakeClient) GetApplicationsCallCount() int {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return len(fake.getApplicationsArgsForCall)
}

func (fake *FakeClient) GetApplicationsArgsForCall(i int) url.Values {
	fake.getApplicationsMutex.RLock()
	defer fake.getApplicationsMutex.RUnlock()
	return fake.getApplicationsArgsForCall[i].queryParams
}

func (fake *FakeClient) GetApplicationsReturns(result1 []byte, result2 error) {
	fake.GetApplicationsStub = nil
	fake.getApplicationsReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetResource(path string) ([]byte, error) {
	fake.getResourceMutex.Lock()
	fake.getResourceArgsForCall = append(fake.getResourceArgsForCall, struct {
		path string
	}{path})
	fake.getResourceMutex.Unlock()
	if fake.GetResourceStub != nil {
		return fake.GetResourceStub(path)
	} else {
		return fake.getResourceReturns.result1, fake.getResourceReturns.result2
	}
}

func (fake *FakeClient) GetResourceCallCount() int {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return len(fake.getResourceArgsForCall)
}

func (fake *FakeClient) GetResourceArgsForCall(i int) string {
	fake.getResourceMutex.RLock()
	defer fake.getResourceMutex.RUnlock()
	return fake.getResourceArgsForCall[i].path
}

func (fake *FakeClient) GetResourceReturns(result1 []byte, result2 error) {
	fake.GetResourceStub = nil
	fake.getResourceReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetResources(path string, limit int) ([]byte, error) {
	fake.getResourcesMutex.Lock()
	fake.getResourcesArgsForCall = append(fake.getResourcesArgsForCall, struct {
		path  string
		limit int
	}{path, limit})
	fake.getResourcesMutex.Unlock()
	if fake.GetResourcesStub != nil {
		return fake.GetResourcesStub(path, limit)
	} else {
		return fake.getResourcesReturns.result1, fake.getResourcesReturns.result2
	}
}

func (fake *FakeClient) GetResourcesCallCount() int {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return len(fake.getResourcesArgsForCall)
}

func (fake *FakeClient) GetResourcesArgsForCall(i int) (string, int) {
	fake.getResourcesMutex.RLock()
	defer fake.getResourcesMutex.RUnlock()
	return fake.getResourcesArgsForCall[i].path, fake.getResourcesArgsForCall[i].limit
}

func (fake *FakeClient) GetResourcesReturns(result1 []byte, result2 error) {
	fake.GetResourcesStub = nil
	fake.getResourcesReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ client.Client = new(FakeClient)
