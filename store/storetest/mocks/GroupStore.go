// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// GroupStore is an autogenerated mock type for the GroupStore type
type GroupStore struct {
	mock.Mock
}

// Create provides a mock function with given fields: group
func (_m *GroupStore) Create(group *model.Group) store.StoreChannel {
	ret := _m.Called(group)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Group) store.StoreChannel); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// CreateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) CreateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	ret := _m.Called(groupSyncable)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) store.StoreChannel); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// CreateMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) CreateMember(groupID string, userID string) store.StoreChannel {
	ret := _m.Called(groupID, userID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: groupID
func (_m *GroupStore) Delete(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) DeleteGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) store.StoreChannel); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// DeleteMember provides a mock function with given fields: groupID, userID
func (_m *GroupStore) DeleteMember(groupID string, userID string) store.StoreChannel {
	ret := _m.Called(groupID, userID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string) store.StoreChannel); ok {
		r0 = rf(groupID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Get provides a mock function with given fields: groupID
func (_m *GroupStore) Get(groupID string) store.StoreChannel {
	ret := _m.Called(groupID)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string) store.StoreChannel); ok {
		r0 = rf(groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllGroupSyncablesByGroupPage provides a mock function with given fields: groupID, syncableType, offset, limit
func (_m *GroupStore) GetAllGroupSyncablesByGroupPage(groupID string, syncableType model.GroupSyncableType, offset int, limit int) store.StoreChannel {
	ret := _m.Called(groupID, syncableType, offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, model.GroupSyncableType, int, int) store.StoreChannel); ok {
		r0 = rf(groupID, syncableType, offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetAllPage provides a mock function with given fields: offset, limit
func (_m *GroupStore) GetAllPage(offset int, limit int) store.StoreChannel {
	ret := _m.Called(offset, limit)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(int, int) store.StoreChannel); ok {
		r0 = rf(offset, limit)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// GetGroupSyncable provides a mock function with given fields: groupID, syncableID, syncableType
func (_m *GroupStore) GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) store.StoreChannel {
	ret := _m.Called(groupID, syncableID, syncableType)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(string, string, model.GroupSyncableType) store.StoreChannel); ok {
		r0 = rf(groupID, syncableID, syncableType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// Update provides a mock function with given fields: group
func (_m *GroupStore) Update(group *model.Group) store.StoreChannel {
	ret := _m.Called(group)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.Group) store.StoreChannel); ok {
		r0 = rf(group)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}

// UpdateGroupSyncable provides a mock function with given fields: groupSyncable
func (_m *GroupStore) UpdateGroupSyncable(groupSyncable *model.GroupSyncable) store.StoreChannel {
	ret := _m.Called(groupSyncable)

	var r0 store.StoreChannel
	if rf, ok := ret.Get(0).(func(*model.GroupSyncable) store.StoreChannel); ok {
		r0 = rf(groupSyncable)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StoreChannel)
		}
	}

	return r0
}
