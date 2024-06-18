export default function PostForm() {
  return (
    <div className="mb-3">
      <label
        htmlFor="TweetFormTextarea"
        className="block text-gray-700 text-sm font-bold mb-2"
      >
        今の気持ちは？？
      </label>
      <textarea
        className="form-control border rounded py-2 px-3 w-full text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        id="TweetFormTextarea"
        placeholder="今の気持ちは？？"
      ></textarea>
      <button
        type="button"
        className="mt-3 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        Button
      </button>
    </div>
  );
}
