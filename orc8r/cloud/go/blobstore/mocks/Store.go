// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	blobstore "magma/orc8r/cloud/go/blobstore"

	mock "github.com/stretchr/testify/mock"

	storage "magma/orc8r/cloud/go/storage"
)

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Commit provides a mock function with given fields:
func (_m *Store) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: networkID, ids
func (_m *Store) Delete(networkID string, ids storage.TKs) error {
	ret := _m.Called(networkID, ids)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, storage.TKs) error); ok {
		r0 = rf(networkID, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: networkID, id
func (_m *Store) Get(networkID string, id storage.TK) (blobstore.Blob, error) {
	ret := _m.Called(networkID, id)

	var r0 blobstore.Blob
	if rf, ok := ret.Get(0).(func(string, storage.TK) blobstore.Blob); ok {
		r0 = rf(networkID, id)
	} else {
		r0 = ret.Get(0).(blobstore.Blob)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, storage.TK) error); ok {
		r1 = rf(networkID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetExistingKeys provides a mock function with given fields: keys, filter
func (_m *Store) GetExistingKeys(keys []string, filter blobstore.SearchFilter) ([]string, error) {
	ret := _m.Called(keys, filter)

	var r0 []string
	if rf, ok := ret.Get(0).(func([]string, blobstore.SearchFilter) []string); ok {
		r0 = rf(keys, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]string, blobstore.SearchFilter) error); ok {
		r1 = rf(keys, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMany provides a mock function with given fields: networkID, ids
func (_m *Store) GetMany(networkID string, ids storage.TKs) (blobstore.Blobs, error) {
	ret := _m.Called(networkID, ids)

	var r0 blobstore.Blobs
	if rf, ok := ret.Get(0).(func(string, storage.TKs) blobstore.Blobs); ok {
		r0 = rf(networkID, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blobstore.Blobs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, storage.TKs) error); ok {
		r1 = rf(networkID, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementVersion provides a mock function with given fields: networkID, id
func (_m *Store) IncrementVersion(networkID string, id storage.TK) error {
	ret := _m.Called(networkID, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, storage.TK) error); ok {
		r0 = rf(networkID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Rollback provides a mock function with given fields:
func (_m *Store) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Search provides a mock function with given fields: filter, criteria
func (_m *Store) Search(filter blobstore.SearchFilter, criteria blobstore.LoadCriteria) (map[string]blobstore.Blobs, error) {
	ret := _m.Called(filter, criteria)

	var r0 map[string]blobstore.Blobs
	if rf, ok := ret.Get(0).(func(blobstore.SearchFilter, blobstore.LoadCriteria) map[string]blobstore.Blobs); ok {
		r0 = rf(filter, criteria)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]blobstore.Blobs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(blobstore.SearchFilter, blobstore.LoadCriteria) error); ok {
		r1 = rf(filter, criteria)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Write provides a mock function with given fields: networkID, blobs
func (_m *Store) Write(networkID string, blobs blobstore.Blobs) error {
	ret := _m.Called(networkID, blobs)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, blobstore.Blobs) error); ok {
		r0 = rf(networkID, blobs)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
