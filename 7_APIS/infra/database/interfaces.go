package database

import "github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
