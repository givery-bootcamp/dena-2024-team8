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
        <h5 className="card-title">{item.Title}</h5>
        <div className="card-text" dangerouslySetInnerHTML={{ __html: bodyHTML }} ></div>
        <p className="card-text">{new Date(item.CreatedAt).toLocaleString("ja-JP")}</p>
      </div>
      </div>
    );
}