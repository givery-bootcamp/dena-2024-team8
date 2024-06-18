export type PostList = Array<Post>;
  
export type Post = {
    Id: number;
    UserId: number;
    Title: string;
    Body: string;
    CreatedAt: string;
    UpdatedAt: string;
    DeletedAt: string;
};