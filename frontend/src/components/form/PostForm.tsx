export default function PostForm() {
    return (
        <div className="mb-3" style={{ margin: '1%' }}>
          <label htmlFor="exampleFormControlTextarea1" className="form-label">今の気持ちは？？</label>
          <textarea className="form-control" id="exampleFormControlTextarea1">

          </textarea>
          <button type="button" className="btn btn-primary">Button</button>
        </div>
    );
}