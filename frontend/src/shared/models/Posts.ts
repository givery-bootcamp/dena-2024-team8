export type PostList = Array<Post>;
  
export type Post = {
    id: number;
    UserId: number;
    Title: string;
    Body: string;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
};