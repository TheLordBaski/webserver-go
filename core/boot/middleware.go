// Package boot handles the initialization of the web components.
package boot

import (
	"net/http"
	"../router"
	"github.com/gorilla/context"
)

// SetUpMiddleware contains the middleware that applies to every request.
func SetUpMiddleware(h http.Handler) http.Handler {
	return router.ChainHandler( // Chain middleware, top middlware runs first
		h,                    // Handler to wrap
		//TODO add more chains fe logger rest handler etc
		context.ClearHandler, // Prevent memory leak with gorilla.sessions
	)
}
