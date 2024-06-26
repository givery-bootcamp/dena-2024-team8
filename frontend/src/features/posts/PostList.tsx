import { Post } from "../../shared/models";
import { PostCard } from "../../components/card/PostCard";

type PostListProps = {
  posts: Post[];
};

export const PostList = ({ posts }: PostListProps) => {
  return (
    <div className="w-full mt-4 py-4">
      {posts &&
        posts.map((post: Post, index: number) => (
          <PostCard
            key={index}
            title={post.Title}
            content={post.Body}
            date={post.CreatedAt}
            postId={post.id}
          />
        ))}
    </div>
  );
};
