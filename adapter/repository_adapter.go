package adapter

import "gorm.io/gorm"

type RepositoryAdapter struct {
	db *gorm.DB
}
