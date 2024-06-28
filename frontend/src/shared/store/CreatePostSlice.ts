import { createSlice } from "@reduxjs/toolkit";

import { Post } from "../models";
import { APIService } from "../services";

export type CreatePostState = {
  createPost?: Post;
};

export const initialState: CreatePostState = {};

export const createPostSlice = createSlice({
  name: "post",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getPostDetail.fulfilled, (state, action) => {
      state.createPost = action.payload;
    });
  },
});

export default createPostSlice.reducer;
