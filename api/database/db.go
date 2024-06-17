package database

import (
	"context"
	"os"

	"gorm.io/gorm"
)

// Layer upon the gorm.DB
type Store struct {
	DB *gorm.DB
}

// Create a new Store, which has a connection to a gorm.DB.
// If env:DB is pg then it will use a sqlite connection, otherwise it will use a postgres connection
func New() *Store {
	t := os.Getenv("DB")

	if t == "sqlite" {
		db, err := SQLiteConn()
		if err != nil {
			panic(err)
		}

		return &Store{
			DB: db,
		}
	}

	db, err := PgConn()
	if err != nil {
		panic(err)
	}

	return &Store{
		DB: db,
	}
}

// Pings the database, returns any errors when trying to ping
func (s *Store) Ping() error {
	ctx := context.Background()

	db, err := s.DB.DB()
	if err != nil {
		return err
	}

	return db.PingContext(ctx)
}

// Close the database connection
func (s *Store) Close() error {
	db, err := s.DB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

// Gets all the records from table
func (s *Store) List(i interface{}, t string) error {
	result := s.DB.Table(t).Find(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Gets all the specific records from table
func (s *Store) ListWhere(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Find(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Gets a record using the primary key
func (s *Store) Get(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Find(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Adds a new model to the database
func (s *Store) Add(i interface{}, t string) error {
	result := s.DB.Table(t).Create(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Updates record in database
func (s *Store) Update(i interface{}, t string) error {
	result := s.DB.Table(t).Save(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Delete a record in the database
func (s *Store) Delete(i interface{}, t string, primary ...interface{}) error {
	result := s.DB.Table(t).Delete(i, primary...)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Check if a record is within the database
func (s *Store) Contains(i interface{}, t string, primary ...interface{}) bool {
	result := s.DB.Table(t).Model(i).First(i, primary...)
	return result.Error == nil
}

// Automatically Migrate all of the given Models
func (s *Store) AutoMigrate(i ...interface{}) error {
	if len(i) > 0 {
		for _, m := range i {
			err := s.DB.Migrator().AutoMigrate(&m)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
