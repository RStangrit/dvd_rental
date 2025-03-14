package actor

// import (
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// func Test_RegisterActorRoutes(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	r := gin.New()

// 	RegisterActorRoutes(r)

// 	routes := r.Routes()

// 	tests := []struct {
// 		method  string
// 		path    string
// 		handler string
// 	}{
// 		{"POST", "/actor", "PostActorHandler"},
// 		{"POST", "/actors", "PostActorsHandler"},
// 		{"GET", "/actors", "GetActorsHandler"},
// 		{"GET", "/actor/:id", "GetActorHandler"},
// 		{"GET", "/actor/:id/films", "GetActorFilmsHandler"},
// 		{"PUT", "/actor/:id", "PutActorHandler"},
// 		{"DELETE", "/actor/:id", "DeleteActorHandler"},
// 	}

// 	for _, test := range tests {
// 		if !containsRoute(routes, test.method, test.path, test.handler) {
// 			t.Errorf("Route not found: %s %s -> %s", test.method, test.path, test.handler)
// 		}
// 	}

// }

// func containsRoute(routes []gin.RouteInfo, method, path, handler string) bool {
// 	for _, r := range routes {
// 		if r.Method == method && r.Path == path && strings.HasSuffix(r.Handler, handler) {
// 			return true
// 		}
// 	}
// 	return false
// }
