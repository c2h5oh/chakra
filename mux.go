package chakra

import "github.com/pressly/chi"

// UseParentRoute passed as route to NewRouter results in nothing being added to
// acl.route on subrouter creation - parent route is used instead
const UseParentRoute = "#"

// Mux implements chi.Router and makes access control with chi a bit easier
type Mux struct {
	*chi.Mux
}

var accessControl interface{}

// SetAC sets the function that is going to be used by chakra for resource
// control
func SetAC(fn func(chi.Handler) chi.Handler) {
	accessControl = fn
}

// Use appends middleware(s) to chi middleware stack
func (r *Mux) Use(middlewares ...interface{}) {
	r.Mux.Use(middlewares...)
}

// Handle is a wrapper of chi.Handle for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Handle(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Handle(pattern, h...)
}

// Connect is a wrapper of chi.Connect for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Connect(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Connect(pattern, h...)
}

// Head is a wrapper of chi.Head for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Head(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Head(pattern, h...)
}

// Get is a wrapper of chi.Get for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Get(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Get(pattern, h...)
}

// Post is a wrapper of chi.Post for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Post(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Post(pattern, h...)
}

// Put a wrapper of chi.Put for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Put(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Put(pattern, h...)
}

// Patch is a wrapper of chi.Patch for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Patch(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Patch(pattern, h...)
}

// Delete is a wrapper of chi.Delete for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Delete(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Delete(pattern, h...)
}

// Trace is a wrapper of chi.Trace for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Trace(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Trace(pattern, h...)
}

// Options is a wrapper of chi.Options for chakra.Mux. It appends the pattern
// passed to it to acl.route and injects access control middleware as
// penultimate one - just before handler(s)
func (r *Mux) Options(pattern string, handlers ...interface{}) {
	h := append([]interface{}{Route(pattern), accessControl}, handlers...)
	r.Mux.Options(pattern, h...)
}

// Group is a wrapper of chi.Group for chakra.Mux. It works exactly the same
// as the underlying chi method
func (r *Mux) Group(fn func(r chi.Router)) chi.Router {
	return r.Mux.Group(fn)
}

// Route is a wrapper of chi.Route for chakra.Mux. It works exactly the same
// as the underlying chi method
func (r *Mux) Route(pattern string, fn func(r chi.Router)) chi.Router {
	subRouter := NewRouter(UseParentRoute)
	r.Mount(pattern, subRouter)
	if fn != nil {
		fn(subRouter)
	}
	return subRouter
}

// Mount is a wrapper of chi.Mount for chakra.Mux. It works exactly the same
// as the underlying chi method
func (r *Mux) Mount(path string, handlers ...interface{}) {
	r.Mux.Mount(path, handlers...)
}

// NewRouter creates new chakra Mux and appends its route to acl.route, unless
// the passed route is chakra.UseParentRoute.
func NewRouter(route string) *Mux {
	if route == "" {
		panic("new router created without acl route")
	}
	if accessControl == nil {
		panic("you have not defined access control function with SetAC")
	}

	r := &Mux{chi.NewRouter()}
	if route != UseParentRoute {
		r.Mux.Use(Route(route))
	}

	return r
}
