package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

type GraphQLFieldCollector struct{}

func (fc *GraphQLFieldCollector) CollectAllFields(ctx context.Context) []string {
	return graphql.CollectAllFields(ctx)
}
