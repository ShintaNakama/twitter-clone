import {
  PostDetailFragment,
 } from "../../../graphql";

type Props = {
  data: PostDetailFragment;
};

const Component = ({ data }: Props) => {
  const {
    id,
    user,
    body,
  } = data;

  return (
    <>
    <div className="id">{id}</div>
    <div className="userName">{user.name}</div>
    <div className="body">{body}</div>
    </>
  )
};

export const PostDetail = Component
