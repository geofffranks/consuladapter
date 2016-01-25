// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/consuladapter"
	"github.com/hashicorp/consul/api"
)

type FakeKV struct {
	GetStub        func(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		key string
		q   *api.QueryOptions
	}
	getReturns struct {
		result1 *api.KVPair
		result2 *api.QueryMeta
		result3 error
	}
	ListStub        func(prefix string, q *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		prefix string
		q      *api.QueryOptions
	}
	listReturns struct {
		result1 api.KVPairs
		result2 *api.QueryMeta
		result3 error
	}
	PutStub        func(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		p *api.KVPair
		q *api.WriteOptions
	}
	putReturns struct {
		result1 *api.WriteMeta
		result2 error
	}
}

func (fake *FakeKV) Get(key string, q *api.QueryOptions) (*api.KVPair, *api.QueryMeta, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		key string
		q   *api.QueryOptions
	}{key, q})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(key, q)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2, fake.getReturns.result3
	}
}

func (fake *FakeKV) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeKV) GetArgsForCall(i int) (string, *api.QueryOptions) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].key, fake.getArgsForCall[i].q
}

func (fake *FakeKV) GetReturns(result1 *api.KVPair, result2 *api.QueryMeta, result3 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *api.KVPair
		result2 *api.QueryMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeKV) List(prefix string, q *api.QueryOptions) (api.KVPairs, *api.QueryMeta, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		prefix string
		q      *api.QueryOptions
	}{prefix, q})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(prefix, q)
	} else {
		return fake.listReturns.result1, fake.listReturns.result2, fake.listReturns.result3
	}
}

func (fake *FakeKV) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeKV) ListArgsForCall(i int) (string, *api.QueryOptions) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return fake.listArgsForCall[i].prefix, fake.listArgsForCall[i].q
}

func (fake *FakeKV) ListReturns(result1 api.KVPairs, result2 *api.QueryMeta, result3 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 api.KVPairs
		result2 *api.QueryMeta
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeKV) Put(p *api.KVPair, q *api.WriteOptions) (*api.WriteMeta, error) {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		p *api.KVPair
		q *api.WriteOptions
	}{p, q})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(p, q)
	} else {
		return fake.putReturns.result1, fake.putReturns.result2
	}
}

func (fake *FakeKV) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeKV) PutArgsForCall(i int) (*api.KVPair, *api.WriteOptions) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].p, fake.putArgsForCall[i].q
}

func (fake *FakeKV) PutReturns(result1 *api.WriteMeta, result2 error) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 *api.WriteMeta
		result2 error
	}{result1, result2}
}

var _ consuladapter.KV = new(FakeKV)
