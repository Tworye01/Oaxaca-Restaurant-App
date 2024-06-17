package models

import (
	"team-project/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

// this test does not work
func TestUsersAutoMigrate(t *testing.T) {
	s := database.New()
	err := s.AutoMigrate(&Credentials{})
	assert.NoError(t, err)
}

func TestRole(t *testing.T) {
	str := Customer.String()
	assert.Equal(t, "Customer", str)

	str = Waiter.String()
	assert.Equal(t, "Waiter", str)

	str = Kitchen.String()
	assert.Equal(t, "Kitchen", str)

	str = Manager.String()
	assert.Equal(t, "Manager", str)

	str = Admin.String()
	assert.Equal(t, "Admin", str)
}
