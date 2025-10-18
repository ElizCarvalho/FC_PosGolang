//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/events"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/entity"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/event"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/database"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/infra/web"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/usecase"
	"github.com/google/wire"
	"github.com/streadway/amqp"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

func NewHealthHandler(db *sql.DB, rabbitMQChannel *amqp.Channel) *web.HealthHandler {
	wire.Build(
		web.NewHealthHandler,
	)
	return &web.HealthHandler{}
}
