package catalog

import "net/http"

type CatalogRouter interface {
	Route() *http.ServeMux
}

type CatalogRouterImpl struct {
	Handler CatalogHandler
}

func NewCatalogRouter(handler CatalogHandler) CatalogRouter {
	return &CatalogRouterImpl{Handler: handler}
}

func (r *CatalogRouterImpl) Route() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/v1/categories", r.Handler.List)
	mux.HandleFunc("GET /api/v1/categories/{id}", r.Handler.Get)
	mux.HandleFunc("POST /api/v1/categories", r.Handler.Create)

	return mux
}
