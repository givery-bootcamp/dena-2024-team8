import { createSlice } from '@reduxjs/toolkit';

import { User } from '../models';
import { APIService } from '../services';

export type SigninState = {
    Signin ?: User;
}

export const initialState : SigninState =  {}; 

export const signinSlice = createSlice({
    name: 'signin',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(APIService.signin.fulfilled, (state, action) => {
            state.Signin = action.payload;
        });
    },
});

export default signinSlice.reducer;