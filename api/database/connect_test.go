package database

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	v := m.Run()
	os.Exit(v)
}

func TestPgConn(t *testing.T) {
	db, err := PgConn()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}

func TestSQLiteConn(t *testing.T) {
	db, err := SQLiteConn()
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
