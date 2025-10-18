package graph

import (
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/entity"
	"github.com/ElizCarvalho/FC_PosGolang/20_CleanArch/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	OrderRepository    entity.OrderRepositoryInterface
}
