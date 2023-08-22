/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  mutation Login($email: String!, $password: String!) {\n    login(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n": types.LoginDocument,
    "\n  mutation RegisterUser($email: String!, $password: String!) {\n    register(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n": types.RegisterUserDocument,
    "\n  query ListGenres {\n    listGenres {\n      id\n      name\n    }\n  }\n": types.ListGenresDocument,
    "\n  mutation CreateMovie(\n    $title: String!\n    $rating: String!\n    $cast: [String!]!\n    $director: String!\n    $poster: String\n    $userRating: Int\n    $genres: [String!]!\n  ) {\n    createMovie(input: {\n      title: $title,\n      rating: $rating,\n      cast: $cast,\n      director: $director,\n      poster: $poster,\n      userRating: $userRating,\n      genres: $genres\n    }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n      genres {\n        id\n        name\n      }\n    }\n  }\n": types.CreateMovieDocument,
    "\n  mutation CreateGenres($input: [CreateGenreInput!]!) {\n    createGenres(input: $input) {\n      id\n      name\n    }\n  }\n": types.CreateGenresDocument,
    "\n  query ListMovies($genre: String) {\n    listMovies(input: { genre: $genre }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n": types.ListMoviesDocument,
    "\n  query GetMovie($id: ID!) {\n    getMovie(input: { id: $id }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n": types.GetMovieDocument,
};

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = graphql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function graphql(source: string): unknown;

/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation Login($email: String!, $password: String!) {\n    login(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n"): (typeof documents)["\n  mutation Login($email: String!, $password: String!) {\n    login(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation RegisterUser($email: String!, $password: String!) {\n    register(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n"): (typeof documents)["\n  mutation RegisterUser($email: String!, $password: String!) {\n    register(input: { email: $email, password: $password}) {\n      token\n      expires_in\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ListGenres {\n    listGenres {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  query ListGenres {\n    listGenres {\n      id\n      name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateMovie(\n    $title: String!\n    $rating: String!\n    $cast: [String!]!\n    $director: String!\n    $poster: String\n    $userRating: Int\n    $genres: [String!]!\n  ) {\n    createMovie(input: {\n      title: $title,\n      rating: $rating,\n      cast: $cast,\n      director: $director,\n      poster: $poster,\n      userRating: $userRating,\n      genres: $genres\n    }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n      genres {\n        id\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation CreateMovie(\n    $title: String!\n    $rating: String!\n    $cast: [String!]!\n    $director: String!\n    $poster: String\n    $userRating: Int\n    $genres: [String!]!\n  ) {\n    createMovie(input: {\n      title: $title,\n      rating: $rating,\n      cast: $cast,\n      director: $director,\n      poster: $poster,\n      userRating: $userRating,\n      genres: $genres\n    }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n      genres {\n        id\n        name\n      }\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  mutation CreateGenres($input: [CreateGenreInput!]!) {\n    createGenres(input: $input) {\n      id\n      name\n    }\n  }\n"): (typeof documents)["\n  mutation CreateGenres($input: [CreateGenreInput!]!) {\n    createGenres(input: $input) {\n      id\n      name\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query ListMovies($genre: String) {\n    listMovies(input: { genre: $genre }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n"): (typeof documents)["\n  query ListMovies($genre: String) {\n    listMovies(input: { genre: $genre }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n"];
/**
 * The graphql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function graphql(source: "\n  query GetMovie($id: ID!) {\n    getMovie(input: { id: $id }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n"): (typeof documents)["\n  query GetMovie($id: ID!) {\n    getMovie(input: { id: $id }) {\n      id\n      title\n      rating\n      cast\n      director\n      poster\n      userRating\n    }\n  }\n"];

export function graphql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;