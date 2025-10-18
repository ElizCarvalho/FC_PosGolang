package usecase

import (
	"testing"
	"time"

	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/events"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetTotal() (int, error) {
	args := m.Called()
	return args.Int(0), args.Error(1)
}

func (m *MockOrderRepository) FindAll() ([]*entity.Order, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Order), args.Error(1)
}

type MockEventDispatcher struct {
	mock.Mock
}

func (m *MockEventDispatcher) Register(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *MockEventDispatcher) Dispatch(event events.EventInterface) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *MockEventDispatcher) Remove(eventName string, handler events.EventHandlerInterface) error {
	args := m.Called(eventName, handler)
	return args.Error(0)
}

func (m *MockEventDispatcher) Has(eventName string, handler events.EventHandlerInterface) bool {
	args := m.Called(eventName, handler)
	return args.Bool(0)
}

func (m *MockEventDispatcher) Clear() {
	m.Called()
}

type MockEvent struct {
	mock.Mock
}

func (m *MockEvent) GetName() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockEvent) GetPayload() interface{} {
	args := m.Called()
	return args.Get(0)
}

func (m *MockEvent) GetDateTime() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func (m *MockEvent) SetPayload(payload interface{}) {
	m.Called(payload)
}

func TestCreateOrderUseCase_Execute(t *testing.T) {
	orderRepository := new(MockOrderRepository)
	eventDispatcher := new(MockEventDispatcher)
	event := new(MockEvent)

	orderRepository.On("Save", mock.AnythingOfType("*entity.Order")).Return(nil)
	event.On("SetPayload", mock.AnythingOfType("usecase.OrderOutputDTO"))
	eventDispatcher.On("Dispatch", event).Return(nil)

	createOrderUseCase := NewCreateOrderUseCase(orderRepository, event, eventDispatcher)

	input := OrderInputDTO{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}

	output, err := createOrderUseCase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, "123", output.ID)
	assert.Equal(t, 10.0, output.Price)
	assert.Equal(t, 2.0, output.Tax)
	assert.Equal(t, 12.0, output.FinalPrice)

	orderRepository.AssertExpectations(t)
	event.AssertExpectations(t)
	eventDispatcher.AssertExpectations(t)
}

func TestCreateOrderUseCase_Execute_WhenRepositoryFails(t *testing.T) {
	orderRepository := new(MockOrderRepository)
	eventDispatcher := new(MockEventDispatcher)
	event := new(MockEvent)

	orderRepository.On("Save", mock.AnythingOfType("*entity.Order")).Return(assert.AnError)

	createOrderUseCase := NewCreateOrderUseCase(orderRepository, event, eventDispatcher)

	input := OrderInputDTO{
		ID:    "123",
		Price: 10.0,
		Tax:   2.0,
	}

	output, err := createOrderUseCase.Execute(input)

	assert.Error(t, err)
	assert.Equal(t, OrderOutputDTO{}, output)

	orderRepository.AssertExpectations(t)
	event.AssertNotCalled(t, "SetPayload")
	eventDispatcher.AssertNotCalled(t, "Dispatch")
}
