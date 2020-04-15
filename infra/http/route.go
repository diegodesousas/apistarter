package http

type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Middlewares []Middleware
}
