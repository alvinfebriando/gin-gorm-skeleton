package valueobject

import "fmt"

type Operator string

const (
	Equal            Operator = "="
	NotEqual                  = "!="
	LessThan                  = "<"
	GreaterThan               = ">"
	LessThanEqual             = "<="
	GreaterThanEqual          = ">="
	Is                        = "IS"
	In                        = "IN"
	Not                       = "NOT"
	Like                      = "LIKE"
	NotLike                   = "NOT LIKE"
	ILike                     = "ILIKE"
	NotILike                  = "NOT ILIKE"
)

type Condition struct {
	Field     string
	Operation Operator
	Value     any
}

func NewCondition(field string, operation Operator, value any) *Condition {
	if value == "" {
		return nil
	}
	if operation == ILike || operation == Like || operation == NotILike || operation == NotLike {
		value = "%" + fmt.Sprintf("%v", value) + "%"
	}
	return &Condition{
		Field:     field,
		Operation: operation,
		Value:     value,
	}
}
