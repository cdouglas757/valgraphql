package helpers

import (
	"strings"

	linq "github.com/ahmetb/go-linq/v3"

	"github.com/cdouglas757/valgraphql/graph/customTypes"
	apiModels "github.com/cdouglas757/valgraphql/valapi/models"
)

func GetAgents(getApiAgents func() []apiModels.Agent) []*customTypes.Agent {
	apiAgents := getApiAgents()

	agents := []*customTypes.Agent{}

	for _, a := range apiAgents {
		agent := &customTypes.Agent{
			Name: a.Name,
			UUID: a.Uuid,
			Role: a.Role.Name,
		}
		agents = append(agents, agent)
	}

	return agents
}

func FilterByName(agents []*customTypes.Agent, name string) []*customTypes.Agent {
	if name == "" {
		return agents
	}

	var nameMatch []*customTypes.Agent
	linq.From(agents).WhereT(func(a *customTypes.Agent) bool {
		return strings.EqualFold(a.Name, name)
	}).ToSlice(&nameMatch)

	return nameMatch
}

func FilterByRole(agents []*customTypes.Agent, role string) []*customTypes.Agent {
	if role == "" {
		return agents
	}

	var roleMatch []*customTypes.Agent
	linq.From(agents).WhereT(func(a *customTypes.Agent) bool {
		return strings.EqualFold(a.Role, role)
	}).ToSlice(&roleMatch)

	return roleMatch
}
