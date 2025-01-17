package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user, err := CreateUser("John Doe", "john@doe.com", "Az#123456")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@doe.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := CreateUser("John Doe", "john@doe.com", "Az#123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("Az#123456"))
	assert.False(t, user.ValidatePassword("Az#123457"))
	assert.NotEmpty(t, "Az#123456", user.Password)

}
