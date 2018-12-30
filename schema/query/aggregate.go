package query

import (
	"strings"

	"github.com/oktacode/rest-layer/schema"
)

const (
	opGroup = "$group"
)

// Aggregate defines an expression against a schema to perform a match on schema's data.
type Aggregate []Expression

// Match implements Expression interface.
func (e Aggregate) Match(payload map[string]interface{}) bool {
	if e == nil || len(e) == 0 {
		// nil or empty predicates always match
		return true
	}
	// Run each sub queries like a root query, stop/pass on first match
	for _, subQuery := range e {
		if !subQuery.Match(payload) {
			return false
		}
	}
	return true
}

// String implements Expression interface.
func (e Aggregate) String() string {
	if len(e) == 0 {
		return "{}"
	}
	s := make([]string, 0, len(e))
	for _, subQuery := range e {
		s = append(s, subQuery.String())
	}
	return "{" + strings.Join(s, ", ") + "}"
}

// Prepare implements Expression interface.
func (e Aggregate) Prepare(validator schema.Validator) error {
	return prepareExpressions(e, validator)
}

// Group matches values that match to a specified regular expression.
// Exist matches all values which are present, even if nil
type Group struct {
	Field string
}

// Match implements Expression interface.
func (e Group) Match(payload map[string]interface{}) bool {
	_, found := getFieldExist(payload, e.Field)
	return found
}

// Prepare implements Expression interface.
func (e *Group) Prepare(validator schema.Validator) error {
	return validateField(e.Field, validator)
}

// String implements Expression interface.
func (e Group) String() string {
	return quoteField(e.Field) + ": {" + opExists + ": true}"
}
