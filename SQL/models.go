// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dbsqlc

import (
	"database/sql"
)

type User struct {
	ID     int64
	Name   string
	Status sql.NullBool
}
