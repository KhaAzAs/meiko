package information

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/melodiez14/meiko/src/util/conn"
)

func Get(column ...string) QueryGet {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColTitle,
			ColDescription,
			ColType,
			ColCourseID,
			ColCreatedAt,
			ColUpdatedAt,
		}
	} else {
		for _, v := range column {
			c = append(c, v)
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QueryGet{fmt.Sprintf(queryGet, columnQuery)}
}

func (q QueryGet) Where(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryGet) AndWhere(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryGet) OrWhere(column, operator string, value interface{}) QueryGet {
	switch value.(type) {
	case int, int8, int64:
		return QueryGet{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryGet{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryGet) Exec() (Information, error) {
	var information Information
	query := fmt.Sprintf("%s LIMIT 1", q.string)
	err := conn.DB.Get(&information, query)
	if err != nil {
		return information, err
	}
	return information, nil
}

func Select(column ...string) QuerySelect {
	var c []string
	if len(column) < 1 {
		c = []string{
			ColID,
			ColTitle,
			ColDescription,
			ColType,
			ColCourseID,
			ColCreatedAt,
			ColUpdatedAt,
		}
	} else {
		for _, v := range column {
			c = append(c, v)
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QuerySelect{fmt.Sprintf(querySelect, columnQuery)}
}

func (q QuerySelect) Where(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		if len(rv) < 1 {
			return q
		}
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QuerySelect{fmt.Sprintf("%s WHERE %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

func (q QuerySelect) AndWhere(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		if len(rv) < 1 {
			return q
		}
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QuerySelect{fmt.Sprintf("%s AND %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

func (q QuerySelect) OrWhere(column, operator string, value interface{}) QuerySelect {
	switch value.(type) {
	case int, int8, int64:
		return QuerySelect{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QuerySelect{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	case []int64:
		var vals []string
		rv := reflect.ValueOf(value).Interface().([]int64)
		if len(rv) < 1 {
			return q
		}
		for _, v := range rv {
			vals = append(vals, fmt.Sprintf("%d", v))
		}
		str := strings.Join(vals, ", ")
		return QuerySelect{fmt.Sprintf("%s OR %s %s (%s)", q.string, column, operator, str)}
	default:
		return q
	}
}

func (q QuerySelect) OrderBy(column, order string) QuerySelect {

	if order != OrderAsc && order != OrderDesc {
		return q
	}

	return QuerySelect{fmt.Sprintf("%s ORDER BY %s %s", q.string, column, order)}
}

func (q QuerySelect) Limit(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s LIMIT %d", q.string, value)}
}

func (q QuerySelect) Offset(value uint16) QuerySelect {
	return QuerySelect{fmt.Sprintf("%s OFFSET %d", q.string, value)}
}

func (q QuerySelect) Exec() ([]Information, error) {
	var informations []Information
	err := conn.DB.Select(&informations, q.string)
	if err != nil {
		return informations, err
	}
	return informations, nil
}

func Insert(column map[string]interface{}) QueryInsert {

	c := []string{"created_at", "updated_at"}
	v := []string{"NOW()", "NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int8, int64:
			c = append(c, i)
			v = append(v, fmt.Sprintf("(%d)", val))
		case string:
			c = append(c, i)
			v = append(v, fmt.Sprintf("('%s')", val))
		case sql.NullString:
			str := reflect.ValueOf(val).Interface().(sql.NullString)
			c = append(c, i)
			if str.Valid {
				v = append(v, fmt.Sprintf("('%s')", str.String))
			} else {
				v = append(v, fmt.Sprintf("(NULL)"))
			}
		}
	}
	columnQuery := strings.Join(c, ", ")
	valueQuery := strings.Join(v, ", ")
	return QueryInsert{fmt.Sprintf(queryInsert, columnQuery, valueQuery)}
}

func (q QueryInsert) Exec() error {
	result, err := conn.DB.Exec(q.string)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}

func Update(column map[string]interface{}) QueryUpdate {
	c := []string{"updated_at = NOW()"}
	for i, val := range column {
		switch val.(type) {
		case int, int8, int64:
			c = append(c, fmt.Sprintf("%s = (%d)", i, val))
		case string:
			c = append(c, fmt.Sprintf("%s = ('%s')", i, val))
		case sql.NullString:
			str := reflect.ValueOf(val).Interface().(sql.NullString)
			if str.Valid {
				c = append(c, fmt.Sprintf("%s = ('%s')", i, str.String))
			} else {
				c = append(c, fmt.Sprintf("%s = NULL", i))
			}
		}
	}
	columnQuery := strings.Join(c, ", ")
	return QueryUpdate{fmt.Sprintf(queryUpdate, columnQuery)}
}

func (q QueryUpdate) Where(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s WHERE %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s WHERE %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryUpdate) AndWhere(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s AND %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s AND %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryUpdate) OrWhere(column, operator string, value interface{}) QueryUpdate {
	switch value.(type) {
	case int, int8, int64:
		return QueryUpdate{fmt.Sprintf("%s OR %s %s (%d)", q.string, column, operator, value)}
	case string:
		return QueryUpdate{fmt.Sprintf("%s OR %s %s ('%s')", q.string, column, operator, value)}
	default:
		return q
	}
}

func (q QueryUpdate) Exec() error {
	result, err := conn.DB.Exec(q.string)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("No rows affected")
	}
	return nil
}
