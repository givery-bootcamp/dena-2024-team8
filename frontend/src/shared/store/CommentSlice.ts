import { createSlice } from "@reduxjs/toolkit";

import { CommentList } from "../models";
import { APIService } from "../services";

export type CommentListState = {
  commentList?: CommentList;
};

export const initialState: CommentListState = {};

export const getCommentSlice = createSlice({
  name: "getCommentList",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder.addCase(APIService.getCommentList.fulfilled, (state, action) => {
      state.commentList = action.payload;
    });
  },
});

export default getCommentSlice.reducer;
