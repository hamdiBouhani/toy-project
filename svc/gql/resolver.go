package gql

import (
	"strings"

	"github.com/sirupsen/logrus"
)

type RootResolver struct{}

func GetRootSchema(extraQueries, extraMutations, extraTypes string) string {
	return `
	schema {
		query: Query
		mutation: Mutation
	}

	type Query {
   ` + extraQueries +
		`}
   
   type Mutation {
	` + extraMutations +
		`}
	
` + extraTypes

}

// RootResolvers wrapper for model's RootResolver
type RootResolvers struct {
	RootResolver
	Logger logrus.FieldLogger
}

type extraGraphQL struct {
	queries   []string
	mutations []string
	types     []string
}

func (e *extraGraphQL) queryRegister(query string) {
	if e.queries == nil {
		e.queries = make([]string, 0)
	}
	e.queries = append(e.queries, query)
}

func (e *extraGraphQL) mutationRegister(mutation string) {
	if e.mutations == nil {
		e.mutations = make([]string, 0)
	}
	e.mutations = append(e.mutations, mutation)
}

func (e *extraGraphQL) typeRegister(typ string) {
	if e.types == nil {
		e.types = make([]string, 0)
	}
	e.types = append(e.types, typ)
}

func (e *extraGraphQL) getQueries() string {
	return strings.Join(e.queries, "\n")
}

func (e *extraGraphQL) getMutations() string {
	return strings.Join(e.mutations, "\n")
}

func (e *extraGraphQL) getTypes() string {
	return strings.Join(e.types, "\n")
}

var extraGQL extraGraphQL

// GetRootSchemas get wrapper model's schema, and self defined graphQL query/mutation
func GetRootSchemas() string {
	return GetRootSchema(extraGQL.getQueries(), extraGQL.getMutations(), extraGQL.getTypes())
}

// TODO: refactor existed queries/mutations/types to current package
func init() {
	extraGQL.queryRegister(ExtraQueries)
	extraGQL.mutationRegister(ExtraMutations)
	extraGQL.typeRegister(ExtraTypes)
}
