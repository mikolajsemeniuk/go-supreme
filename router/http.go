package router

import "net/http"

type HTTPRouter struct {
}

func (r *HTTPRouter) Route(listen string) error {
	http.ListenAndServe(listen, nil)
	return nil
}
