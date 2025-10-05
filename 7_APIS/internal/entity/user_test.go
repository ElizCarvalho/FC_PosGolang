package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testUserName  = "John Doe"
	testUserEmail = "john.doe@example.com"
	testPassword  = "123456"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser(testUserName, testUserEmail, "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, testUserName, user.Name)
	assert.Equal(t, testUserEmail, user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser(testUserName, testUserEmail, testPassword)
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword(testPassword))
	assert.False(t, user.ValidatePassword("123456789"))
	assert.NotEqual(t, testPassword, user.Password)
}
