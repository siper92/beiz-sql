package beiz_sql

type EntityDefinition struct {
	Table   string
	IDField string
}

type (
	EntityInterface interface {
		TableName() string
		Definition() EntityDefinition
	}
)
