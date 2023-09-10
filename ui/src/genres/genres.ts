import { useQuery } from "@tanstack/react-query";
import { graphql } from "../gql";
import { client } from "../api/client";
import { useAuth } from "../auth/auth.context";
import { Genre } from "./genres.interface";

const ListGenresQuery = graphql(`
  query ListGenres {
    listGenres {
      id
      name
    }
  }
`)

export const useGenres = () => {
  const { token } = useAuth()
  
  const { isLoading, error, data } = useQuery(['ListGenres'], () => 
    client.request(ListGenresQuery, {}, {
      Authorization: `Bearer ${token}`
    })
  )

  const genres: Genre[] = data?.listGenres
    .map(({id,name}) => ({ id, name })) ?? []

  return {
    loading: isLoading,
    genres,
    error
  }
}