import { useEffect } from "react";
import PostForm from "../../components/form/PostForm";
import { useAppSelector, useAppDispatch } from "../../shared/hooks";
import { APIService } from "../../shared/services";
import { PostList } from "../posts/PostList";

export function Home() {
  const { postList } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostList());
  }, [dispatch]);

  return (
    <div className="h-full">
      <div className="px-8">
        <div>
          <PostForm />
          <hr />
          <PostList posts={postList ?? []} />
        </div>
      </div>
    </div>
  );
}
