package swb

import (
	"fmt"
	"testing"

	"github.com/xeipuuv/gojsonschema"

	"github.com/Starish-Wars-Backend/internal/swb/persistence"
)

// TestCreate test the Create function
func TestCreate(t *testing.T) {
	var schemaLoader = gojsonschema.NewReferenceLoader("file:///Users/johntodd/workspace/go/src/github.com/Starish-Wars-Backend/api/state_schema.json")

	mockPersister := persistence.MockPersister{}
	gameID, response, err := Create(mockPersister)
	if err != nil {
		t.Errorf("Did not expect an error")
	}
	if gameID == "" {
		t.Errorf("gameID should not be empty")
	}
	responseLoader := gojsonschema.NewStringLoader(response)
	result, parseErr := gojsonschema.Validate(schemaLoader, responseLoader)
	if parseErr != nil {
		t.Errorf("Did not expect an error while parsing")
		t.Errorf("%s", parseErr)
		return
	}
	if !result.Valid() {
		fmt.Printf("The response is not valid. See errors :\n")
		for _, err := range result.Errors() {
			fmt.Printf("- %s]n", err)
		}
	}
}
