package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valedaniel/golang-product-manager/internal/handler"
	"github.com/valedaniel/golang-product-manager/internal/storage"
)

func main() {
	fmt.Println("Iniciando o servidor da API de produtos")

	logger := log.New(os.Stdout, "product-manager-API: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	err := godotenv.Load()

	if err != nil {
		logger.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	addr := os.Getenv("DB_ADDR")

	if addr == "" {
		addr = ":3000"
	}

	logger.Println("Conectando ao banco de dados...")

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	))

	if err != nil {
		logger.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logger.Fatalf("Não foi possível verificar a conexão com o banco de dados: %v", err)
	}

	logger.Println("Conexão com o banco de dados estabelecida com sucesso.")

	productStorage := storage.NewPostgresStore(db)

	productRouter := handler.NewRouter(productStorage, logger)

	srv := &http.Server{
		Addr:         addr,
		Handler:      productRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	logger.Printf("Iniciando o servidor na porta %s", addr)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Não foi possível iniciar o servidor: %v", err)
	}
}
