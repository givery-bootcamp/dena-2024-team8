import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";
import { Spinner } from "../../components/Spinner";

export const UpdatePost = () => {
  const { postId = "" } = useParams();
  const { postDetail } = useAppSelector((state) => state.detail);
  const dispatch = useAppDispatch();
  const [title, setTitle] = useState("");
  const [body, setBody] = useState("");
  const navigator = useNavigate();
  const [errorMessage, setErrorMessage] = useState("");
  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    console.log("submit");
    dispatch(APIService.updatePost({ postId, title, body }))
      .then((response) => {
        console.log(response);
        navigator("/posts/" + postId);
      })
      .catch((error) => {
        setErrorMessage(`Error updating post: ${error}`);
      });
  };

  useEffect(() => {
    dispatch(APIService.getPostDetail(postId));
  }, [dispatch, postId]);

  useEffect(() => {
    if (postDetail) {
      setTitle(postDetail.title);
      setBody(postDetail.body);
    }
  }, [postDetail]);

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
        <form className="p-8" onSubmit={handleSubmit}>
          <p>Postの更新</p>
          <input
            className="form-control border rounded py-2 px-3 w-full text-indigo-500 leading-tight focus:outline-none focus:shadow-outline"
            id="titleInput"
            value={title}
            onChange={(e) => setTitle(e.target.value)} // この行を追加
          ></input>
          <textarea
            className="form-control border rounded py-2 px-3 w-full text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            id="bodyTextarea"
            value={body}
            placeholder="今の気持ちは？？"
            onChange={(e) => setBody(e.target.value)} // この行を追加
          ></textarea>
          <button
            type="submit"
            className="mt-3 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
          >
            Update
          </button>
          {errorMessage && <p className="text-red-500">{errorMessage}</p>}
        </form>
      </div>
    </div>
  );
};
