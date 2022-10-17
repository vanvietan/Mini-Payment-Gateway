package util

import (
	"gorm.io/gorm"
	"io/ioutil"
)

// ExecuteTestData execute test data sql
func ExecuteTestData(db *gorm.DB, filePath string) error {
	s, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	if tx := db.Exec(string(s)); tx.Error != nil {
		return tx.Error
	}
	return nil
}
