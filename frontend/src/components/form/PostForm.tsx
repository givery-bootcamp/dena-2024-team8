export default function PostForm() {
    return (
        <div className="mb-3">
          <label htmlFor="TweetFormTextarea" >今の気持ちは？？</label>
          <textarea className="form-control" id="TweetFormTextarea" placeholder="今の気持ちは？？">

          </textarea>
          <button type="button" className="btn btn-primary">Button</button>
        </div>
    );
}