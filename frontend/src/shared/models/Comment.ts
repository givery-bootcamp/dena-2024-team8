export type CommentList = Comment[];

export type Comment = {
  id: number;
  user_id: number;
  post_id: number;
  body: string;
  created_at: string;
  update_at: string;
  deleted_at: string;
};
