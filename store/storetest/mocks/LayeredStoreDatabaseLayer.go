// Code generated by mockery v1.0.0. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mattermost/mattermost-server/model"
import store "github.com/mattermost/mattermost-server/store"

// LayeredStoreDatabaseLayer is an autogenerated mock type for the LayeredStoreDatabaseLayer type
type LayeredStoreDatabaseLayer struct {
	mock.Mock
}

// Audit provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Audit() store.AuditStore {
	ret := _m.Called()

	var r0 store.AuditStore
	if rf, ok := ret.Get(0).(func() store.AuditStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.AuditStore)
		}
	}

	return r0
}

// Channel provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Channel() store.ChannelStore {
	ret := _m.Called()

	var r0 store.ChannelStore
	if rf, ok := ret.Get(0).(func() store.ChannelStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ChannelStore)
		}
	}

	return r0
}

// ChannelMemberHistory provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) ChannelMemberHistory() store.ChannelMemberHistoryStore {
	ret := _m.Called()

	var r0 store.ChannelMemberHistoryStore
	if rf, ok := ret.Get(0).(func() store.ChannelMemberHistoryStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ChannelMemberHistoryStore)
		}
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Close() {
	_m.Called()
}

// ClusterDiscovery provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) ClusterDiscovery() store.ClusterDiscoveryStore {
	ret := _m.Called()

	var r0 store.ClusterDiscoveryStore
	if rf, ok := ret.Get(0).(func() store.ClusterDiscoveryStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ClusterDiscoveryStore)
		}
	}

	return r0
}

// Command provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Command() store.CommandStore {
	ret := _m.Called()

	var r0 store.CommandStore
	if rf, ok := ret.Get(0).(func() store.CommandStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.CommandStore)
		}
	}

	return r0
}

// CommandWebhook provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) CommandWebhook() store.CommandWebhookStore {
	ret := _m.Called()

	var r0 store.CommandWebhookStore
	if rf, ok := ret.Get(0).(func() store.CommandWebhookStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.CommandWebhookStore)
		}
	}

	return r0
}

// Compliance provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Compliance() store.ComplianceStore {
	ret := _m.Called()

	var r0 store.ComplianceStore
	if rf, ok := ret.Get(0).(func() store.ComplianceStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ComplianceStore)
		}
	}

	return r0
}

// DropAllTables provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) DropAllTables() {
	_m.Called()
}

// Emoji provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Emoji() store.EmojiStore {
	ret := _m.Called()

	var r0 store.EmojiStore
	if rf, ok := ret.Get(0).(func() store.EmojiStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.EmojiStore)
		}
	}

	return r0
}

// FileInfo provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) FileInfo() store.FileInfoStore {
	ret := _m.Called()

	var r0 store.FileInfoStore
	if rf, ok := ret.Get(0).(func() store.FileInfoStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.FileInfoStore)
		}
	}

	return r0
}

// Group provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Group() store.GroupStore {
	ret := _m.Called()

	var r0 store.GroupStore
	if rf, ok := ret.Get(0).(func() store.GroupStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.GroupStore)
		}
	}

	return r0
}

// GroupCreate provides a mock function with given fields: ctx, group, hints
func (_m *LayeredStoreDatabaseLayer) GroupCreate(ctx context.Context, group *model.Group, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, group)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Group, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, group, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupCreateGroupSyncable provides a mock function with given fields: ctx, groupSyncable, hints
func (_m *LayeredStoreDatabaseLayer) GroupCreateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupSyncable)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.GroupSyncable, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupSyncable, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupCreateMember provides a mock function with given fields: ctx, groupID, userID, hints
func (_m *LayeredStoreDatabaseLayer) GroupCreateMember(ctx context.Context, groupID string, userID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, userID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDelete provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreDatabaseLayer) GroupDelete(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDeleteGroupSyncable provides a mock function with given fields: ctx, groupID, syncableID, syncableType, hints
func (_m *LayeredStoreDatabaseLayer) GroupDeleteGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableID, syncableType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.GroupSyncableType, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableID, syncableType, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupDeleteMember provides a mock function with given fields: ctx, groupID, userID, hints
func (_m *LayeredStoreDatabaseLayer) GroupDeleteMember(ctx context.Context, groupID string, userID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, userID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGet provides a mock function with given fields: ctx, groupID, hints
func (_m *LayeredStoreDatabaseLayer) GroupGet(ctx context.Context, groupID string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetAllGroupSyncablesByGroupPage provides a mock function with given fields: ctx, groupID, syncableType, offset, limit, hints
func (_m *LayeredStoreDatabaseLayer) GroupGetAllGroupSyncablesByGroupPage(ctx context.Context, groupID string, syncableType model.GroupSyncableType, offset int, limit int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableType, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, model.GroupSyncableType, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableType, offset, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetAllPage provides a mock function with given fields: ctx, offset, limit, hints
func (_m *LayeredStoreDatabaseLayer) GroupGetAllPage(ctx context.Context, offset int, limit int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, offset, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupGetGroupSyncable provides a mock function with given fields: ctx, groupID, syncableID, syncableType, hints
func (_m *LayeredStoreDatabaseLayer) GroupGetGroupSyncable(ctx context.Context, groupID string, syncableID string, syncableType model.GroupSyncableType, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID, syncableID, syncableType)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, string, model.GroupSyncableType, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupID, syncableID, syncableType, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupUpdate provides a mock function with given fields: ctx, group, hints
func (_m *LayeredStoreDatabaseLayer) GroupUpdate(ctx context.Context, group *model.Group, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, group)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Group, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, group, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// GroupUpdateGroupSyncable provides a mock function with given fields: ctx, groupSyncable, hints
func (_m *LayeredStoreDatabaseLayer) GroupUpdateGroupSyncable(ctx context.Context, groupSyncable *model.GroupSyncable, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupSyncable)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.GroupSyncable, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, groupSyncable, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// Job provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Job() store.JobStore {
	ret := _m.Called()

	var r0 store.JobStore
	if rf, ok := ret.Get(0).(func() store.JobStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.JobStore)
		}
	}

	return r0
}

// License provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) License() store.LicenseStore {
	ret := _m.Called()

	var r0 store.LicenseStore
	if rf, ok := ret.Get(0).(func() store.LicenseStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LicenseStore)
		}
	}

	return r0
}

// LockToMaster provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) LockToMaster() {
	_m.Called()
}

// MarkSystemRanUnitTests provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) MarkSystemRanUnitTests() {
	_m.Called()
}

// Next provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Next() store.LayeredStoreSupplier {
	ret := _m.Called()

	var r0 store.LayeredStoreSupplier
	if rf, ok := ret.Get(0).(func() store.LayeredStoreSupplier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.LayeredStoreSupplier)
		}
	}

	return r0
}

// OAuth provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) OAuth() store.OAuthStore {
	ret := _m.Called()

	var r0 store.OAuthStore
	if rf, ok := ret.Get(0).(func() store.OAuthStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.OAuthStore)
		}
	}

	return r0
}

// Plugin provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Plugin() store.PluginStore {
	ret := _m.Called()

	var r0 store.PluginStore
	if rf, ok := ret.Get(0).(func() store.PluginStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PluginStore)
		}
	}

	return r0
}

// Post provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Post() store.PostStore {
	ret := _m.Called()

	var r0 store.PostStore
	if rf, ok := ret.Get(0).(func() store.PostStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PostStore)
		}
	}

	return r0
}

// Preference provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Preference() store.PreferenceStore {
	ret := _m.Called()

	var r0 store.PreferenceStore
	if rf, ok := ret.Get(0).(func() store.PreferenceStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.PreferenceStore)
		}
	}

	return r0
}

// Reaction provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Reaction() store.ReactionStore {
	ret := _m.Called()

	var r0 store.ReactionStore
	if rf, ok := ret.Get(0).(func() store.ReactionStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ReactionStore)
		}
	}

	return r0
}

// ReactionDelete provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreDatabaseLayer) ReactionDelete(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionDeleteAllWithEmojiName provides a mock function with given fields: ctx, emojiName, hints
func (_m *LayeredStoreDatabaseLayer) ReactionDeleteAllWithEmojiName(ctx context.Context, emojiName string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, emojiName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, emojiName, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionGetForPost provides a mock function with given fields: ctx, postId, hints
func (_m *LayeredStoreDatabaseLayer) ReactionGetForPost(ctx context.Context, postId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, postId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, postId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionPermanentDeleteBatch provides a mock function with given fields: ctx, endTime, limit, hints
func (_m *LayeredStoreDatabaseLayer) ReactionPermanentDeleteBatch(ctx context.Context, endTime int64, limit int64, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, endTime, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, endTime, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ReactionSave provides a mock function with given fields: ctx, reaction, hints
func (_m *LayeredStoreDatabaseLayer) ReactionSave(ctx context.Context, reaction *model.Reaction, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, reaction)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Reaction, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, reaction, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// Role provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Role() store.RoleStore {
	ret := _m.Called()

	var r0 store.RoleStore
	if rf, ok := ret.Get(0).(func() store.RoleStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.RoleStore)
		}
	}

	return r0
}

// RoleDelete provides a mock function with given fields: ctx, roldId, hints
func (_m *LayeredStoreDatabaseLayer) RoleDelete(ctx context.Context, roldId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, roldId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, roldId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGet provides a mock function with given fields: ctx, roleId, hints
func (_m *LayeredStoreDatabaseLayer) RoleGet(ctx context.Context, roleId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, roleId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, roleId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByName provides a mock function with given fields: ctx, name, hints
func (_m *LayeredStoreDatabaseLayer) RoleGetByName(ctx context.Context, name string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, name, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleGetByNames provides a mock function with given fields: ctx, names, hints
func (_m *LayeredStoreDatabaseLayer) RoleGetByNames(ctx context.Context, names []string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, names)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, []string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, names, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RolePermanentDeleteAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreDatabaseLayer) RolePermanentDeleteAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// RoleSave provides a mock function with given fields: ctx, role, hints
func (_m *LayeredStoreDatabaseLayer) RoleSave(ctx context.Context, role *model.Role, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, role)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Role, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, role, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// Scheme provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Scheme() store.SchemeStore {
	ret := _m.Called()

	var r0 store.SchemeStore
	if rf, ok := ret.Get(0).(func() store.SchemeStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.SchemeStore)
		}
	}

	return r0
}

// SchemeDelete provides a mock function with given fields: ctx, schemeId, hints
func (_m *LayeredStoreDatabaseLayer) SchemeDelete(ctx context.Context, schemeId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGet provides a mock function with given fields: ctx, schemeId, hints
func (_m *LayeredStoreDatabaseLayer) SchemeGet(ctx context.Context, schemeId string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeId)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeId, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGetAllPage provides a mock function with given fields: ctx, scope, offset, limit, hints
func (_m *LayeredStoreDatabaseLayer) SchemeGetAllPage(ctx context.Context, scope string, offset int, limit int, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, scope, offset, limit)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, int, int, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, scope, offset, limit, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeGetByName provides a mock function with given fields: ctx, schemeName, hints
func (_m *LayeredStoreDatabaseLayer) SchemeGetByName(ctx context.Context, schemeName string, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, schemeName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, string, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, schemeName, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemePermanentDeleteAll provides a mock function with given fields: ctx, hints
func (_m *LayeredStoreDatabaseLayer) SchemePermanentDeleteAll(ctx context.Context, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// SchemeSave provides a mock function with given fields: ctx, scheme, hints
func (_m *LayeredStoreDatabaseLayer) SchemeSave(ctx context.Context, scheme *model.Scheme, hints ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult {
	_va := make([]interface{}, len(hints))
	for _i := range hints {
		_va[_i] = hints[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, scheme)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *store.LayeredStoreSupplierResult
	if rf, ok := ret.Get(0).(func(context.Context, *model.Scheme, ...store.LayeredStoreHint) *store.LayeredStoreSupplierResult); ok {
		r0 = rf(ctx, scheme, hints...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*store.LayeredStoreSupplierResult)
		}
	}

	return r0
}

// ServiceTerms provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) ServiceTerms() store.ServiceTermsStore {
	ret := _m.Called()

	var r0 store.ServiceTermsStore
	if rf, ok := ret.Get(0).(func() store.ServiceTermsStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.ServiceTermsStore)
		}
	}

	return r0
}

// Session provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Session() store.SessionStore {
	ret := _m.Called()

	var r0 store.SessionStore
	if rf, ok := ret.Get(0).(func() store.SessionStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.SessionStore)
		}
	}

	return r0
}

// SetChainNext provides a mock function with given fields: _a0
func (_m *LayeredStoreDatabaseLayer) SetChainNext(_a0 store.LayeredStoreSupplier) {
	_m.Called(_a0)
}

// Status provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Status() store.StatusStore {
	ret := _m.Called()

	var r0 store.StatusStore
	if rf, ok := ret.Get(0).(func() store.StatusStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.StatusStore)
		}
	}

	return r0
}

// System provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) System() store.SystemStore {
	ret := _m.Called()

	var r0 store.SystemStore
	if rf, ok := ret.Get(0).(func() store.SystemStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.SystemStore)
		}
	}

	return r0
}

// Team provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Team() store.TeamStore {
	ret := _m.Called()

	var r0 store.TeamStore
	if rf, ok := ret.Get(0).(func() store.TeamStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.TeamStore)
		}
	}

	return r0
}

// Token provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Token() store.TokenStore {
	ret := _m.Called()

	var r0 store.TokenStore
	if rf, ok := ret.Get(0).(func() store.TokenStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.TokenStore)
		}
	}

	return r0
}

// TotalMasterDbConnections provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) TotalMasterDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// TotalReadDbConnections provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) TotalReadDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// TotalSearchDbConnections provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) TotalSearchDbConnections() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// UnlockFromMaster provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) UnlockFromMaster() {
	_m.Called()
}

// User provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) User() store.UserStore {
	ret := _m.Called()

	var r0 store.UserStore
	if rf, ok := ret.Get(0).(func() store.UserStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.UserStore)
		}
	}

	return r0
}

// UserAccessToken provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) UserAccessToken() store.UserAccessTokenStore {
	ret := _m.Called()

	var r0 store.UserAccessTokenStore
	if rf, ok := ret.Get(0).(func() store.UserAccessTokenStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.UserAccessTokenStore)
		}
	}

	return r0
}

// Webhook provides a mock function with given fields:
func (_m *LayeredStoreDatabaseLayer) Webhook() store.WebhookStore {
	ret := _m.Called()

	var r0 store.WebhookStore
	if rf, ok := ret.Get(0).(func() store.WebhookStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(store.WebhookStore)
		}
	}

	return r0
}
