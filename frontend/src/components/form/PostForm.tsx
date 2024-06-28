import { useEffect, useState } from "react";
import { APIService } from "../../shared/services";
import { useAppDispatch } from "../../shared/hooks";

export default function PostForm() {
  const [postContent, setPostContent] = useState(""); // 投稿内容
  const dispatch = useAppDispatch();
  const title = "hoge";

  const resetText = () => {
    setPostContent("");
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    dispatch(APIService.createPost({ title, content: postContent })).then(
      () => {
        dispatch(APIService.getPostList());
      },
    );
    resetText();
  };

  useEffect(() => {
    console.log(postContent);
  }, [postContent]);

  return (
    <form className="mb-3" onSubmit={handleSubmit}>
      <label
        htmlFor="TweetFormTextarea"
        className="block text-gray-700 text-sm font-bold mb-2"
      >
        今の気持ちは？？
      </label>
      <textarea
        className="form-control border rounded py-2 px-3 w-full text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        id="TweetFormTextarea"
        value={postContent}
        placeholder="今の気持ちは？？"
        onChange={(e) => setPostContent(e.target.value)}
      ></textarea>
      <button
        type="submit"
        className="mt-3 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Tweet
      </button>
    </form>
  );
}
