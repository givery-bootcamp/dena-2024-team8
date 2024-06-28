import { createSlice } from '@reduxjs/toolkit';

import { Post } from '../models';
import { APIService } from '../services';

export type PostDetailState = {
    postDetail ?: Post;
}

export const initialState : PostDetailState =  {}; 

export const postDetailSlice = createSlice({
    name: 'post',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.getPostDetail.fulfilled, (state, action) => {
            state.postDetail = action.payload;
        });
    },
});

export default postDetailSlice.reducer;