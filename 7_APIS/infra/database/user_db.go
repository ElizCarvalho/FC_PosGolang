package database

import (
	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) Create(user *entity.User) error {
	return u.db.Create(user).Error
}

func (u *UserDB) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := u.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
