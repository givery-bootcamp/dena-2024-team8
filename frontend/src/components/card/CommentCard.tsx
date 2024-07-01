import { Comment } from "../../shared/models";

type CommentCardProps = {
  comment: Comment;
};

export const CommentCard = ({ comment }: CommentCardProps) => {
  return (
    <div className="border p-4 mb-4 rounded shadow-sm">
      <p className="mb-2">{comment.body}</p>
      <time className="text-gray-500">{new Date(comment.created_at).toLocaleString()}</time>
    </div>
  );
};
