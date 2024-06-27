import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";
import { Spinner } from "../../components/Spinner";

export const Post = () => {
  const { postId } = useParams();
  const { postDetail } = useAppSelector((state) => state.detail);
  const dispatch = useAppDispatch();

  // ここでpostIdを使ってデータベースから対象のpostの詳細を取得
  useEffect(() => {
    if (postId != undefined) {
      dispatch(APIService.getPostDetail(postId));
    }
  }, [dispatch, postId]);

  if (postDetail == undefined) {
    return (
      <div className="max-w-xl mx-auto bg-white overflow-hidden md:max-w-2xl my-4">
        <Spinner />
      </div>
    );
  }

  return (
    <div className="max-w-xl mx-auto bg-white shadow-md overflow-hidden md:max-w-2xl my-4">
      <div className="md:flex">
        <div className="p-8">
          <h2 className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">
            {postDetail.title}
          </h2>
          <p className="mt-2 text-gray-500">{postDetail.body}</p>
          <div className="mt-4">
            <span className="text-gray-600 text-sm">
              Posted by User {postDetail.user_id}
            </span>
          </div>
          <div className="mt-4 text-gray-500 text-xs">
            <p>
              Created At: {new Date(postDetail.created_at).toLocaleString()}
            </p>
            <p>Updated At: {new Date(postDetail.update_at).toLocaleString()}</p>
            {postDetail.deleted_at ? (
              <p className="text-red-500">
                Deleted At: {new Date(postDetail.deleted_at).toLocaleString()}
              </p>
            ) : null}
          </div>
        </div>
      </div>
    </div>
  );
};
