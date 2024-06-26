import { createSlice } from '@reduxjs/toolkit';

import { ErrorResponse, User } from '../models';
import { APIService } from '../services';

export type UserState = {
    User ?: User | ErrorResponse;
}

export const initialState : UserState =  {}; 

export const userSlice = createSlice({
    name: 'getUser',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.getUser.fulfilled, (state, action) => {
            state.User = action.payload;
        });
    },
});

export default userSlice.reducer;