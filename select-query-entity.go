package beiz_sql

// SelectQuery -> BeizEntityQueryBuilder

func (qb *SelectQuery) Entity(e EntityInterface) BeizQueryBuilder {
	qb.entity = e
	qb.table = Table{
		raw:   e.TableName() + " e",
		name:  e.TableName(),
		alias: "e",
	}

	return qb
}

func (qb *SelectQuery) GetAttributeMap() map[string]EntityAttribute {
	//TODO implement me
	panic("implement me")
}

func (qb *SelectQuery) GetAttributeField(propKey string) EntityAttribute {
	//TODO implement me
	panic("implement me")
}

func (qb *SelectQuery) JoinAttributes(propKey string) BeizQueryBuilder {
	//TODO implement me
	panic("implement me")
}
