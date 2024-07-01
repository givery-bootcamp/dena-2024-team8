import { useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";
import { Spinner } from "../../components/Spinner";
import { actions } from "../../shared/store";
import { CommentCard } from "../../components/card/CommentCard";

export const PostDetailRouter = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const { postId = "" } = useParams();

  const { postDetail } = useAppSelector((state) => state.detail);
  const { isDeleted, error_message } = useAppSelector(
    (state) => state.deletePost,
  );
  const { commentList } = useAppSelector((state) => state.commentList);
  useEffect(() => {
    dispatch(APIService.getCommentList(Number(postId)));
  }, [dispatch, postId]);

  const handleDelete = (id: number) => {
    dispatch(APIService.deletePost(id));
  };
  const navigator = useNavigate();

  // ここでpostIdを使ってデータベースから対象のpostの詳細を取得
  useEffect(() => {
    dispatch(APIService.getPostDetail(postId));
  }, [dispatch, postId]);

  useEffect(() => {
    if (isDeleted) {
      dispatch(actions.resetDeletePostState());
      navigate("/");
    }
  }, [isDeleted, navigate, dispatch]);

  if (postDetail == undefined) {
    return (
      <div className="max-w-xl mx-auto bg-white overflow-hidden md:max-w-2xl my-4">
        <Spinner />
      </div>
    );
  }
  console.log("postDetail:", postDetail);
  const handleRedirectEdit = () => {
    navigator("/posts/" + postId + "/edit");
  };

  return (
    <div className="max-w-xl mx-auto md:max-w-2xl">
      {error_message && <p className="text-red-500">{error_message}</p>}
      <div className="bg-white shadow-md overflow-hidden my-4">
        <div className="p-8">
          <div className="flex justify-between items-center">
            <h2 className="uppercase tracking-wide text-sm text-indigo-500 font-semibold">
              {postDetail.title}
            </h2>
            <button
              onClick={handleRedirectEdit}
              className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
            >
              編集をする
            </button>
          </div>
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
      <div className="flex justify-end p-4">
        <button
          className="px-4 py-2 bg-red-500 text-white rounded-md"
          onClick={() => handleDelete(postDetail.id)}
        >
          Delete
        </button>
      </div>
      {commentList?.map((comment) => (
        <CommentCard key={comment.id} comment={comment} />
      ))}
    </div>
  );
};
