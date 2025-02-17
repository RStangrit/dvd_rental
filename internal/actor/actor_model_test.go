package actor

import (
	"main/internal/models"
	"testing"
)

func Test_TableName(t *testing.T) {
	var actor Actor
	tableNameFromFunction := actor.TableName()

	if tableNameFromFunction != tableName {
		t.Errorf("TableName method is compromised")
	}
}

func Test_RegisterModels(t *testing.T) {
	registerModels()

	if !models.IsModelRegistered(&Actor{}) {
		t.Errorf("Actor model was not registered")
	}
}
