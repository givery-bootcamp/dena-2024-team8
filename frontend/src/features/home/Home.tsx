import PostForm from "../../components/form/PostForm";
import { PostList } from "../posts/PostList";

export function Home() {
  return (
    <div className="container-xxl">
      <div className="row">
        <div className="border bg-white vh-100 p-3">
          <PostForm />
          <hr />
          <PostList />
        </div>
      </div>
    </div>
  );
}
