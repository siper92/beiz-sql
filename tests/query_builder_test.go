package tests

import (
	beiz_sql "github.com/siper92/beiz-sql"
	core_utils "github.com/siper92/core-utils"
	"testing"
)

func TestSimpleQuery(t *testing.T) {
	qb := beiz_sql.NewSelectQuery().
		From("test t").
		Select("t.id", "t.name").
		Where("t.id = ?", 1).
		OrderBy("t.id desc")

	qResult, err := qb.SQL()
	if err != nil {
		t.Fatal(err)
	}

	core_utils.AMatchesB(t,
		"select t.id, t.name from test t\nwhere t.id = ?\norder by t.id desc",
		qResult,
	)
}

func TestSimpleJoinQuery(t *testing.T) {
	qb := beiz_sql.NewSelectQuery().
		From("test t").
		Select("t.id", "t.name").
		Where("t.id = ?", 1).
		Join("test2 t2 on t2.id = t.id")

	qResult, err := qb.SQL()
	if err != nil {
		t.Fatal(err)
	}

	core_utils.AMatchesB(t,
		`select t.id, t.name from test t
join test2 t2 on t2.id = t.id
where t.id = ?`,
		qResult,
	)

	qb = beiz_sql.NewSelectQuery().
		From("test t").
		Select("t.id", "t.name").
		Where("t.id = ?", 1).
		RightJoin("test2 t2 on t2.id = t.id")

	core_utils.AMatchesB(t,
		`select t.id, t.name from test t
right join test2 t2 on t2.id = t.id
where t.id = ?`,
		qb.MustBuild(),
	)

	qb = beiz_sql.NewSelectQuery().
		From("test t").
		Select("t.id", "t.name").
		Where("t.id = ?", 1).
		LeftJoin("test2 t2 on t2.id = t.id")

	core_utils.AMatchesB(t,
		`select t.id, t.name from test t
left join test2 t2 on t2.id = t.id
where t.id = ?`,
		qb.MustBuild(),
	)
}

func TestMultipleSelects(t *testing.T) {
	qb := beiz_sql.NewSelectQuery().
		From("test t").
		Select("t.id").
		Where("t.id = ?", 1).
		Join("test2 t2 on t2.id = t.id").
		Select("t2.name, t2.id").
		Select("t.name")

	core_utils.AMatchesB(t,
		`select t.id, t2.name, t2.id, t.name from test t
join test2 t2 on t2.id = t.id
where t.id = ?`,
		qb.MustBuild(),
	)
}

func TestMultipleJoins(t *testing.T) {
	qb := beiz_sql.NewSelectQuery().
		From("test t").
		Select("*").
		Where("t.id = ?", 1).
		Join("test2 t2 on t2.id = t.id").
		Join("test3 t3 on t3.name = t.name").
		Where("t3.id > ?", 2)

	core_utils.AMatchesB(t,
		`select * from test t
join test2 t2 on t2.id = t.id
join test3 t3 on t3.name = t.name
where t.id = ? and t3.id > ?`,
		qb.MustBuild(),
	)
}

func TestJoinsFormatting(t *testing.T) {
	qb := beiz_sql.NewSelectQuery().
		From("test t").
		Join("test2 t2 on t2.id = t.id and code = '%s'", "TEST")

	core_utils.AMatchesB(t,
		`select * from test t
join test2 t2 on t2.id = t.id and code = 'TEST'`,
		qb.MustBuild(),
	)
}
