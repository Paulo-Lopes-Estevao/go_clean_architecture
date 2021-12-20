package injection

import "gorm.io/gorm"

type RepositoryAdapter struct {
	db *gorm.DB
}
