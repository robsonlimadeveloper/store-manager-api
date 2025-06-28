// app/modules/store/service_test.go
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
