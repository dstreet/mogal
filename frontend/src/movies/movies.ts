import { useQuery } from "@tanstack/react-query";
import { graphql } from "../gql";
import { client } from "../api/client";
import { useAuth } from "../auth/auth.context";
import { Movie } from "./movie.interface";

const ListMoviesQuery = graphql(`
  query ListMovies($genre: String) {
    listMovies(input: { genre: $genre }) {
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

const GetMovieQuery = graphql(`
  query GetMovie($id: ID!) {
    getMovie(input: { id: $id }) {
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

export const useMovies = (genre?: string) => {
  const { token } = useAuth()

  const { isLoading, error, data } = useQuery(['ListMovies'], () => 
    client.request(ListMoviesQuery, { genre }, {
      Authorization: `Bearer ${token}`
    })
  )

  let movies: Movie[] = []

  if (data?.listMovies) {
    movies = data.listMovies.map(m => ({
      ...m,
      poster: m.poster ?? undefined,
      userRating: m.userRating ?? undefined,
    }))
  }

  return {
    loading: isLoading,
    error,
    movies
  }
}

export const useMovie = (id: string) => {
  const { token } = useAuth()

  const { isLoading, error, data } = useQuery(['GetMovie'], () => 
    client.request(GetMovieQuery, { id }, {
      Authorization: `Bearer ${token}`
    })
  )

  let movie: Movie | undefined

  if (data?.getMovie) {
    movie = {
      ...data.getMovie,
      poster: data.getMovie.poster ?? undefined,
      userRating: data.getMovie.userRating ?? undefined
    }
  }

  return {
    loading: isLoading,
    error,
    movie
  }
}