import { Link } from 'react-router-dom';

type PostCardProps = {
  title: string;
  content: string;
  date: string;
  postId : number;
}

export const PostCard = ({ title, content, date ,postId}: PostCardProps) => {
  return (
    <Link to={`/posts/${postId}`}>
      <div className="border p-4 mb-4 rounded shadow-sm">
        <h3 className="font-bold text-lg mb-2">{title}</h3>
        <p className="mb-2">{content}</p>
        <time className="text-gray-500">{date}</time>
      </div>
    </Link>
  );
};
