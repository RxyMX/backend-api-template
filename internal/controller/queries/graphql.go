package queries

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/graphql-go/graphql/language/kinds"
	"github.com/graphql-go/graphql/language/parser"
	"github.com/graphql-go/graphql/language/printer"
	"github.com/graphql-go/graphql/language/source"
	"log"
)

type QueryRequest struct {
	Query         string                     `json:"query"`
	OperationName string                     `json:"operationName,omitempty"`
	Variables     map[string]json.RawMessage `json:"variables,omitempty"`
}

func CreateQueryRequest(queryBytes []byte) (QueryRequest, error) {
	queryRequest := QueryRequest{}
	return queryRequest, json.Unmarshal(queryBytes, &queryRequest)
}

type QueryJson struct {
	Type          string                     `json:"type"`
	OperationName string                     `json:"operationName"`
	Arguments     map[string]interface{}     `json:"arguments,omitempty"`
	Fields        []string                   `json:"fields,omitempty"`
	Variables     map[string]json.RawMessage `json:"variables,omitempty"`
	Data          json.RawMessage            `json:"data,omitempty"`
}

type Query struct {
	*QueryJson
	*ast.OperationDefinition
}

func (q Query) ToQueryRequest() QueryRequest {
	return QueryRequest{
		Query:     printer.Print(q.OperationDefinition).(string),
		Variables: q.Variables,
	}
}

func ToGraphqlQuery(query []byte) ([]Query, error) {
	var queries []Query
	document, err := parser.Parse(parser.ParseParams{
		Source: source.NewSource(&source.Source{
			Body: query,
		}),
	})

	if err == nil {
		if len(document.Definitions) == 1 {
			node := document.Definitions[0]
			if node.GetKind() == kinds.OperationDefinition {
				operation := node.(*ast.OperationDefinition)
				for _, selection := range operation.SelectionSet.Selections {
					queries = append(queries, createQuery(operation, selection.(*ast.Field)))
				}
			} else {
				err = errors.New("Found no operation definition in the graphql query")
			}
		} else {
			err = errors.New(
				fmt.Sprintf("Only support graphql defintions of size 1, but found %v",
					len(document.Definitions)))
		}
	}

	return queries, err
}

func createQuery(operation *ast.OperationDefinition, field *ast.Field) Query {
	var query = Query{
		QueryJson: &QueryJson{
			Type:          field.Name.Value,
			OperationName: operation.Operation,
		},
		OperationDefinition: operation,
	}

	applyFields("", field.GetSelectionSet().Selections, &query)

	if len(field.Arguments) > 0 {
		query.Arguments = make(map[string]interface{})
		applyArguments(field.Arguments, &query)
	}

	return query
}

func applyFields(parent string, selections []ast.Selection, query *Query) {
	for _, selection := range selections {
		name := selection.(*ast.Field).Name.Value
		if parent != "" {
			name = parent + "." + name
		}

		query.Fields = append(query.Fields, name)

		if selection.GetSelectionSet() != nil {
			applyFields(name, selection.GetSelectionSet().Selections, query)
		}
	}
}

func applyArguments(arguments []*ast.Argument, query *Query) {
	for _, argument := range arguments {
		query.Arguments[argument.Name.Value] = parseLiteral(argument.Value, &query.Variables)
	}
}

func parseLiteral(astValue ast.Value, variables *map[string]json.RawMessage) interface{} {
	kind := astValue.GetKind()

	switch kind {
	case kinds.StringValue:
		return astValue.GetValue()
	case kinds.BooleanValue:
		return astValue.GetValue()
	case kinds.IntValue:
		return astValue.GetValue()
	case kinds.FloatValue:
		return astValue.GetValue()
	case kinds.EnumValue:
		return astValue.GetValue()
	case kinds.Variable:
		if *variables == nil {
			*variables = make(map[string]json.RawMessage)
		}

		variableName := astValue.GetValue().(*ast.Name).Value
		(*variables)[variableName] = []byte("")
		return variableName
	case kinds.ObjectValue:
		obj := make(map[string]interface{})
		for _, v := range astValue.GetValue().([]*ast.ObjectField) {
			obj[v.Name.Value] = parseLiteral(v.Value, variables)
		}
		return obj
	case kinds.ListValue:
		list := make([]interface{}, 0)
		for _, v := range astValue.GetValue().([]ast.Value) {
			list = append(list, parseLiteral(v, variables))
		}
		return list
	default:
		log.Printf("Unsupported parse for value type: %v", kind)
		return nil
	}
}
