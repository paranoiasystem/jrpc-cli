package model

import "encoding/json"

type JsonSchema struct {
	Schema     string                `json:"$schema"`
	Vocabulary map[string]bool       `json:"$vocabulary"`
	ID         string                `json:"$id"`
	Ref        string                `json:"$ref"`
	DynamicRef string                `json:"$dynamicRef"`
	Defs       map[string]JsonSchema `json:"$defs"`
	Comment    string                `json:"$comment"`

	AllOf []JsonSchema `json:"allOf"`
	AnyOf []JsonSchema `json:"anyOf"`
	OneOf []JsonSchema `json:"oneOf"`
	Not   []JsonSchema `json:"not"`

	If               *JsonSchema           `json:"if"`
	Then             *JsonSchema           `json:"then"`
	Else             *JsonSchema           `json:"else"`
	DependentSchemas map[string]JsonSchema `json:"dependentSchemas"`

	PrefixItems []JsonSchema `json:"prefixItems"`
	Items       *JsonSchema  `json:"items"`
	Contains    *JsonSchema  `json:"contains"`

	Properties           map[string]JsonSchema `json:"properties"`
	PatternProperties    map[string]JsonSchema `json:"patternProperties"`
	AdditionalProperties interface{}           `json:"additionalProperties"`
	PropertyNames        *JsonSchema           `json:"propertyNames"`

	Type  interface{}   `json:"type"`
	Enum  []interface{} `json:"enum"`
	Const interface{}   `json:"const"`

	MultipleOf       json.Number `json:"multipleOf"`
	Maximum          json.Number `json:"maximum"`
	ExclusiveMaximum json.Number `json:"exclusiveMaximum"`
	Minimum          json.Number `json:"minimum"`
	ExclusiveMinimum json.Number `json:"exclusiveMinimum"`

	MaxLength int    `json:"maxLength"`
	MinLength int    `json:"minLength"`
	Pattern   string `json:"pattern"`

	MaxItems    int  `json:"maxItems"`
	MinItems    int  `json:"minItems"`
	UniqueItems bool `json:"uniqueItems"`
	MaxContains int  `json:"maxContains"`
	MinContains int  `json:"minContains"`

	MaxProperties     int                 `json:"maxProperties"`
	MinProperties     int                 `json:"minProperties"`
	Required          []string            `json:"required"`
	DependentRequired map[string][]string `json:"dependentRequired"`

	Title       string        `json:"title"`
	Description string        `json:"description"`
	Default     interface{}   `json:"default"`
	Deprecated  bool          `json:"deprecated"`
	ReadOnly    bool          `json:"readOnly"`
	WriteOnly   bool          `json:"writeOnly"`
	Examples    []interface{} `json:"examples"`
}
