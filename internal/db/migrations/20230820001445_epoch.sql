-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "last_login" timestamp NOT NULL
);

CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");

CREATE TABLE "movies" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "title" text NOT NULL,
  "rating" text NOT NULL,
  "cast" text[] NOT NULL,
  "director" text NOT NULL,
  "poster" text,
  "user" uuid NOT NULL,
  "user_rating" int
);

ALTER TABLE "movies" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

CREATE TABLE "user_movies" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user" uuid NOT NULL,
  "movie" uuid NOT NULL
);

ALTER TABLE "user_movies" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");
ALTER TABLE "user_movies" ADD FOREIGN KEY ("movie") REFERENCES "movies" ("id");

CREATE TABLE "genres" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "user" uuid NOT NULL,
  "name" text NOT NULL
);

CREATE UNIQUE INDEX "idx_genre_name" ON "genres" ("name");
ALTER TABLE "genres" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

CREATE TABLE "movie_genres" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "movie" uuid NOT NULL,
  "genre" uuid NOT NULL
);

ALTER TABLE "movie_genres" ADD FOREIGN KEY ("movie") REFERENCES "movies" ("id");
ALTER TABLE "movie_genres" ADD FOREIGN KEY ("genre") REFERENCES "genres" ("id");

-- migrate:down
DROP INDEX "idx_genre_name";
DROP INDEX "idx_users_email";
DROP TABLE "movie_genres";
DROP TABLE "genres";
DROP TABLE "user_movies";
DROP TABLE "movies";
DROP TABLE "users";
