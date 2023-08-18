package valapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	models "github.com/cdouglas757/valgraphql/valapi/models"
)

func GetAllAgents() []models.Agent {
	response, err := http.Get("https://valorant-api.com/v1/agents?isPlayableCharacter=true")

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var agentResponse models.AgentResponse
	jsonErr := json.Unmarshal(responseData, &agentResponse)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return agentResponse.Data
}
