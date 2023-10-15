package beiz_sql

import (
	"fmt"
	core_utils "github.com/siper92/core-utils"
	"strings"
)

func (qb *SelectQuery) buildSelect() string {
	selectString := fmt.Sprintf("select selected_data from %s",
		qb.table.String(),
	)

	fieldsString := "*"
	if len(qb.fields) > 0 {
		fieldsString = ""
		for _, field := range qb.fields {
			if fieldsString != "" {
				fieldsString += ", "
			}

			fieldsString += field.String()
		}
	}

	return strings.Replace(selectString, "selected_data", fieldsString, 1)
}

func (qb *SelectQuery) buildWhere() string {
	selectString := "where "
	if len(qb.conditions) > 0 {
		for i, condition := range qb.conditions {
			if i > 0 {
				// AND/OR
				selectString += string(condition.logicType) + " "
			}

			selectString += condition.String() + " "
		}
	}

	return selectString
}

func (qb *SelectQuery) buildJoins() string {
	var joinString string
	for _, join := range qb.joints {
		joinString += fmt.Sprintf("\n%s", join.String())
	}

	return joinString
}

func (qb *SelectQuery) SQL() (string, error) {
	selectString := qb.buildSelect()

	if len(qb.joints) > 0 {
		selectString += qb.buildJoins()
	}

	if len(qb.conditions) > 0 {
		selectString += "\n" + strings.TrimSpace(
			qb.buildWhere(),
		)
	}

	selectString += "\n" + qb.orderBy
	selectString = strings.Trim(selectString, "\n")
	selectString = strings.TrimSpace(selectString)

	return selectString, nil
}

func (qb *SelectQuery) MustBuild() string {
	buildResult, err := qb.SQL()
	core_utils.ErrorWarning(err)
	return buildResult
}

func (qb *SelectQuery) DebugSql() string {
	sql, err := qb.SQL()
	if err != nil {
		core_utils.ErrorWarning(err)
		return ""
	}

	params := qb.GetParams()
	count := strings.Count(sql, "?")

	for i := 0; i < count; i++ {
		param := ""
		switch params[i].(type) {
		case string:
			param = fmt.Sprintf("'%s'", params[i])
		default:
			param = fmt.Sprintf("%v", params[i])
		}

		sql = strings.Replace(sql, "?", param, 1)
	}

	return sql
}
