import PostCard from "./PostCard";

export default function PostCardList({itemList}: {itemList: any[]}) { // TODO 型の定義 Add type for itemList
        return (
            <div className="card-list">
                {itemList.map((item: any) => (
                    <PostCard key={item.id} /> // TODO propsの設定 item={item} /> // Render PostCard component with item data
                ))}
            </div>
        );
}