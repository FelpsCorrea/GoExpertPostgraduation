package events

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (te TestEvent) GetName() string {
	return te.Name
}

func (te TestEvent) GetDateTime() time.Time {
	return time.Now()
}

func (te TestEvent) GetPayLoad() interface{} {
	return te.Payload
}

type TestEventHandler struct{}

func (teh TestEventHandler) Handle(event EventInterface) {}

// Ajuda a preparar os testes para não repetir código
type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

// Antes do teste rodar o setup roda
func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{}
	suite.handler2 = TestEventHandler{}
	suite.handler3 = TestEventHandler{}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}

// Rodar a suite
func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
