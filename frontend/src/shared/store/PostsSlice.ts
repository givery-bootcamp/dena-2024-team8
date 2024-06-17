import { createSlice } from '@reduxjs/toolkit';

import { PostList } from '../models';
import { APIService } from '../services';

export type PostState = {
  postList?: PostList;
};

export const initialState: PostState = {};

export const postSlice = createSlice({
    name: 'post',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.getPostList.fulfilled, (state, action) => {
            state.postList = action.payload;
        });
    },
});

export default postSlice.reducer;

