export interface User {
  id: number;
  name: string;
  password: string;
  created_at: string;
  updated_at: string;
  deleted_at: string;
}

export interface SignOutResponse {
  message: string;
}
