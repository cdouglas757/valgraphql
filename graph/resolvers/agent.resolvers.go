package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"strings"

	"github.com/cdouglas757/valgraphql/graph/customTypes"
	"github.com/cdouglas757/valgraphql/graph/generated"
	"github.com/cdouglas757/valgraphql/graph/helpers"
	apiModels "github.com/cdouglas757/valgraphql/valapi/models"
	apiWrappers "github.com/cdouglas757/valgraphql/valapi/wrappers"
	"golang.org/x/exp/slices"
)

// GetAgentByName is the resolver for the getAgentByName field.
func (r *queryResolver) GetAgentByName(ctx context.Context, agentName string) (*customTypes.Agent, error) {
	agents := apiWrappers.GetAllAgents()

	idx := slices.IndexFunc(agents, func(a apiModels.Agent) bool { return strings.EqualFold(a.Name, agentName) })

	if idx < 0 {
		return nil, nil
	}

	agent := &customTypes.Agent{
		UUID: agents[idx].Uuid,
		Name: agents[idx].Name,
		Role: agents[idx].Role.Name,
	}

	return agent, nil
}

// Agents is the resolver for the Agents field.
func (r *queryResolver) Agents(ctx context.Context, agentName *string, role *string) ([]*customTypes.Agent, error) {
	agents := helpers.GetAgents(apiWrappers.GetAllAgents)

	if agentName != nil {
		n := *agentName
		agents = helpers.FilterByName(agents, n)
	}

	if role != nil {
		r := *role
		agents = helpers.FilterByRole(agents, r)
	}

	return agents, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }