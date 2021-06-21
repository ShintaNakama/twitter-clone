import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions =  {}
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

/** Mutation */
export type Mutation = {
  __typename?: 'Mutation';
  /** ユーザー作成 */
  createUser: Scalars['Boolean'];
  /** 投稿作成 */
  createPost: Scalars['Boolean'];
  /** 投稿削除 */
  deletePost: Scalars['Boolean'];
};


/** Mutation */
export type MutationCreateUserArgs = {
  input?: Maybe<UserInput>;
};


/** Mutation */
export type MutationCreatePostArgs = {
  input?: Maybe<PostInput>;
};


/** Mutation */
export type MutationDeletePostArgs = {
  id: Scalars['ID'];
};

/** post */
export type Post = {
  __typename?: 'Post';
  /** ID */
  id: Scalars['ID'];
  /** ユーザー */
  user: User;
  /** 本文 */
  body: Scalars['String'];
};

/** post input */
export type PostInput = {
  /** ID */
  id: Scalars['ID'];
  /** ユーザー */
  userID: Scalars['ID'];
  /** 本文 */
  body: Scalars['String'];
};

/** Query */
export type Query = {
  __typename?: 'Query';
  /** 投稿リスト */
  postList: Array<Post>;
  /** 投稿 */
  post: Post;
};


/** Query */
export type QueryPostArgs = {
  id: Scalars['ID'];
};

/** user */
export type User = {
  __typename?: 'User';
  /** ID */
  id: Scalars['ID'];
  /** email */
  email: Scalars['String'];
  /** name */
  name: Scalars['String'];
  /** image */
  image: Scalars['String'];
};

/** user input */
export type UserInput = {
  /** ID */
  id: Scalars['ID'];
  /** email */
  email: Scalars['String'];
  /** name */
  name: Scalars['String'];
  /** image */
  image: Scalars['String'];
};

export type PostDetailFragment = (
  { __typename?: 'Post' }
  & Pick<Post, 'id' | 'body'>
  & { user: (
    { __typename?: 'User' }
    & Pick<User, 'id' | 'name' | 'image'>
  ) }
);

export type DetailPageGetDataQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type DetailPageGetDataQuery = (
  { __typename?: 'Query' }
  & { post: (
    { __typename?: 'Post' }
    & PostDetailFragment
  ) }
);

export const PostDetailFragmentDoc = gql`
    fragment PostDetail on Post {
  id
  user {
    id
    name
    image
  }
  body
}
    `;
export const DetailPageGetDataDocument = gql`
    query DetailPageGetData($id: ID!) {
  post(id: $id) {
    ...PostDetail
  }
}
    ${PostDetailFragmentDoc}`;

/**
 * __useDetailPageGetDataQuery__
 *
 * To run a query within a React component, call `useDetailPageGetDataQuery` and pass it any options that fit your needs.
 * When your component renders, `useDetailPageGetDataQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useDetailPageGetDataQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useDetailPageGetDataQuery(baseOptions: Apollo.QueryHookOptions<DetailPageGetDataQuery, DetailPageGetDataQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<DetailPageGetDataQuery, DetailPageGetDataQueryVariables>(DetailPageGetDataDocument, options);
      }
export function useDetailPageGetDataLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<DetailPageGetDataQuery, DetailPageGetDataQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<DetailPageGetDataQuery, DetailPageGetDataQueryVariables>(DetailPageGetDataDocument, options);
        }
export type DetailPageGetDataQueryHookResult = ReturnType<typeof useDetailPageGetDataQuery>;
export type DetailPageGetDataLazyQueryHookResult = ReturnType<typeof useDetailPageGetDataLazyQuery>;
export type DetailPageGetDataQueryResult = Apollo.QueryResult<DetailPageGetDataQuery, DetailPageGetDataQueryVariables>;