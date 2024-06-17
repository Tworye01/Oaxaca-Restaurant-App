package database

import (
	"team-project/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test Mock struct
type Mock struct {
	models.Model
	Name        string
	Description string
}

func TestNew(t *testing.T) {
	s := New()
	assert.NotNil(t, s)
}

func TestPing(t *testing.T) {
	s := New()

	err := s.Ping()
	assert.NoError(t, err)
}

func TestClose(t *testing.T) {
	s := New()

	err := s.Close()
	assert.NoError(t, err)
}

func TestList(t *testing.T) {
	s := New()

	var m []Mock
	err := s.List(&m, "mocks")
	assert.NoError(t, err)
	assert.True(t, len(m) > 0)
}

func TestGet(t *testing.T) {
	s := New()

	m := Mock{}
	err := s.Get(&m, "mocks", 1)
	assert.NoError(t, err)
	assert.NotNil(t, m)

	type Mock2 struct{}

	m2 := Mock2{}
	err = s.Get(&m2, "mocks", 500)
	assert.Error(t, err)
}

func TestAdd(t *testing.T) {
	s := New()
	m := Mock{Name: "Potato", Description: "It is not poisonous"}

	err := s.Add(&m, "mocks")
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	s := New()
	mock := Mock{Model: models.Model{ID: 1}, Name: "Apple", Description: "It is poisonous"}

	err := s.Update(&mock, "mocks")
	assert.NoError(t, err)
}

func TestDelete(t *testing.T) {
	s := New()
	mock := Mock{Name: "Apple", Description: "It is poisonous"}

	err := s.Add(&mock, "mocks")
	assert.NoError(t, err)

	err = s.Delete(&mock, "mocks")
	assert.NoError(t, err)
}

func TestContains(t *testing.T) {
	s := New()
	mock := Mock{Model: models.Model{ID: 1}}

	e := s.Contains(&mock, "mocks")
	assert.True(t, e)

	mock = Mock{Model: models.Model{ID: 500}}

	e = s.Contains(&mock, "mocks")
	assert.False(t, e)
}

func TestAutoMigrate(t *testing.T) {
	s := New()

	err := s.AutoMigrate(&Mock{})
	assert.NoError(t, err)
}
