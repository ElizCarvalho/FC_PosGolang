package database

import (
	"testing"

	"github.com/ElizCarvalho/FC_PosGolang/7_APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	db.AutoMigrate(&entity.User{})

	user, _ := entity.NewUser("John Doe", "john.doe@example.com", "123456")
	userDB := NewUserDB(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound *entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)

	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

