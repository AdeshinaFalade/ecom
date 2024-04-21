package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/adeshinafalade/ecom/service/cart"
	"github.com/adeshinafalade/ecom/service/order"
	"github.com/adeshinafalade/ecom/service/products"
	"github.com/adeshinafalade/ecom/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	productStore := products.NewStore(s.db)
	productHandler := products.NewHandler(productStore)
	productHandler.RegisterRoutes(subRouter)

	orderStore := order.NewStore(s.db)
	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
