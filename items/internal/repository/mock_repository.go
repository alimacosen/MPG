package repository

import (
	"context"
	"github.com/stretchr/testify/mock"
	"math/rand"
	"mpg/items/internal/model"
	"time"
)

type MockItemRepository struct {
	mock.Mock
}

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

func NewMockItemRepository() *MockItemRepository {
	return &MockItemRepository{}
}

func generateID(length int) string {
	id := make([]byte, length)
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range id {
		id[i] = charset[r.Intn(len(charset))]
	}
	return string(id)
}

func (m *MockItemRepository) Create(ctx context.Context, item *model.Item) (*model.Item, error) {
	args := m.Called(item)
	mockItem := args.Get(0).(*model.Item)
	mockItem.ID = generateID(24)
	return mockItem, nil
}

func (m *MockItemRepository) FindByID(ctx context.Context, idList []string) ([]*model.Item, error) {
	args := m.Called(idList)

	ids := []string{"645878736c300c61a5780076", "6458789d6c300c61a5780077", "6458793e59655d088fd08f4a"}
	names := []string{"sword", "shield", "horse"}
	description := []string{"good", "very good", "awesome"}
	damage := []int{10, 20, 30}
	healing := []int{100, 90, 80}
	protection := []int{45, 34, 26}

	items := make([]*model.Item, len(ids))
	for i, _ := range ids {
		items[i] = &model.Item{
			ID:          ids[i],
			Name:        &names[i],
			Description: &description[i],
			Damage:      &damage[i],
			Healing:     &healing[i],
			Protection:  &protection[i],
		}
	}

	res := make([]*model.Item, 0)
	for _, idPassedIn := range args.Get(0).([]string) {
		for j, id := range ids {
			if id == idPassedIn {
				res = append(res, items[j])
				break
			}
		}
	}

	return res, nil
}

func (m *MockItemRepository) FindAll(ctx context.Context) ([]*model.Item, error) {
	return nil, nil
}

func (m *MockItemRepository) Update(ctx context.Context, id string, updateFields *model.Item) (int, error) {
	return 1, nil
}

func (m *MockItemRepository) Delete(ctx context.Context, id string) (int, error) {
	return 1, nil
}
