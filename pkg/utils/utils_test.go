package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_JoinStrings(t *testing.T) {
	string1 := "string1"
	string2 := "string2"
	result := JoinStrings(string1, string2)

	expected := "string1 string2"

	if result != expected {
		t.Errorf("JoinStrings(%s, %s) = %s, expected %s", string1, string2, result, expected)
	}

}

func Test_GetIntParam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		paramName  string
		paramValue string
		expected   int64
		expectErr  bool
	}{
		{"positive", "id", "1", 1, false},
		{"negative", "id", "-1", -1, false},
		{"string", "id", "first", 0, true},
		{"empty", "id", "", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			writer := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(writer)
			context.AddParam(test.paramName, test.paramValue)

			result, err := GetIntParam(context, "id")

			if test.expectErr {
				if err == nil {
					t.Fatalf("expected error but got none for input: %s", test.paramValue)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != test.expected {
				t.Errorf("GetIntParam(%s, %s) = %d, expected %d", test.paramName, test.paramValue, result, test.expected)
			}

		})

	}
}
