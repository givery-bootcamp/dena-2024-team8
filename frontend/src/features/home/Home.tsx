import PostForm from "../../components/form/PostForm";
import { PostList } from "../posts/PostList";

export function Home() {
  return (
    <div className="h-full">
      <div className="row px-8">
        <div>
          <PostForm />
          <hr />
          <PostList />
        </div>
      </div>
    </div>
  );
}
