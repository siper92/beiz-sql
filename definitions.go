package beiz_sql

type HasRawSql interface {
	Raw() string
}

type EntityDefinition struct {
	Table   string
	IDField string
}

type EntityAttribute struct {
	DbFieldName string
	DbTable     string
}

type EntityInterface interface {
	TableName() string
	Definition() EntityDefinition
}

type BeizQueryRunner interface {
	Execute(bind interface{}) error
}

type BeizEntityQueryBuilder interface {
	Entity(e EntityInterface) BeizQueryBuilder
	GetAttributeMap() map[string]EntityAttribute
	GetAttributeField(propKey string) EntityAttribute
	JoinAttributes(propKey string) BeizQueryBuilder
}

type BeizQueryBuilder interface {
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
