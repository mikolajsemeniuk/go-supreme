package router

import "net/http"

type Router interface {
	Route(string) error
}

type ListHandler interface {
	HandleList(w http.ResponseWriter, r *http.Request)
}

type ReadHandler interface {
	HandleRead(w http.ResponseWriter, r *http.Request)
}

type CreateHandler interface {
	HandleCreate(w http.ResponseWriter, r *http.Request)
}

type UpdateHandler interface {
	HandleUpdate(w http.ResponseWriter, r *http.Request)
}

type RemoveHandler interface {
	HandleRemove(w http.ResponseWriter, r *http.Request)
}
