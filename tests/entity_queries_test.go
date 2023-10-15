package tests

import (
	beiz_sql "github.com/siper92/beiz-sql"
	"testing"
)

func TestEntityGetAttributeFields(t *testing.T) {
	qb := beiz_sql.NewSelectQuery()

	_ = qb.From("eav_attribute a")
}
