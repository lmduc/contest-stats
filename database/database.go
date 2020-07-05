package database

import "github.com/jinzhu/gorm"

// Database is an interface that could connect or close the connection
// to a specific database
type Database interface {
	Connect() error
	Close() error
	Connection() *gorm.DB
}
