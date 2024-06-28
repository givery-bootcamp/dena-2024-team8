import { User } from "./User";

export type PostList = Post[];

export type Post = {
  id: number;
  user_id: number;
  user: User;
  title: string;
  body: string;
  created_at: string;
  update_at: string;
  deleted_at: string;
};
