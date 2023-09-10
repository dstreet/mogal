import { useMutation } from "@tanstack/react-query";
import { graphql } from "../gql";
import { client } from "../api/client";
import { Movie } from "./movie.interface";
import { useAuth } from "../auth/auth.context";
import { CreateGenresMutationVariables, CreateMovieMutationVariables } from "../gql/graphql";
import { useState } from "react";

const CreateMovieMutation = graphql(`
  mutation CreateMovie(
    $title: String!
    $rating: String!
    $cast: [String!]!
    $director: String!
    $poster: String
    $userRating: Int
    $genres: [String!]!
  ) {
    createMovie(input: {
      title: $title,
      rating: $rating,
      cast: $cast,
      director: $director,
      poster: $poster,
      userRating: $userRating,
      genres: $genres
    }) {
      id
      title
      rating
      cast
      director
      poster
      userRating
      genres {
        id
        name
      }
    }
  }
`)

const CreateGenresMutation = graphql(`
  mutation CreateGenres($input: [CreateGenreInput!]!) {
    createGenres(input: $input) {
      id
      name
    }
  }
`)

export const useCreateMovie = () => {
  const { token } = useAuth()
  const headers = {
    Authorization: `Bearer ${token}`
  }

  const [error, setError] = useState<unknown>(null)

  const createMovieMutation = useMutation(['CreateMovie'], (input: CreateMovieMutationVariables) =>
    client.request(CreateMovieMutation, input, headers)
  )

  const createGenreMutation = useMutation(['CreateGenre'], (input: CreateGenresMutationVariables) =>
    client.request(CreateGenresMutation, input, headers)
  )

  const createMovie = async (movie: CreateMovieMutationVariables, newGenres: string[]) => {
    try {
      if (newGenres.length) {
        const { createGenres } = await createGenreMutation.mutateAsync({ input: newGenres.map(g => ({ name: g })) })
        const createdGeneres = createGenres.map(g => g.id)

        const genres = Array.isArray(movie.genres)
          ? [...movie.genres, ...createdGeneres]
          : [movie.genres, ...createdGeneres]

        movie.genres = genres
      }
      createMovieMutation.mutate(movie)
    } catch (err) {
      setError(err)
    }
  }

  let movie: Movie | undefined = undefined

  if (createMovieMutation.data?.createMovie) {
    movie = {
      ...createMovieMutation.data.createMovie,
      poster: createMovieMutation.data.createMovie.poster ?? undefined,
      userRating: createMovieMutation.data.createMovie.userRating ?? undefined
    }
  }

  return {
    loading: createMovieMutation.isLoading,
    success: createMovieMutation.isSuccess,
    error: error || createMovieMutation.error,
    movie,
    createMovie
  }
}