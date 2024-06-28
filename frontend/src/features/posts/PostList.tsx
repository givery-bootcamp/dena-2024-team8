import { Post } from "../../shared/models";
import { PostCard } from "../../components/card/PostCard";

type PostListProps = {
  posts: Post[];
};

export const PostList = ({ posts }: PostListProps) => {
  return (
    <div className="w-3/4 p-4">
      {posts &&
        posts.map((post: Post, index: number) => (
          <PostCard key={index} post={post} />
        ))}
    </div>
  );
};
