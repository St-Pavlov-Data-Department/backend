package utils

import "gorm.io/gorm"

// -------- Database

func WithTransaction(db *gorm.DB, f func(*gorm.DB) error) error {
	transaction := db.Begin()
	err := f(db)
	if err != nil {
		transaction.Rollback()
	} else {
		transaction.Commit()
	}
	return err
}
