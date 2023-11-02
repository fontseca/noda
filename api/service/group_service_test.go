package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"noda/api/data/model"
	"noda/api/data/transfer"
	"noda/api/data/types"
	"testing"
)

type groupRepositoryMock struct {
	mock.Mock
}

func (o *groupRepositoryMock) InsertGroup(ownerID string, next *transfer.GroupCreation) (string, error) {
	args := o.Called(ownerID, next)
	return args.String(0), args.Error(1)
}

func (o *groupRepositoryMock) FetchGroupByID(ownerID, groupID string) (*model.Group, error) {
	args := o.Called(ownerID, groupID)
	var group *model.Group
	arg1 := args.Get(0)
	if nil != arg1 {
		group = arg1.(*model.Group)
	}
	return group, args.Error(1)
}

func (o *groupRepositoryMock) FetchGroups(ownerID string, page, rpp int64, needle, sortBy string) ([]*model.Group, error) {
	args := o.Called(ownerID, page, rpp, needle, sortBy)
	var groups []*model.Group
	arg1 := args.Get(0)
	if nil != arg1 {
		groups = arg1.([]*model.Group)
	}
	return groups, args.Error(1)
}

func (o *groupRepositoryMock) UpdateGroup(ownerID, groupID string, up *transfer.GroupUpdate) (ok bool, err error) {
	args := o.Called(ownerID, groupID, up)
	return args.Bool(0), args.Error(1)
}

func (o *groupRepositoryMock) DeleteGroup(ownerID, groupID string) (ok bool, err error) {
	args := o.Called(ownerID, groupID)
	return args.Bool(0), args.Error(1)
}

func TestGroupService_SaveGroup(t *testing.T) {
	var (
		m       *groupRepositoryMock
		ownerID = uuid.New()
		next    = new(transfer.GroupCreation)
		s       *GroupService
		res     string
		err     error
	)

	/* Success.  */

	t.Run("success", func(t *testing.T) {
		m = new(groupRepositoryMock)
		m.On("InsertGroup", ownerID.String(), next).
			Return(ownerID.String(), nil)
		s = NewGroupService(m)
		res, err = s.SaveGroup(ownerID, next)
		assert.Equal(t, ownerID.String(), res)
		assert.NoError(t, err)
	})

	/* Got an error.  */

	t.Run("got an error", func(t *testing.T) {
		unexpected := errors.New("unexpected error")
		m = new(groupRepositoryMock)
		m.On("InsertGroup", ownerID.String(), next).
			Return("", unexpected)
		s = NewGroupService(m)
		res, err = s.SaveGroup(ownerID, next)
		assert.Empty(t, res)
		assert.ErrorIs(t, err, unexpected)
	})
}

func TestGroupService_FindGroupByID(t *testing.T) {
	var (
		m                *groupRepositoryMock
		ownerID, groupID = uuid.New(), uuid.New()
		s                *GroupService
		res              *model.Group
		err              error
	)

	/* Success.  */

	t.Run("success", func(t *testing.T) {
		current := new(model.Group)
		m = new(groupRepositoryMock)
		m.On("FetchGroupByID", ownerID.String(), groupID.String()).
			Return(current, nil)
		s = NewGroupService(m)
		res, err = s.FindGroupByID(ownerID, groupID)
		assert.Equal(t, current, res)
		assert.NoError(t, err)
	})

	/* Got an error.  */

	t.Run("got an error", func(t *testing.T) {
		unexpected := errors.New("unexpected error")
		m = new(groupRepositoryMock)
		m.On("FetchGroupByID", ownerID.String(), groupID.String()).
			Return(nil, unexpected)
		s = NewGroupService(m)
		res, err = s.FindGroupByID(ownerID, groupID)
		assert.Nil(t, res)
		assert.ErrorIs(t, err, unexpected)
	})
}

func TestGroupService_FindGroups(t *testing.T) {
	var (
		m       *groupRepositoryMock
		ownerID = uuid.New()
		s       *GroupService
		err     error
		res     *types.Result[model.Group]
		pag     = &types.Pagination{Page: 1, RPP: 10}
	)

	/* Success.  */

	t.Run("success", func(t *testing.T) {
		var groups = make([]*model.Group, 0)
		current := &types.Result[model.Group]{
			Page:      1,
			RPP:       10,
			Payload:   groups,
			Retrieved: int64(len(groups)),
		}
		m = new(groupRepositoryMock)
		m.On("FetchGroups", ownerID.String(), pag.Page, pag.RPP, "", "").
			Return(groups, nil)
		s = NewGroupService(m)
		res, err = s.FindGroups(ownerID, pag, "", "")
		assert.Equal(t, current, res)
		assert.NoError(t, err)
	})

	/* Got an error.  */

	t.Run("got an error", func(t *testing.T) {
		unexpected := errors.New("unexpected error")
		m = new(groupRepositoryMock)
		m.On("FetchGroups", ownerID.String(), pag.Page, pag.RPP, "", "").
			Return(nil, unexpected)
		s = NewGroupService(m)
		res, err = s.FindGroups(ownerID, pag, "", "")
		assert.Nil(t, res)
		assert.ErrorIs(t, err, unexpected)
	})
}

func TestGroupService_UpdateGroup(t *testing.T) {
	var (
		m                *groupRepositoryMock
		ownerID, groupID = uuid.New(), uuid.New()
		s                *GroupService
		res              bool
		err              error
		up               = new(transfer.GroupUpdate)
	)

	/* Success.  */

	t.Run("success", func(t *testing.T) {
		m = new(groupRepositoryMock)
		m.On("UpdateGroup", ownerID.String(), groupID.String(), up).
			Return(true, nil)
		s = NewGroupService(m)
		res, err = s.UpdateGroup(ownerID, groupID, up)
		assert.True(t, res)
		assert.NoError(t, err)
	})

	/* Got an error.  */

	t.Run("got an error", func(t *testing.T) {
		unexpected := errors.New("unexpected error")
		m = new(groupRepositoryMock)
		m.On("UpdateGroup", ownerID.String(), groupID.String(), up).
			Return(false, unexpected)
		s = NewGroupService(m)
		res, err = s.UpdateGroup(ownerID, groupID, up)
		assert.False(t, res)
		assert.ErrorIs(t, err, unexpected)
	})
}

func TestGroupService_DeleteGroup(t *testing.T) {
	var (
		m                *groupRepositoryMock
		ownerID, groupID = uuid.New(), uuid.New()
		s                *GroupService
		res              bool
		err              error
	)

	/* Success.  */

	t.Run("success", func(t *testing.T) {
		m = new(groupRepositoryMock)
		m.On("DeleteGroup", ownerID.String(), groupID.String()).
			Return(true, nil)
		s = NewGroupService(m)
		res, err = s.DeleteGroup(ownerID, groupID)
		assert.True(t, res)
		assert.NoError(t, err)
	})

	/* Got an error.  */

	t.Run("got an error", func(t *testing.T) {
		unexpected := errors.New("unexpected error")
		m = new(groupRepositoryMock)
		m.On("DeleteGroup", ownerID.String(), groupID.String()).
			Return(false, unexpected)
		s = NewGroupService(m)
		res, err = s.DeleteGroup(ownerID, groupID)
		assert.False(t, res)
		assert.ErrorIs(t, err, unexpected)
	})
}
