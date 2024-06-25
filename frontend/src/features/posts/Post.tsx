import { useEffect } from "react";
import { useParams } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";

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
    return <>ロード中</>;
  }

  return (
    <div className="max-w-xl mx-auto bg-white shadow-md overflow-hidden md:max-w-2xl my-4">
      <div className="md:flex">
        <div className="p-8">
          <div className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">
            {postDetail.Title}
          </div>
          <p className="mt-2 text-gray-500">{postDetail.Body}</p>
          <div className="mt-4">
            <span className="text-gray-600 text-sm">
              Posted by User {postDetail.UserId}
            </span>
          </div>
          <div className="mt-4 text-gray-500 text-xs">
            <p>Created At: {new Date(postDetail.CreatedAt).toLocaleString()}</p>
            <p>Updated At: {new Date(postDetail.UpdatedAt).toLocaleString()}</p>
            {postDetail.DeletedAt ? (
              <p className="text-red-500">
                Deleted At: {new Date(postDetail.DeletedAt).toLocaleString()}
              </p>
            ) : null}
          </div>
        </div>
      </div>
    </div>
  );
};
