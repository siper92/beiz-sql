package beiz_sql

import (
	"strings"
)

type WhereCondition struct {
	raw       string
	logicType LogicType
	field     SelectField
	params    []interface{}
}

func (wc *WhereCondition) String() string {
	return wc.raw
}

func (qb *SelectQuery) GetParams() []interface{} {
	var params []interface{}

	for _, condition := range qb.conditions {
		for _, part := range condition.params {
			params = append(params, part)
		}
	}

	return params
}

func getFieldsFromCondition(condition string) []string {
	var fields []string
	for _, part := range strings.Split(condition, " ") {
		if strings.Contains(part, ".") {
			fields = append(fields, part)
		}
	}

	return fields
}

func (qb *SelectQuery) Where(condition string, params ...interface{}) BeizQueryBuilder {

	rawCondition := strings.TrimSpace(strings.ToLower(condition))
	logicType := AND

	if strings.Index(rawCondition, "or ") == 0 {
		logicType = OR
		rawCondition = strings.Replace(rawCondition, "or ", "", 1)
	}

	qb.conditions = append(qb.conditions, WhereCondition{
		raw:       rawCondition,
		logicType: logicType,
		params:    params,
	})
	return qb
}
