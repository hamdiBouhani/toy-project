package gql

// ExtraQueries is a list of GraphQL custom queries.
const ExtraQueries = `
	ping: String!
`

// ExtraMutations is a list of GraphQL custom mutations.
const ExtraMutations = ``

// ExtraTypes is a list of GraphQL custom types.
const ExtraTypes = ``

// Ping returns the string "pong".
func (r RootResolver) Ping() string {
	return "pong"
}
