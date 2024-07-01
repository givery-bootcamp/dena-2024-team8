import { createSlice } from "@reduxjs/toolkit";
import { APIService } from "../services";

export type CreatePostState = {
  isDeleted: boolean;
  error_message?: string;
};

export const initialState: CreatePostState = { isDeleted: false };

export const deletePostSlice = createSlice({
  name: "post",
  initialState,
  reducers: {
    resetDeletePostState: (state) => {
      state.isDeleted = false;
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(APIService.deletePost.pending, (state) => {
        // 削除処理が開始されたことを示す
        state.isDeleted = false;
      })
      .addCase(APIService.deletePost.fulfilled, (state) => {
        state.isDeleted = true;
      })
      .addCase(APIService.deletePost.rejected, (state, action) => {
        state.isDeleted = false;
        state.error_message = action.error.message;
      });
  },
});

export default deletePostSlice.reducer;
