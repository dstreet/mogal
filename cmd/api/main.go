package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dstreet/mogal/internal/auth"
	"github.com/dstreet/mogal/internal/db"
	"github.com/dstreet/mogal/internal/graphql"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		panic("missing DB_HOST")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		panic("missing DB_NAME")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		panic("missing DB_USER")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		panic("missing DB_PASSWORD")
	}

	authIssuer := os.Getenv("AUTH_ISSUER")
	if authIssuer == "" {
		panic("missing AUTH_ISSUER")
	}

	authSigningKey := os.Getenv("AUTH_SIGNING_KEY")
	if authSigningKey == "" {
		panic("missing AUTH_SIGNING_KEY")
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	cs := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
	)

	fmt.Println(cs)
	conn, err := sql.Open("postgres", cs)
	if err != nil {
		panic(err)
	}

	resolver := &graphql.Resolver{
		Logger: logger.WithGroup("Resolver"),
		UserRepository: db.NewDBUserRepository(
			logger.WithGroup("UserRepository"),
			conn,
			auth.NewBcryptPasswordHasher(12),
		),
		TokenProvider: auth.NewJWTProvider(authIssuer, authSigningKey),
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("connect to http://localhost:%s/ for GraphQL playground", "port", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Error("failed to start api server", "err", err)
	}
}
