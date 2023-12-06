package valueobject

type Query struct {
	Page       int
	Limit      int
	Sort       string
	Conditions []*Condition
	Joins      []string
	Locked     bool
}

func NewQuery() *Query {
	query := &Query{}
	query.Conditions = make([]*Condition, 0)
	query.Joins = make([]string, 0)
	return query
}

func (q *Query) Condition(field string, operator Operator, value any) *Query {
	condition := NewCondition(field, operator, value)
	q.Conditions = append(q.Conditions, condition)
	return q
}

func (q *Query) GetConditionValue(field string) any {
	for _, condition := range q.Conditions {
		if condition.Field == field {
			return condition.Value
		}
	}
	return nil
}

func (q *Query) IsConditionExist(field string) bool {
	for _, condition := range q.Conditions {
		if condition.Field == field {
			return true
		}
	}
	return false
}

func (q *Query) Paginate(page int, perPage int) *Query {
	q.Page = page
	q.Limit = perPage
	return q
}

func (q *Query) Order(orderedBy string) *Query {
	q.Sort = orderedBy
	return q
}

func (q *Query) With(entity string) *Query {
	q.Joins = append(q.Joins, entity)
	return q
}

func (q *Query) Lock() *Query {
	q.Locked = true
	return q
}
