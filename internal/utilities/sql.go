package utilities

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateSaveQuery(
	tableName string,
	entity any,
	updateOnDuplicate bool,
) string {
	columns := []string{}
	valuePlaceholders := []string{}
	updates := []string{}

	element := reflect.ValueOf(entity)
	for i := 0; i < element.NumField(); i++ {
		fieldDefinition := element.Type().Field(i)

		columnName := fieldDefinition.Tag.Get("db")
		if columnName == "" {
			continue
		}

		if strings.Contains(columnName, ".") {
			continue
		}

		columns = append(columns, fmt.Sprintf("`%s`", columnName))
		valuePlaceholders = append(valuePlaceholders, fmt.Sprintf(":%s", columnName))
		updates = append(updates, fmt.Sprintf("`%s` = :%s", columnName, columnName))
	}

	queryText := fmt.Sprintf(
		"INSERT INTO `%s` (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(valuePlaceholders, ", "),
	)

	if updateOnDuplicate {
		queryText += fmt.Sprintf(
			" ON DUPLICATE KEY UPDATE %s;",
			strings.Join(updates, ", "),
		)
	}

	return queryText
}

func GenerateDeleteByPrimaryKeyQuery(tableName string, entity any) string {
	deletes := []string{}

	element := reflect.ValueOf(entity)
	for i := 0; i < element.NumField(); i++ {
		fieldDefinition := element.Type().Field(i)

		_, ok := fieldDefinition.Tag.Lookup("primaryKey")
		if !ok {
			continue
		}

		columnName := fieldDefinition.Tag.Get("db")

		deletes = append(deletes, fmt.Sprintf("`%s` = :%s", columnName, columnName))
	}

	queryText := fmt.Sprintf(
		"DELETE FROM `%s` WHERE %s",
		tableName,
		strings.Join(deletes, " AND "),
	)

	return queryText
}

func GenerateSelect(tableName string, joins string, entity any) string {
	selects := []string{}

	element := reflect.ValueOf(entity)
	for i := 0; i < element.NumField(); i++ {
		fieldDefinition := element.Type().Field(i)

		columnName := fieldDefinition.Tag.Get("db")
		if columnName == "" {
			continue
		}

		parts := strings.Split(columnName, ".")

		if len(parts) == 2 {
			selects = append(selects, fmt.Sprintf("`%s`.`%s` as `%s`", parts[0], parts[1], columnName))

			continue
		}

		selects = append(selects, fmt.Sprintf("`%s`", columnName))
	}

	return strings.TrimSpace(fmt.Sprintf("SELECT %s FROM `%s` %s", strings.Join(selects, ", "), tableName, joins))
}
