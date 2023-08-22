import { useMutation } from "@tanstack/react-query";
import { graphql } from "../gql";
import { CreateMovieMutationVariables } from "../gql/graphql";
import { client } from "../api/client";
import { Movie } from "./movie.interface";
import { useAuth } from "../auth/auth.context";

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

export const useCreateMovie = () => {
  const { token } = useAuth()

  const { isLoading, isSuccess, error, data, mutate } = useMutation(['CreateMovie'], (input: CreateMovieMutationVariables) =>
    client.request(CreateMovieMutation, input, {
      Authorization: `Bearer ${token}`
    })
  )

  let movie: Movie | undefined = undefined

  if (data?.createMovie) {
    movie = {
      ...data.createMovie,
      poster: data.createMovie.poster ?? undefined,
      userRating: data.createMovie.userRating ?? undefined
    }
  }

  return {
    loading: isLoading,
    success: isSuccess,
    error,
    movie,
    createMovie: mutate
  }
}