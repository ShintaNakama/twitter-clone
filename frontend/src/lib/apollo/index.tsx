import { ApolloClient, InMemoryCache } from '@apollo/client';

const cache = new InMemoryCache();
export const newApolloClient = new ApolloClient({
  uri: `${process.env.BFF_ENDPOINT}/graphql`,
  cache,
});

