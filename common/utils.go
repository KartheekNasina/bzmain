package common

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/lib/pq"
)

func GenerateUpdateQuery(tableName string, idColumn string, idValue interface{}, obj interface{}) (string, []interface{}, error) {
	var columns []string
	var values []interface{}
	var placeholders []string

	// Get the type of the struct
	typ := reflect.TypeOf(obj)

	// Iterate through the fields of the struct
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("db")

		// Skip fields without a "db" tag
		if tag == "" {
			continue
		}

		// Get the field value
		value := reflect.ValueOf(obj).Field(i).Interface()

		// Skip fields with zero or nil values
		if reflect.DeepEqual(value, reflect.Zero(field.Type).Interface()) {
			continue
		}

		// Special handling for "updated_at" column
		if tag == "updated_at" {
			columns = append(columns, fmt.Sprintf("%s = NOW()", tag))
		} else {
			columns = append(columns, fmt.Sprintf("%s = $%d", tag, len(values)+1))
			values = append(values, value)
			placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
		}
	}

	if len(columns) == 0 {
		return "", nil, fmt.Errorf("no fields to update")
	}

	// Construct the SQL query with placeholders
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = %s", tableName, strings.Join(columns, ", "), idColumn, fmt.Sprintf("$%d", len(values)+1))

	// Append the ID value to the slice of values
	values = append(values, idValue)

	return query, values, nil
}

func ScanRowToModel(row pgx.Row, model interface{}) error {

	val := reflect.ValueOf(model).Elem()
	args := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		args[i] = val.Field(i).Addr().Interface()
	}

	if err := row.Scan(args...); err != nil {
		return err
	}

	return nil
}

func ScanRowsToModel(rows pgx.Rows, model interface{}) error {

	val := reflect.ValueOf(model).Elem()
	args := make([]interface{}, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		args[i] = val.Field(i).Addr().Interface()
	}

	if err := rows.Scan(args...); err != nil {
		return err
	}

	return nil
}

func GetInsertSQLAndArgs(model interface{}, tableName string, ignoreFields ...string) (string, []interface{}) {
	val := reflect.ValueOf(model).Elem()
	typ := val.Type()

	ignoreMap := make(map[string]bool)
	for _, field := range ignoreFields {
		ignoreMap[field] = true
	}

	var columns, placeholders []string
	var args []interface{}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		if _, shouldIgnore := ignoreMap[field.Name]; shouldIgnore {
			continue
		}

		dbTag, ok := field.Tag.Lookup("db")
		if !ok {
			continue // skip if there's no 'db' tag
		}

		columns = append(columns, dbTag)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(columns))) // Adjusted to match the number of columns

		// Handle special types like slices which need to be converted to PostgreSQL arrays
		if val.Field(i).Type() == reflect.TypeOf([]string{}) {
			args = append(args, pq.Array(val.Field(i).Interface().([]string)))
		} else {
			args = append(args, val.Field(i).Interface())
		}
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	return query, args
}
