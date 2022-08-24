package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fuerback/subscription/adapter/http/rest/productservice"
	"github.com/Fuerback/subscription/adapter/http/rest/subscriptionservice"
	"github.com/Fuerback/subscription/adapter/sqlite"
	"github.com/Fuerback/subscription/adapter/sqlite/productrepository"
	"github.com/Fuerback/subscription/adapter/sqlite/subscriptionrepository"
	"github.com/Fuerback/subscription/core/usecase/productusecase"
	"github.com/Fuerback/subscription/core/usecase/subscriptionusecase"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

// use a single instance of Validate, it caches struct info
var jsonValidate *validator.Validate

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := sqlite.GetConnection(ctx)

	jsonValidate = validator.New()

	// Config Product
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase, jsonValidate)

	// Config Subscription
	subscriptionRepository := subscriptionrepository.New(conn)
	subscriptionUseCase := subscriptionusecase.New(subscriptionRepository)
	subscriptionService := subscriptionservice.New(subscriptionUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/v1/product", productService.Fetch).Methods("GET")
	router.HandleFunc("/v1/product/{id}", productService.FetchOne).Methods("GET")
	router.HandleFunc("/v1/product/purchase/{id}", productService.Purchase).Methods("POST")

	router.HandleFunc("/v1/subscription/{id}", subscriptionService.FetchOne).Methods("GET")
	router.HandleFunc("/v1/subscription/cancel/{id}", subscriptionService.UpdateStatus).Methods("PUT")
	router.HandleFunc("/v1/subscription/status/{id}", subscriptionService.UpdateStatus).Methods("PUT")

	port := viper.GetString("server.port")
	if port == "" {
		port = os.Getenv("PORT")
	}

	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)

}
