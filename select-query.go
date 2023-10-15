package beiz_sql

import (
	"fmt"
	"strings"
)

type LogicType string

const (
	AND LogicType = "and"
	OR  LogicType = "or"
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
	Join  JoinType = "join"
	Left  JoinType = "left join"
	Right JoinType = "right join"
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

	return fmt.Sprintf(
		"%s %s",
		jc._type,
		fmt.Sprintf(jc.raw, jc.params...),
	)
}

// SelectQuery -> BeizQueryBuilder
// SelectQuery -> BeizEntityQueryBuilder
type SelectQuery struct {
	entity     EntityInterface
	table      Table
	fields     []SelectField
	conditions []WhereCondition
	joints     []JoinCondition
	orderBy    string
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
	selectV := sf.raw
	if sf.table.alias != "" && sf.field != "" {
		if sf.table.alias != "" {
			selectV = sf.table.alias + "." + sf.field
		} else {
			selectV = sf.field
		}

		if sf.alias != "" {
			selectV += " as " + sf.alias
		}
	}

	return selectV
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

func (qb *SelectQuery) join(joiType JoinType, join string, params ...interface{}) BeizQueryBuilder {
	qb.joints = append(qb.joints, JoinCondition{
		_type:  joiType,
		raw:    join,
		params: params,
		result: fmt.Sprintf(join, params...),
	})
	return qb
}

func (qb *SelectQuery) Join(joins string, params ...interface{}) BeizQueryBuilder {
	return qb.join(Join, joins, params...)
}

func (qb *SelectQuery) LeftJoin(joins string, params ...interface{}) BeizQueryBuilder {
	return qb.join(Left, joins, params...)
}

func (qb *SelectQuery) RightJoin(joins string, params ...interface{}) BeizQueryBuilder {
	return qb.join(Right, joins, params...)
}

func (qb *SelectQuery) OrderBy(order string) BeizQueryBuilder {
	if !strings.Contains(order, "order by") {
		order = fmt.Sprintf("order by %s", order)
	}

	qb.orderBy = order

	return qb
}

func NewSelectQuery() *SelectQuery {
	return &SelectQuery{}
}

func NewEntitySelectQuery(e EntityInterface) *SelectQuery {
	return &SelectQuery{
		table: Table{
			raw:   e.TableName() + " e",
			name:  e.TableName(),
			alias: "e",
		},
	}
}
