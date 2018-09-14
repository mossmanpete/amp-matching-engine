// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "gopkg.in/mgo.v2/bson"
import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/Proofsuite/amp-matching-engine/types"

// TradeDao is an autogenerated mock type for the TradeDao type
type TradeDao struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: q
func (_m *TradeDao) Aggregate(q []bson.M) ([]*types.Tick, error) {
	ret := _m.Called(q)

	var r0 []*types.Tick
	if rf, ok := ret.Get(0).(func([]bson.M) []*types.Tick); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Tick)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]bson.M) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: o
func (_m *TradeDao) Create(o ...*types.Trade) error {
	_va := make([]interface{}, len(o))
	for _i := range o {
		_va[_i] = o[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*types.Trade) error); ok {
		r0 = rf(o...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Drop provides a mock function with given fields:
func (_m *TradeDao) Drop() {
	_m.Called()
}

// GetAll provides a mock function with given fields:
func (_m *TradeDao) GetAll() ([]types.Trade, error) {
	ret := _m.Called()

	var r0 []types.Trade
	if rf, ok := ret.Get(0).(func() []types.Trade); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByHash provides a mock function with given fields: hash
func (_m *TradeDao) GetByHash(hash common.Hash) (*types.Trade, error) {
	ret := _m.Called(hash)

	var r0 *types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) *types.Trade); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByOrderHash provides a mock function with given fields: hash
func (_m *TradeDao) GetByOrderHash(hash common.Hash) ([]*types.Trade, error) {
	ret := _m.Called(hash)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Hash) []*types.Trade); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPairAddress provides a mock function with given fields: baseToken, quoteToken
func (_m *TradeDao) GetByPairAddress(baseToken common.Address, quoteToken common.Address) ([]*types.Trade, error) {
	ret := _m.Called(baseToken, quoteToken)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address, common.Address) []*types.Trade); ok {
		r0 = rf(baseToken, quoteToken)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address, common.Address) error); ok {
		r1 = rf(baseToken, quoteToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPairName provides a mock function with given fields: name
func (_m *TradeDao) GetByPairName(name string) ([]*types.Trade, error) {
	ret := _m.Called(name)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(string) []*types.Trade); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUserAddress provides a mock function with given fields: addr
func (_m *TradeDao) GetByUserAddress(addr common.Address) ([]*types.Trade, error) {
	ret := _m.Called(addr)

	var r0 []*types.Trade
	if rf, ok := ret.Get(0).(func(common.Address) []*types.Trade); ok {
		r0 = rf(addr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Trade)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(addr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: t
func (_m *TradeDao) Update(t *types.Trade) error {
	ret := _m.Called(t)

	var r0 error
	if rf, ok := ret.Get(0).(func(*types.Trade) error); ok {
		r0 = rf(t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByHash provides a mock function with given fields: hash, t
func (_m *TradeDao) UpdateByHash(hash common.Hash, t *types.Trade) error {
	ret := _m.Called(hash, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, *types.Trade) error); ok {
		r0 = rf(hash, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateTradeStatus provides a mock function with given fields: hash, status
func (_m *TradeDao) UpdateTradeStatus(hash common.Hash, status string) error {
	ret := _m.Called(hash, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.Hash, string) error); ok {
		r0 = rf(hash, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
