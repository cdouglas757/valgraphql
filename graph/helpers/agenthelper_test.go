package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cdouglas757/valgraphql/graph/customTypes"
	apiModels "github.com/cdouglas757/valgraphql/valapi/models"
)

func TestAgentHelper_GetAgents_ReturnsAgents(t *testing.T) {
	//arrange
	testName1 := "test1"
	testUid2 := "uuid2"
	testRole3 := "role3"
	apiAgents := []apiModels.Agent{
		{
			Name: testName1,
			Uuid: "uuid1",
			Role: apiModels.Role{
				Name: "role1",
			},
		},
		{
			Name: "test2",
			Uuid: testUid2,
			Role: apiModels.Role{
				Name: "role2",
			},
		},
		{
			Name: "test3",
			Uuid: "uuid3",
			Role: apiModels.Role{
				Name: testRole3,
			},
		},
	}

	//act
	agents := GetAgents(func() []apiModels.Agent {
		return apiAgents
	})

	//assert
	assert.NotNil(t, agents)
	assert.Equal(t, 3, len(agents))
	assert.Equal(t, testName1, agents[0].Name)
	assert.Equal(t, testUid2, agents[1].UUID)
	assert.Equal(t, testRole3, agents[2].Role)
}

func TestAgentHelper_FilterByName_EmptyStringReturnsOriginal(t *testing.T) {
	//arrange
	agents := []*customTypes.Agent{
		{
			UUID: "uuid1",
			Name: "test1",
			Role: "role1",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
	}
	agentName := ""

	//act
	returnedAgents := FilterByName(agents, agentName)

	//assert
	assert.Equal(t, agents, returnedAgents)
}

func TestAgentHelper_FilterByName_ReturnsCorrectValue(t *testing.T) {
	//arrange
	agents := []*customTypes.Agent{
		{
			UUID: "uuid1",
			Name: "test1",
			Role: "role1",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
		{
			UUID: "uuid3",
			Name: "test3",
			Role: "role3",
		},
	}
	agentName := "test2"

	//act
	returnedAgents := FilterByName(agents, agentName)

	//assert
	assert.Equal(t, 1, len(returnedAgents))
	assert.Equal(t, agentName, returnedAgents[0].Name)
}

func TestAgentHelper_FilterByRole_EmptyStringReturnsOriginal(t *testing.T) {
	//arrange
	agents := []*customTypes.Agent{
		{
			UUID: "uuid1",
			Name: "test1",
			Role: "role1",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
	}
	agentRole := ""

	//act
	returnedAgents := FilterByRole(agents, agentRole)

	//assert
	assert.Equal(t, agents, returnedAgents)
}

func TestAgentHelper_FilterByRole_ReturnsCorrectValue(t *testing.T) {
	//arrange
	agents := []*customTypes.Agent{
		{
			UUID: "uuid1",
			Name: "test1",
			Role: "role1",
		},
		{
			UUID: "uuid2",
			Name: "test2",
			Role: "role2",
		},
		{
			UUID: "uuid3",
			Name: "test3",
			Role: "role3",
		},
	}
	agentRole := "role3"

	//act
	returnedAgents := FilterByRole(agents, agentRole)

	//assert
	assert.Equal(t, 1, len(returnedAgents))
	assert.Equal(t, agentRole, returnedAgents[0].Role)
}
