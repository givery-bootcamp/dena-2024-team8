import { useEffect, useState } from "react";
import { APIService } from "../../shared/services";
import { useAppDispatch } from "../../shared/hooks";

export default function PostForm() {
  const [postTitle, setPostTitle] = useState(""); // 投稿内容
  const [postContent, setPostContent] = useState(""); // 投稿内容
  const dispatch = useAppDispatch();

  const resetForm = () => {
    setPostContent("");
    setPostTitle("");
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    dispatch(
      APIService.createPost({ title: postTitle, content: postContent })
    ).then(() => {
      dispatch(APIService.getPostList());
    });
    resetForm();
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
        タイトル
      </label>
      <input
        type="text"
        value={postTitle}
        onChange={(e) => setPostTitle(e.target.value)}
        placeholder="タイトルを入力"
        className="w-3/4 border border-gray-300 rounded-md px-4 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <label
        htmlFor="TweetFormTextarea"
        className="block text-gray-700 text-sm font-bold mb-2 mt-2"
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
