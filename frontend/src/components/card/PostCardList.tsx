import { PostList, Post } from "../../shared/models";
import PostCard from "./PostCard";

export default function PostCardList({itemList = []}: {itemList: PostList}) { // TODO 型の定義 Add type for itemList
    return (
        <div className="card-list">
            {itemList.map((item: Post, index: number) => (
                <PostCard key={index} item={item}/> // TODO propsの設定 item={item} /> // Render PostCard component with item data
            ))}
        </div>
    );
}