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
          <h4 className="card-title">{item.Title}</h4>
          <div>
            <p className="card-text" dangerouslySetInnerHTML={{ __html: bodyHTML }} ></p>
            <time className="card-text">{new Date(item.CreatedAt).toLocaleString("ja-JP")}</time>
          </div>
        </div>
      </div>
    );
}