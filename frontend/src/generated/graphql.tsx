import { gql } from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
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
