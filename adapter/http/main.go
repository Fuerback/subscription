package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fuerback/subscription/adapter/http/rest/productservice"
	"github.com/Fuerback/subscription/adapter/sqlite"
	"github.com/Fuerback/subscription/adapter/sqlite/productrepository"
	"github.com/Fuerback/subscription/core/usecase/productusecase"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

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

	// Config Product
	productRepository := productrepository.New(conn)
	productUseCase := productusecase.New(productRepository)
	productService := productservice.New(productUseCase)

	router := mux.NewRouter()

	router.HandleFunc("/v1/product", productService.Fetch).Methods("GET")
	router.HandleFunc("/v1/product/{id}", productService.FetchOne).Methods("GET")

	port := viper.GetString("server.port")
	if port == "" {
		port = os.Getenv("PORT")
	}

	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)

}
