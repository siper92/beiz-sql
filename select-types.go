package beiz_sql

type HasRawSql interface {
	Raw() string
}

type BeizQueryBuilder interface {
	Entity(e EntityInterface) BeizQueryBuilder
	From(from string) BeizQueryBuilder
	Select(fields ...string) BeizQueryBuilder
	Where(condition string, params ...interface{}) BeizQueryBuilder
	Join(joins string, params ...interface{}) BeizQueryBuilder
	LeftJoin(joins string, params ...interface{}) BeizQueryBuilder
	RightJoin(joins string, params ...interface{}) BeizQueryBuilder
	OrderBy(order string) BeizQueryBuilder

	GetParams() []interface{}
	SQL() (string, error)
	MustBuild() string
	DebugSql() string
}

type BeizQueryRunner interface {
	Execute(bind interface{}) error
}
