import request, { GraphQLClient } from 'graphql-request'
import { useMutation } from '@tanstack/react-query'
import { graphql } from "../gql";
import { RegisterUserMutationVariables } from '../gql/graphql';
import { client } from '../api/client';

const RegisterUserMutation = graphql(`
  mutation RegisterUser($email: String!, $password: String!) {
    register(input: { email: $email, password: $password}) {
      token
      expires_in
    }
  }
`)

export const useRegisterUser = () => {
  const { isLoading, error, data, mutate } = useMutation(['RegisterUser'], (input: RegisterUserMutationVariables) =>
    client.request(RegisterUserMutation, input)
  )

  const register = (email: string, password: string) => {
    mutate({email, password})
  }

  return {
    loading: isLoading,
    token: data?.register.token,
    expiresIn: data?.register.expires_in,
    error,
    register
  }
}