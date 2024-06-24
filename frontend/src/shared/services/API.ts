import { createAsyncThunk } from '@reduxjs/toolkit';

import { Hello } from '../models';
import { PostList,Post } from '../models';

const API_ENDPOINT_PATH =
  import.meta.env.VITE_API_ENDPOINT_PATH ?? '';

export const getHello = createAsyncThunk<Hello>('getHello', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPostList = createAsyncThunk<PostList>('getPostList', async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/posts`);
  return await response.json();
});

export const getPostDetail = createAsyncThunk<Post,string>('getPostDetail',async( postId : string ) => {
  const response = await fetch(`${API_ENDPOINT_PATH}/posts/${postId}`);
  return await response.json();
}) 

export const signin = createAsyncThunk<Post, { username: string; password: string }>(
  'signin',
  async ({ username, password }) => {
    const response = await fetch(`${API_ENDPOINT_PATH}/signin`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
      throw new Error('サインインに失敗しました。');
    }

    const data = await response.json();

    // JWTをcookieに設定する処理はフロントエンドでは実施しないのが一般的ですが、
    // 必要に応じてサーバーからのレスポンスヘッダーを確認し、適切に処理を行ってください。

    return data;
  }
);