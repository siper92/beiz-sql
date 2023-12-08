package tests

import (
	beiz_sql "github.com/siper92/beiz-sql"
	"testing"
)

type TestEntity struct {
	ID    int    `db:"entity_id"`
	Name  string `db:"name" attr:"path"`
	Test  string `db:"test_field_1"`
	Test2 string
	Path  string `db:"path"`
}

func (t TestEntity) TableName() string {
	return "test"
}

func (t TestEntity) Definition() beiz_sql.EntityDefinition {
	return beiz_sql.EntityDefinition{
		Table:   t.TableName(),
		IDField: "entity_id",
	}
}

func TestEntityGetAttributeFields(t *testing.T) {
	e := TestEntity{}
	testFieldAst := beiz_sql.GetStructFieldOrNil(e, "Test")
	_ = testFieldAst
}
