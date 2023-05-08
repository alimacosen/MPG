package service

import (
	"context"
	"github.com/stretchr/testify/assert"
	"mpg/items/internal/model"
	"mpg/items/internal/repository"
	"testing"
)

func TestService_Create(t *testing.T) {
	mockRepo := repository.NewMockItemRepository()
	itemService := NewItemService(mockRepo)

	name := "test item"
	description := "It's a test item"
	damage := 10
	healing := 20
	protection := 30

	item := model.Item{
		Name:        &name,
		Description: &description,
		Damage:      &damage,
		Healing:     &healing,
		Protection:  &protection,
	}

	// Set the expected behavior for the Create method on the mock repository
	mockRepo.On("Create", &item).Return(&item, nil)
	// Call the Create method on the service with the test item
	res, err := itemService.Create(context.Background(), &item)
	// Assert that there was no error
	assert.NoError(t, err)
	// Assert that the item was created successfully
	assert.NotNil(t, res)
	assert.Equal(t, len(item.ID), 24)
	assert.Equal(t, *item.Name, *res.Name)
	assert.Equal(t, *item.Description, *res.Description)
	assert.Equal(t, *item.Healing, *res.Healing)
	assert.Equal(t, *item.Protection, *res.Protection)
	assert.Equal(t, *item.Damage, *res.Damage)
	// Assert that the Create method was called on the mock repository with the correct argument
	mockRepo.AssertCalled(t, "Create", &item)
}
