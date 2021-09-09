package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"main/handler"
)

func TestServer_Run(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Moscow",
		},
		{
			name: "Berlin",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/weather?city=Moscow", nil)
			response := httptest.NewRecorder()

			handler.WeatherAPI(response, request)

			got := response.Body.String()

			var jsonMap map[string]interface{}
			json.Unmarshal([]byte(got), &jsonMap)

			if jsonMap["main"] == nil {
				t.Errorf("no main structure")
			}

			if jsonMap["name"] != "Moscow" {
				t.Errorf("no name value")
			}

			if jsonMap["temp"] == nil {
				t.Errorf("no temp value")
			}
		})
	}
	t.Run("List", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/listRequests", nil)
		response := httptest.NewRecorder()

		handler.ListRequests(response, request)

		got := response.Body.String()

		var jsonMap map[string]interface{}
		json.Unmarshal([]byte(got), &jsonMap)

		if jsonMap["id"] == nil {
			t.Errorf("no id")
		}

		if jsonMap["city"] == nil {
			t.Errorf("no city value")
		}

		if jsonMap["temperature"] == nil {
			t.Errorf("no temperature value")
		}

		if jsonMap["time_requested"] == nil {
			t.Errorf("no requested time value")
		}
	})
}
