-- migrate:up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "movies" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "title" text NOT NULL,
  "rating" text NOT NULL,
  "genre" text NOT NULL,
  "cast" text[] NOT NULL,
  "director" text NOT NULL,
  "poster" text,
  "user" uuid NOT NULL,
  "user_rating" int
);

CREATE TABLE "users" (
  "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
  "email" text NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamp NOT NULL,
  "deleted_at" timestamp,
  "last_login" timestamp NOT NULL
);

CREATE INDEX "idx_movies_genre" ON "movies" ("genre");
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
ALTER TABLE "movies" ADD FOREIGN KEY ("user") REFERENCES "users" ("id");

-- migrate:down
DROP INDEX "idx_movies_genre";
DROP INDEX "idx_users_email";
DROP TABLE "movies"; 
DROP TABLE "users";
