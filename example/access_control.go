package example

import (
	"log"
	"net/http"
	"strings"

	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

// AccessControl is an example that just prints the ACL route + operation that
// is being checked without actually doing any checking
func AccessControl(next chi.Handler) chi.Handler {
	hn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
		route, ok := ctx.Value("acl.route").([]string)
		if !ok {
			http.Error(w, "undefined acl route", 403)
			return
		}
		// Put ACL code here
		log.Printf("Checking permission to %s %s", r.Method, strings.Join(route, " -> "))

		next.ServeHTTPC(ctx, w, r)
	}
	return chi.HandlerFunc(hn)
}
