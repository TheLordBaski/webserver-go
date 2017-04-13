package router

import (
	"net/http"
	"sync"
	// Middleware chaining
	"github.com/justinas/alice"
)

var(
	routeList []string
	listMutex sync.RWMutex
)

func ChainHandler(h http.Handler, c ...alice.Constructor) http.Handler {
	return alice.New(c...).Then(h)
}