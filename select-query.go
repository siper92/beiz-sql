package beiz_sql

import (
	"fmt"
	"strings"
)

type LogicType string

const (
	AND LogicType = "AND"
	OR  LogicType = "OR"
)

type Table struct {
	raw   string
	alias string
	name  string
}

func (t *Table) String() string {
	if t.alias != "" {
		return fmt.Sprintf("%s %s", t.name, t.alias)
	}

	return t.name
}

type SqlCondition struct {
}

type JoinType string

const (
	Join  JoinType = "JOIN"
	Left  JoinType = "LEFT JOIN"
	Right JoinType = "RIGHT JOIN"
)

type JoinCondition struct {
	raw    string
	params []interface{}
	result string
	_type  JoinType
}

func (jc *JoinCondition) String() string {
	if jc.result == "" {
		return jc.result
	}

	return strings.ToLower(fmt.Sprintf(
		"%s %s",
		jc._type,
		fmt.Sprintf(jc.raw, jc.params...),
	))
}

// SelectQuery -> BeizQueryBuilder
type SelectQuery struct {
	entity     EntityInterface
	table      Table
	fields     []SelectField
	conditions []WhereCondition
	joints     []JoinCondition
	orderBy    string
}

func (qb *SelectQuery) Entity(e EntityInterface) BeizQueryBuilder {
	qb.entity = e
	qb.table = Table{
		raw:   e.TableName() + " e",
		name:  e.TableName(),
		alias: "e",
	}

	return qb
}

func tableStringToTable(raw string) Table {
	var alias string
	var name string

	if strings.Contains(raw, " as ") {
		splits := strings.Split(raw, " as ")
		name = splits[0]
		alias = splits[1]
	} else if strings.Contains(raw, " ") {
		splits := strings.Split(raw, " ")
		name = splits[0]
		alias = splits[1]
	}

	return Table{
		raw:   raw,
		name:  name,
		alias: alias,
	}
}

func (qb *SelectQuery) From(from string) BeizQueryBuilder {
	qb.table = tableStringToTable(from)
	return qb
}

type SelectField struct {
	raw   string
	alias string
	field string
	table Table
}

func (sf *SelectField) String() string {
	var selectV string
	if sf.table.alias != "" || sf.field != "" {
		selectV = sf.table.alias + "." + sf.field
		if sf.alias != "" {
			selectV += " as " + sf.alias
		}

		return selectV
	}

	return sf.raw
}

func stringToSelectFields(s string) []SelectField {
	var parts []string
	if strings.Contains(s, ",") {
		for _, part := range strings.Split(s, ",") {
			parts = append(parts, strings.TrimSpace(part))
		}
	} else {
		parts = append(parts, strings.TrimSpace(s))
	}

	var fields []SelectField
	for _, part := range parts {
		fieldStruct := SelectField{
			raw: strings.TrimSpace(part),
		}
		var alias string
		var field string
		var table Table

		if strings.Contains(fieldStruct.raw, " as ") {
			splits := strings.Split(fieldStruct.raw, " as ")
			field = splits[0]
			alias = splits[1]
		} else {
			field = fieldStruct.raw
		}

		if strings.Contains(field, ".") {
			splits := strings.Split(part, ".")
			table = Table{
				alias: splits[0],
			}
			field = splits[1]
		}

		fieldStruct.alias = alias
		fieldStruct.field = field
		fieldStruct.table = table

		fields = append(fields, fieldStruct)
	}

	return fields
}

func (qb *SelectQuery) Select(fields ...string) BeizQueryBuilder {
	for _, field := range fields {
		for _, f := range stringToSelectFields(field) {
			qb.fields = append(qb.fields, f)
		}
	}

	return qb
}

func (qb *SelectQuery) Join(joins string, params ...interface{}) BeizQueryBuilder {
	qb.joints = append(qb.joints, JoinCondition{
		_type:  Join,
		raw:    joins,
		params: params,
		result: fmt.Sprintf(joins, params...),
	})
	return qb
}

func (qb *SelectQuery) LeftJoin(joins string, params ...interface{}) BeizQueryBuilder {
	qb.joints = append(qb.joints, JoinCondition{
		_type:  Left,
		raw:    joins,
		params: params,
		result: fmt.Sprintf(joins, params...),
	})
	return qb
}

func (qb *SelectQuery) RightJoin(joins string, params ...interface{}) BeizQueryBuilder {
	qb.joints = append(qb.joints, JoinCondition{
		_type:  Right,
		raw:    joins,
		params: params,
		result: fmt.Sprintf(joins, params...),
	})
	return qb
}

func (qb *SelectQuery) OrderBy(order string) BeizQueryBuilder {
	if !strings.Contains(order, "order by") {
		order = fmt.Sprintf("order by %s", order)
	}

	qb.orderBy = order

	return qb
}

func NewSelectQuery() BeizQueryBuilder {
	return &SelectQuery{}
}

func NewEntitySelectQuery(e EntityInterface) BeizQueryBuilder {
	return &SelectQuery{
		table: Table{
			raw:   e.TableName() + " e",
			name:  e.TableName(),
			alias: "e",
		},
	}
}
