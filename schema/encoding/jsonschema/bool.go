package jsonschema

import "github.com/oktacode/rest-layer/schema"

type boolBuilder schema.Bool

func (v boolBuilder) BuildJSONSchema() (map[string]interface{}, error) {
	return map[string]interface{}{"type": "boolean"}, nil
}
