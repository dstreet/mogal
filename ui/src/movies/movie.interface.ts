import { Genre } from "../gql/graphql"

export interface Movie {
  id: string
  title: string
  rating: string
  cast: string[]
  director: string
  poster?: string
  genres: Genre[]
  userRating?: number
}