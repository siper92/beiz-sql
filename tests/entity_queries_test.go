package tests

import (
	beiz_sql "github.com/siper92/beiz-sql"
	"testing"
)

type TestEntity struct {
	ID   int    `db:"entity_id"`
	Name string `db:"name"`
	Test string `db:"test_field_1"`
	Path string `db:"test_path" attr:"path"`
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
	qb, ok := beiz_sql.NewSelectQuery().
		Entity(e).(beiz_sql.BeizEntityQueryBuilder)
	if !ok {
		t.Fatal("Not a BeizEntityQueryBuilder")
	}

	attribute := qb.GetAttributeMap()
}
