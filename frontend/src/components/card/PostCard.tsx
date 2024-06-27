import { Link } from "react-router-dom";
import { Post } from "../../shared/models";

type PostCardProps = {
  post: Post;
};

export const PostCard = ({ post }: PostCardProps) => {
  return (
    <Link to={`/posts/${post.id}`}>
      <div className="border p-4 mb-4 rounded shadow-sm">
        <h3 className="font-bold text-lg mb-2">{post.title}</h3>
        <p className="mb-2">{post.body}</p>
        <time className="text-gray-500">{post.created_at}</time>
      </div>
    </Link>
  );
};
