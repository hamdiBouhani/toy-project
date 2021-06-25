package gql

import (
	"context"

	"github.com/EchoUtopia/zerror"
	gqlErrors "github.com/graph-gophers/graphql-go/errors"
	"github.com/graph-gophers/graphql-go/introspection"
	"github.com/graph-gophers/graphql-go/trace"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type GQLTrace struct {
	Logger logrus.FieldLogger
}

func (g *GQLTrace) TraceQuery(ctx context.Context, queryString string, operationName string, variables map[string]interface{}, varTypes map[string]*introspection.Type) (context.Context, trace.TraceQueryFinishFunc) {

	return ctx, func(errs []*gqlErrors.QueryError) {}
}

func (g *GQLTrace) TraceField(ctx context.Context, label, typeName, fieldName string, trivial bool, args map[string]interface{}) (context.Context, trace.TraceFieldFinishFunc) {

	return ctx, func(err *gqlErrors.QueryError) {
		if err != nil {
			g.Logger.WithFields(logrus.Fields{"graphql.type": typeName, "graphql.field": fieldName, "graphql.error": err.Error()}).Errorln("GraphQL request failed")
			logrus.WithFields(logrus.Fields(args)).Errorln("args details")
			zerr := new(zerror.Error)
			if ok := errors.As(err.ResolverError, &zerr); ok {
				err.Extensions = map[string]interface{}{
					`code`:    zerr.Def.Code,
					`message`: err.Error(),
				}
			}
		}

	}

}
