package web

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestWebOrderHandler_Create(t *testing.T) {
	orderRepository := new(MockOrderRepository)
	eventDispatcher := new(MockEventDispatcher)
	event := new(MockEvent)

	orderRepository.On("Save", mock.AnythingOfType("*entity.Order")).Return(nil)
	event.On("SetPayload", mock.AnythingOfType("usecase.OrderOutputDTO"))
	eventDispatcher.On("Dispatch", event).Return(nil)

	handler := NewWebOrderHandler(eventDispatcher, orderRepository, event)

	requestBody := map[string]interface{}{
		"id":    "123",
		"price": 10.0,
		"tax":   2.0,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "123", response["id"])
	assert.Equal(t, 10.0, response["price"])
	assert.Equal(t, 2.0, response["tax"])
	assert.Equal(t, 12.0, response["final_price"])

	orderRepository.AssertExpectations(t)
	event.AssertExpectations(t)
	eventDispatcher.AssertExpectations(t)
}

func TestWebOrderHandler_Create_WithInvalidJSON(t *testing.T) {
	orderRepository := new(MockOrderRepository)
	eventDispatcher := new(MockEventDispatcher)
	event := new(MockEvent)

	handler := NewWebOrderHandler(eventDispatcher, orderRepository, event)

	req := httptest.NewRequest("POST", "/order", bytes.NewBufferString("invalid json"))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestWebOrderHandler_Create_WhenRepositoryFails(t *testing.T) {
	orderRepository := new(MockOrderRepository)
	eventDispatcher := new(MockEventDispatcher)
	event := new(MockEvent)

	orderRepository.On("Save", mock.AnythingOfType("*entity.Order")).Return(assert.AnError)

	handler := NewWebOrderHandler(eventDispatcher, orderRepository, event)

	requestBody := map[string]interface{}{
		"id":    "123",
		"price": 10.0,
		"tax":   2.0,
	}

	jsonBody, _ := json.Marshal(requestBody)
	req := httptest.NewRequest("POST", "/order", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	orderRepository.AssertExpectations(t)
	event.AssertNotCalled(t, "SetPayload")
	eventDispatcher.AssertNotCalled(t, "Dispatch")
}
