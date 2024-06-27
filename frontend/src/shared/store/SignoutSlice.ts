import { createSlice } from '@reduxjs/toolkit';

import { SignOutResponse } from '../models';
import { APIService } from '../services';

export type SignoutState = {
    signout ?: SignOutResponse;
}

export const initialState : SignoutState =  {}; 

export const signoutSlice = createSlice({
    name: 'signout',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.signout.fulfilled, (state, action) => {
            state.signout = action.payload;
        });
    },
});

export default signoutSlice.reducer;