package establishment_test

import (
	"errors"
	"store-manager-api/app/modules/establishment"
	mocks "store-manager-api/app/modules/establishment/mocks"
	"testing"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	expected := []establishment.Establishment{
		{ID: 1, Name: "Est 1"},
		{ID: 2, Name: "Est 2"},
	}

	repo.EXPECT().FindAll().Return(expected, nil)
	service := establishment.NewService(repo)

	result, err := service.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	expected := &establishment.Establishment{ID: 1, Name: "Est 1"}

	repo.EXPECT().FindByID(1).Return(expected, nil)
	service := establishment.NewService(repo)

	result, err := service.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestService_GetByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	repo.EXPECT().FindByID(999).Return(nil, nil)
	service := establishment.NewService(repo)

	result, err := service.GetByID(999)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, establishment.ErrNotFound, err)
}

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	est := &establishment.Establishment{Name: "New Est"}

	repo.EXPECT().Create(est).Return(nil)
	service := establishment.NewService(repo)

	err := service.Create(est)
	assert.NoError(t, err)
}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	est := &establishment.Establishment{ID: 1, Name: "Updated"}

	repo.EXPECT().Update(est).Return(nil)
	service := establishment.NewService(repo)

	err := service.Update(est)
	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	repo.EXPECT().HasStores(1).Return(false, nil)
	repo.EXPECT().Delete(1).Return(nil)
	service := establishment.NewService(repo)

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestService_Delete_HasStores(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	repo.EXPECT().HasStores(1).Return(true, nil)
	service := establishment.NewService(repo)

	err := service.Delete(1)
	assert.Error(t, err)
	assert.Equal(t, establishment.ErrForeignKeyExists, err)
}

func TestService_Delete_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockEstablishmentRepository(ctrl)
	repo.EXPECT().HasStores(1).Return(false, nil)
	repo.EXPECT().Delete(1).Return(errors.New("delete failed"))
	service := establishment.NewService(repo)

	err := service.Delete(1)
	assert.Error(t, err)
	assert.Equal(t, "delete failed", err.Error())
}