package router

type Router interface {
	Route(string) error
}
