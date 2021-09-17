// Code generated by mockery 2.9.0. DO NOT EDIT.

package repository

import mock "github.com/stretchr/testify/mock"

// MockInventoryRepository is an autogenerated mock type for the InventoryRepository type
type MockInventoryRepository struct {
	mock.Mock
}

// GetByName provides a mock function with given fields: _a0
func (_m *MockInventoryRepository) GetByName(_a0 string) Inventory {
	ret := _m.Called(_a0)

	var r0 Inventory
	if rf, ok := ret.Get(0).(func(string) Inventory); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(Inventory)
	}

	return r0
}

// UpdateQty provides a mock function with given fields: _a0, _a1
func (_m *MockInventoryRepository) UpdateQty(_a0 string, _a1 int) {
	_m.Called(_a0, _a1)
}