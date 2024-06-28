import { createSlice } from '@reduxjs/toolkit';

import { User } from '../models';
import { APIService } from '../services';

export type UserState = {
    user ?: User;
    error ?: string;
}

export const initialState : UserState =  {}; 

export const userSlice = createSlice({
    name: 'getUser',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.getUser.fulfilled, (state, action) => {
            state.user = action.payload;
            state.error = undefined;
        });
        builder.addCase(APIService.getUser.rejected, (state, action) => {
            state.error = action.error.message;
            state.user = undefined;
            console.log("error", action.error.message);
        });
    },
});

export default userSlice.reducer;