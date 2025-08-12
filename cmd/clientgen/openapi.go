package main

import (
	"sort"
	"strings"
)

type OpenAPISpec struct {
	OpenAPI    string                         `yaml:"openapi"`
	Info       Info                           `yaml:"info"`
	Servers    []Server                       `yaml:"servers"`
	Paths      map[string]PathItem            `yaml:"paths"`
	Components Components                     `yaml:"components"`
}

type Info struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
}

type Server struct {
	URL string `yaml:"url"`
}

type PathItem struct {
	Get    *Operation `yaml:"get,omitempty"`
	Post   *Operation `yaml:"post,omitempty"`
	Put    *Operation `yaml:"put,omitempty"`
	Delete *Operation `yaml:"delete,omitempty"`
}

type Operation struct {
	OperationID string                        `yaml:"operationId"`
	Summary     string                        `yaml:"summary"`
	Description string                        `yaml:"description"`
	Tags        []string                      `yaml:"tags"`
	Parameters  []Parameter                   `yaml:"parameters"`
	RequestBody *RequestBody                  `yaml:"requestBody,omitempty"`
	Responses   map[string]Response           `yaml:"responses"`
}

type Parameter struct {
	Name        string  `yaml:"name"`
	In          string  `yaml:"in"`
	Description string  `yaml:"description"`
	Required    bool    `yaml:"required"`
	Schema      *Schema `yaml:"schema"`
}

type RequestBody struct {
	Description string               `yaml:"description"`
	Required    bool                 `yaml:"required"`
	Content     map[string]MediaType `yaml:"content"`
}

type Response struct {
	Description string               `yaml:"description"`
	Content     map[string]MediaType `yaml:"content"`
}

type MediaType struct {
	Schema *Schema `yaml:"schema"`
}

type Components struct {
	Schemas map[string]Schema `yaml:"schemas"`
}

type Schema struct {
	Type        string             `yaml:"type"`
	Format      string             `yaml:"format"`
	Description string             `yaml:"description"`
	Properties  map[string]Schema  `yaml:"properties"`
	Items       *Schema            `yaml:"items"`
	Enum        []interface{}      `yaml:"enum"`
	Required    []string           `yaml:"required"`
	Ref         string             `yaml:"$ref"`
	Example     interface{}        `yaml:"example"`
	AllOf       []Schema           `yaml:"allOf"`
	OneOf       []Schema           `yaml:"oneOf"`
	AnyOf       []Schema           `yaml:"anyOf"`
}

func (s *Schema) GoType() string {
	if s.Ref != "" {
		parts := strings.Split(s.Ref, "/")
		return toPascalCase(parts[len(parts)-1])
	}

	// Handle allOf - typically used for inheritance or combining schemas
	if len(s.AllOf) > 0 {
		// For allOf, we'll use the first reference if available
		for _, schema := range s.AllOf {
			if schema.Ref != "" {
				parts := strings.Split(schema.Ref, "/")
				return toPascalCase(parts[len(parts)-1])
			}
		}
	}

	switch s.Type {
	case "string":
		if len(s.Enum) > 0 {
			return "string"
		}
		switch s.Format {
		case "date-time":
			return "DateTime"
		case "date":
			return "Date"
		default:
			return "string"
		}
	case "integer":
		switch s.Format {
		case "int32":
			return "int32"
		case "int64":
			return "int64"
		default:
			return "int"
		}
	case "number":
		switch s.Format {
		case "float":
			return "float32"
		case "double":
			return "float64"
		default:
			return "float64"
		}
	case "boolean":
		return "bool"
	case "array":
		if s.Items != nil {
			return "[]" + s.Items.GoType()
		}
		return "[]interface{}"
	case "object":
		if len(s.Properties) == 0 {
			return "map[string]interface{}"
		}
		return "interface{}"
	default:
		return "interface{}"
	}
}

func (s *Schema) IsRequired(fieldName string) bool {
	for _, req := range s.Required {
		if req == fieldName {
			return true
		}
	}
	return false
}

func toPascalCase(s string) string {
	if s == "" {
		return ""
	}
	
	words := strings.FieldsFunc(s, func(c rune) bool {
		return c == '_' || c == '-' || c == ' '
	})
	
	for i, word := range words {
		if word != "" {
			words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
		}
	}
	
	return strings.Join(words, "")
}

func toCamelCase(s string) string {
	if s == "" {
		return ""
	}
	
	pascal := toPascalCase(s)
	if pascal == "" {
		return ""
	}
	
	return strings.ToLower(string(pascal[0])) + pascal[1:]
}

func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && (r >= 'A' && r <= 'Z') {
			result.WriteByte('_')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

func (spec *OpenAPISpec) GetSortedSchemas() []string {
	var names []string
	for name := range spec.Components.Schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (spec *OpenAPISpec) GetSortedPaths() []string {
	var paths []string
	for path := range spec.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	return paths
}

func (op *Operation) GetMethodName() string {
	if op.OperationID != "" {
		return toPascalCase(strings.ReplaceAll(op.OperationID, "-", "_"))
	}
	return "UnknownOperation"
}

func (op *Operation) GetSuccessResponse() *Response {
	for code, response := range op.Responses {
		if strings.HasPrefix(code, "2") {
			return &response
		}
	}
	return nil
}