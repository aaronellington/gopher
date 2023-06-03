package utilities_test

import (
	"testing"

	"github.com/aaronellington/gopher/dbal/internal/utilities"
	"github.com/kyberbits/forge/forgetest"
)

func TestSave(t *testing.T) {
	type Entity struct {
		ID   uint64 `db:"id"`
		Name string `db:"user.name"`
		Test bool
	}

	expectedQuery := "INSERT INTO `tableName` (`id`) VALUES (:id) ON DUPLICATE KEY UPDATE `id` = :id;"
	actualQuery := utilities.GenerateSaveQuery("tableName", Entity{}, true)

	if err := forgetest.Assert(expectedQuery, actualQuery); err != nil {
		t.Fatal(err)
	}
}

func TestDelete(t *testing.T) {
	type Entity struct {
		Email  uint64 `db:"email" primaryKey:""`
		Phone  uint64 `db:"phoneNumber" primaryKey:""`
		Active bool   `db:"active"`
		Test   bool
	}

	expectedQuery := "DELETE FROM `tableName` WHERE `email` = :email AND `phoneNumber` = :phoneNumber"
	actualQuery := utilities.GenerateDeleteByPrimaryKeyQuery("tableName", Entity{})

	if err := forgetest.Assert(expectedQuery, actualQuery); err != nil {
		t.Fatal(err)
	}
}

func TestSelect(t *testing.T) {
	type Entity struct {
		Email  uint64 `db:"email" primaryKey:""`
		Phone  uint64 `db:"phoneNumber" primaryKey:""`
		Name   string `db:"user.name"`
		Active bool   `db:"active"`
		Test   bool
	}

	expectedQuery := "SELECT `email`, `phoneNumber`, `user`.`name` as `user.name`, `active` FROM `tableName` LEFT JOIN `user`"
	actualQuery := utilities.GenerateSelect("tableName", "LEFT JOIN `user`", Entity{})

	if err := forgetest.Assert(expectedQuery, actualQuery); err != nil {
		t.Fatal(err)
	}
}
