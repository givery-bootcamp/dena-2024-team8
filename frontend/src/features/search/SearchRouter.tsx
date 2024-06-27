import { useEffect, useState } from "react";
import { PostList } from "../posts";
import { useAppSelector, useAppDispatch } from "../../shared/hooks";
import { APIService } from "../../shared/services";

export const SearchRouter = () => {
  const { posts } = useAppSelector((state) => state.searchPostList);
  const dispatch = useAppDispatch();

  const [query, setQuery] = useState("");

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    dispatch(APIService.getSearchPostList(query));
  };

  useEffect(() => {
    dispatch(APIService.getSearchPostList(query));
  }, [dispatch]);

  return (
    <div className="pl-8 space-y-4">
      <form
        onSubmit={handleSubmit}
        className="flex flex-row items-center gap-4"
      >
        <input
          type="text"
          value={query}
          onChange={(e) => setQuery(e.target.value)}
          placeholder="キーワードを入力"
          className="flex-1 border border-gray-300 rounded-md px-4 py-2 mt-3 focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <button
          type="submit"
          className="mt-3 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          検索
        </button>
      </form>
      <hr />
      {posts && <PostList posts={posts!} />}
    </div>
  );
};

export default SearchRouter;
