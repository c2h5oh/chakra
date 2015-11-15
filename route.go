package chakra

import (
	"net/http"
	"strings"

	"github.com/pressly/chi"
	"golang.org/x/net/context"
)

// Route builds a resource identifier from the middleware chain. This resource
// identifier along with the operation (HTTP verb) can be used for determining
// access to a resource
func Route(part string) func(chi.Handler) chi.Handler {
	return func(next chi.Handler) chi.Handler {
		hn := func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
			part = extractACLRoute(part)
			if part != "" {
				route, ok := ctx.Value("acl.route").([]string)
				if ok {
					route = append(route, part)
				} else {
					route = []string{part}
				}
				ctx = context.WithValue(ctx, "acl.route", route)
			}
			next.ServeHTTPC(ctx, w, r)
		}
		return chi.HandlerFunc(hn)
	}
}

// extractACLRoute strips leading and trailing slashes as well as ":" url param
// prefix
func extractACLRoute(s string) string {
	return strings.Trim(s, "/:")
}
