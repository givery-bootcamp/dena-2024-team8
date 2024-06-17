export type Post = {
    Id: number;
    UserId: number;
    Title: string;
    Body: string;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
};

export default function PostCard({item}: {item: Post}) {
    const bodyHTML = item.Body.replace(/\n/g, "<br>");
    return (
      <div className="card">
        <div className="card-body">
          <p className="card-title fs-5">{item.Title}</p>
          <div className="card-text" dangerouslySetInnerHTML={{ __html: bodyHTML }} ></div>
          <p className="card-text">{new Date(item.CreatedAt).toLocaleString("ja-JP")}</p>
        </div>
      </div>
    );
}