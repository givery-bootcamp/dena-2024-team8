import { useEffect } from "react";

import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";

import { Post } from "../../shared/models";
import { PostCard } from "../../components/card/PostCard";

export const PostList = () => {
  const { postList } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostList());
  }, [dispatch]);

  return (
    <div className="w-3/4 p-4">
      {postList &&
        postList.map((post: Post, index: number) => (
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
