import { createSlice } from "@reduxjs/toolkit";

import { Post } from "../models";
import { APIService } from "../services";

export type UpdatePostState = {
  updatePost?: Post;
};

export const initialState: UpdatePostState = {};

export const updatePostSlice = createSlice({
  name: "updatePost",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.updatePost.fulfilled, (state, action) => {
      state.updatePost = action.payload;
    });
  },
});

export default updatePostSlice.reducer;
