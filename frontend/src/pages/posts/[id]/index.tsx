import { filter } from "graphql-anywhere";
import { GetStaticPaths, GetStaticProps, NextPage } from "next";
import { PostDetail } from "../../../components/layouts";
import { newApolloClient } from "../../../lib/apollo";
import { DetailPageGetDataQuery,
  DetailPageGetDataDocument,
  DetailPageGetDataQueryVariables,
  PostDetailFragment,
  PostDetailFragmentDoc
 } from "../../../graphql";

type Props = {
  data?: DetailPageGetDataQuery;
};

const PostDetailPage: NextPage<Props> = ({ data }) => {
  const post = data?.post;
  return (
    <PostDetail
      data={filter<PostDetailFragment>(
        PostDetailFragmentDoc,
        post
      )}
    />
  )
};

export const getStaticPaths: GetStaticPaths = async () => {
  // Generating static pagesしないので空配列を返す
  return { paths: [], fallback: "blocking" };
};

export const getStaticProps: GetStaticProps = async ({ params }) => {
  const queryId = (params?.id || "") as string;
  console.log(queryId)
  try {
    const { data } = await newApolloClient.query<
      DetailPageGetDataQuery,
      DetailPageGetDataQueryVariables
    >({
      query: DetailPageGetDataDocument,
      variables: { id: queryId },
      fetchPolicy: "no-cache",
    });

    return {
      props: { data },
    };
  } catch (error) {
    console.error({ error });
    return { props: {}};
  }
};

export default PostDetailPage;
