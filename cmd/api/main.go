package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dstreet/mogal/internal/auth"
	"github.com/dstreet/mogal/internal/db"
	"github.com/dstreet/mogal/internal/graphql"
	mhttp "github.com/dstreet/mogal/internal/http"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/net/context"

	_ "github.com/lib/pq"
)

const defaultPort = "8080"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Warn("no .env file found")
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

	cs := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
	)

	// conn, err := sql.Open("postgres", cs)
	conn, err := pgxpool.New(context.TODO(), cs)
	if err != nil {
		panic(err)
	}

	userRepo := db.NewDBUserRepository(
		logger.WithGroup("UserRepository"),
		conn,
		auth.NewBcryptPasswordHasher(12),
	)

	genreRepo := db.NewDBGenreRepository(
		logger.WithGroup("GenreRepository"),
		conn,
	)

	movieRepo := db.NewDBMovieRepository(
		logger.WithGroup("MovieRepository"),
		conn,
	)

	tokenProvider := auth.NewJWTProvider(authIssuer, authSigningKey, time.Second*900)

	authMiddleware := &mhttp.AuthMiddleware{
		Logger:         logger.WithGroup("Auth Middleware"),
		UserRepository: userRepo,
		TokenProvider:  tokenProvider,
	}

	resolver := &graphql.Resolver{
		Logger:          logger.WithGroup("Resolver"),
		UserRepository:  userRepo,
		GenreRepository: genreRepo,
		MovieRepository: movieRepo,
		TokenProvider:   tokenProvider,
		FieldCollector:  &graphql.GraphQLFieldCollector{},
	}

	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", mhttp.CORSMiddleware(authMiddleware.Handler(srv)))

	logger.Info("API started", "endpoint", fmt.Sprintf("http://localhost:%s/graphql", port))
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		logger.Error("failed to start api server", "err", err)
	}
}
