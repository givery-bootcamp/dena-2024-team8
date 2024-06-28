import { createSlice } from "@reduxjs/toolkit";

import { Post } from "../models";
import { APIService } from "../services";

export type PostState = {
  posts?: Post[];
};

export const initialState: PostState = {};

export const searchPostSlice = createSlice({
  name: "searchPostList",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getSearchPostList.fulfilled, (state, action) => {
      state.posts = action.payload;
    });
  },
});

export default searchPostSlice.reducer;
