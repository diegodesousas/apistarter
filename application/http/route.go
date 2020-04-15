package http

type Route struct {
	Path        string
	Method      string
	Handler     Handler
	Middlewares []Middleware
}

type TxRoute struct {
	Path        string
	Method      string
	Handler     TxHandler
	Middlewares []Middleware
}
