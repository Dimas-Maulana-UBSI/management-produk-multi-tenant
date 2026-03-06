package helper

import (
	"errors"
	"regexp"
)

var validDBName = regexp.MustCompile(`^db_[a-zA-Z0-9_]+$`)

func ValidateDBName(dbName string) error {
	if !validDBName.MatchString(dbName) {
		return errors.New("invalid database name")
	}
	return nil
}
