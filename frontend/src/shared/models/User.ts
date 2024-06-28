// {"id":1,"name":"taro","password":"password","created_at":"2024-05-27T10:59:25+09:00","updated_at":"2024-05-27T10:59:25+09:00","deleted_at":"0001-01-01T00:00:00Z"}%

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
