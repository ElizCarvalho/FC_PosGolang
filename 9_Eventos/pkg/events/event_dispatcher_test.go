package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (e *TestEventHandler) Handle(event EventInterface) {
}

type EventDispatcherTestSuit struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuit) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "TestEvent", Payload: "TestPayload"}
	suite.event2 = TestEvent{Name: "TestEvent2", Payload: "TestPayload2"}

}

func (suite *EventDispatcherTestSuit) TestEventDispatcherRegister() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuit) TestEventDispatcherRegisterError() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Equal(suite.T(), ErrHandlerAlreadyRegistered, err)
	assert.Error(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcherClear() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event2.GetName()], 1)

	suite.eventDispatcher.Clear()
	assert.Len(suite.T(), suite.eventDispatcher.handlers, 0)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcherHas() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event2.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcherDispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcherRemove() {
	// Event 1
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 2)

	// Event 2
	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event2.GetName()], 1)

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event.GetName()], 1)
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event2.GetName()], 0)

	err = suite.eventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	assert.Nil(suite.T(), err)
	assert.Len(suite.T(), suite.eventDispatcher.handlers[suite.event2.GetName()], 0)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuit))
}
