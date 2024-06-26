import { createAsyncThunk } from "@reduxjs/toolkit";

import { CommentList, Hello } from "../models";
import { PostList, Post, User, SignOutResponse } from "../models";

const API_ENDPOINT_PATH = import.meta.env.VITE_API_ENDPOINT_PATH ?? "";

export const getHello = createAsyncThunk<Hello>("getHello", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/hello`);
  return await response.json();
});

export const getPostList = createAsyncThunk<PostList>(
  "getPostList",
  async () => {
    const response = await fetch(`${API_ENDPOINT_PATH}/posts`);
    return await response.json();
  },
);

export const getPostDetail = createAsyncThunk<Post, string>(
  "getPostDetail",
  async (postId: string) => {
    const response = await fetch(`${API_ENDPOINT_PATH}/posts/${postId}`);
    return await response.json();
  },
);

export const deletePost = createAsyncThunk<boolean, number>(
  "deletePost",
  async (postId: number) => {
    const response = await fetch(`${API_ENDPOINT_PATH}/posts/${postId}`, {
      method: "DELETE",
      credentials: "include",
    });
    return response.ok;
  },
);

export const signin = createAsyncThunk<
  User,
  { username: string; password: string }
>("signin", async ({ username, password }) => {
  const response = await fetch(`${API_ENDPOINT_PATH}/signin`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: "username=" + username + "&password=" + password,
    credentials: "include",
  });

  if (!response.ok) {
    throw new Error("サインインに失敗しました。");
  }

  const data = await response.json();

  return data;
});

export const signout = createAsyncThunk<SignOutResponse>(
  "signout",
  async () => {
    const response = await fetch(`${API_ENDPOINT_PATH}/signout`, {
      method: "POST",
      credentials: "include",
    });
    if (!response.ok) {
      throw new Error("サインアウトに失敗しました。");
    }
    const data = await response.json();

    return data;
  },
);

export const createPost = createAsyncThunk<
  Post,
  { title: string; content: string }
>("createPost", async ({ title, content: content }) => {
  const response = await fetch(`${API_ENDPOINT_PATH}/posts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: new URLSearchParams({
      title: title,
      body: content,
    }),
    credentials: "include",
  });
  return await response.json();
});

export const getSearchPostList = createAsyncThunk<PostList, string>(
  "getSearchPost",
  async (query) => {
    const response = await fetch(`${API_ENDPOINT_PATH}/posts?q=${query}`);
    return await response.json();
  },
);

export const getUser = createAsyncThunk<User>("getUser", async () => {
  const response = await fetch(`${API_ENDPOINT_PATH}/auth/user`, {
    method: "GET",
    credentials: "include",
  });
  if (response.status != 200) {
    throw new Error("ユーザー情報の取得に失敗しました。");
  }
  return await response.json();
});

export const updatePost = createAsyncThunk<
  Post,
  { postId: string; title: string; body: string }
>("updatePost", async ({ postId, title, body }) => {
  const response = await fetch(`${API_ENDPOINT_PATH}/posts/${postId}`, {
    method: "PUT",
    body: new URLSearchParams({
      title: title,
      body: body,
    }),
    credentials: "include",
  });
  return await response.json();
});

export const getCommentList = createAsyncThunk<CommentList, number>(
  "getCommentList",
  async (postId) => {
    const response = await fetch(
      `${API_ENDPOINT_PATH}/posts/${postId}/comments`,
    );
    return await response.json();
  },
);
