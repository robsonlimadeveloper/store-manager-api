package store_test

import (
    "errors"
    "store-manager-api/app/modules/store"
    mocks "store-manager-api/app/modules/store/mocks"
    "testing"
    "go.uber.org/mock/gomock"
    "github.com/stretchr/testify/assert"
)

// TestService_GetAll tests the GetAll method of the StoreService.
// It verifies that the method returns the expected list of stores without errors.

func TestService_GetAll(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockStoreRepository(ctrl)

    expectedStores := []store.Store{
        {ID: 1, Name: "Loja 1"},
        {ID: 2, Name: "Loja 2"},
    }

    mockRepo.EXPECT().
        FindAll().
        Return(expectedStores, nil)

    service := store.NewService(mockRepo)

    result, err := service.GetAll()
    assert.NoError(t, err)
    assert.Equal(t, expectedStores, result)
}

// TestService_GetAll_Error tests the GetAll method of the StoreService.
// It verifies that the method returns an error when the repository fails to retrieve stores.
func TestService_GetAll_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockStoreRepository(ctrl)
    mockRepo.EXPECT().
        FindAll().
        Return(nil, errors.New("db error"))

    service := store.NewService(mockRepo)

    result, err := service.GetAll()
    assert.Error(t, err)
    assert.Nil(t, result)
}

func TestService_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	expectedStore := &store.Store{ID: 1, Name: "Loja 1"}

	mockRepo.EXPECT().
		FindByID(1).
		Return(expectedStore, nil)

	service := store.NewService(mockRepo)

	result, err := service.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedStore, result)
}

func TestService_GetByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	mockRepo.EXPECT().
		FindByID(1).
		Return(nil, nil)

	service := store.NewService(mockRepo)

	result, err := service.GetByID(1)
	assert.ErrorIs(t, err, store.ErrNotFound)
	assert.Nil(t, result)
}

func TestService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	newStore := &store.Store{Name: "Loja Nova"}

	mockRepo.EXPECT().
		Create(newStore).
		Return(nil)

	service := store.NewService(mockRepo)

	err := service.Create(newStore)
	assert.NoError(t, err)
}

func TestService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	updatedStore := &store.Store{ID: 1, Name: "Loja Atualizada"}

	mockRepo.EXPECT().
		Update(updatedStore).
		Return(nil)

	service := store.NewService(mockRepo)

	err := service.Update(updatedStore)
	assert.NoError(t, err)
}

func TestService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	mockRepo.EXPECT().
		Delete(1).
		Return(nil)

	service := store.NewService(mockRepo)

	err := service.Delete(1)
	assert.NoError(t, err)
}

func TestService_GetByEstablishmentID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockStoreRepository(ctrl)
	expectedStores := []store.Store{
		{ID: 1, EstablishmentID: 10},
		{ID: 2, EstablishmentID: 10},
	}

	mockRepo.EXPECT().
		FindByEstablishmentID(10).
		Return(expectedStores, nil)

	service := store.NewService(mockRepo)

	result, err := service.GetByEstablishmentID(10)
	assert.NoError(t, err)
	assert.Equal(t, expectedStores, result)
}