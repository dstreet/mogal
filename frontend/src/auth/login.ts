import { useMutation } from '@tanstack/react-query'
import { graphql } from "../gql";
import { LoginMutationVariables } from '../gql/graphql';
import { client } from '../api/client';

const LoginMutation = graphql(`
  mutation Login($email: String!, $password: String!) {
    login(input: { email: $email, password: $password}) {
      token
      expires_in
    }
  }
`)

export const useLogin = () => {
  const { isLoading, error, data, mutate } = useMutation(['Login'], (input: LoginMutationVariables) =>
    client.request(LoginMutation, input)
  )

  const login = (email: string, password: string) => {
    mutate({
      email,
      password
    })
  }

  return {
    loading: isLoading,
    token: data?.login.token,
    expiresIn: data?.login.expires_in,
    error,
    login
  }
}