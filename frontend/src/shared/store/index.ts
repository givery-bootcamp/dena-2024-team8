import { configureStore } from '@reduxjs/toolkit';

import helloReducer, { helloSlice } from './HelloSlice';
import postReducer, { postSlice } from './PostsSlice';
import postDetailReducer,{ postDetailSlice } from './PostDetailSlice'
import signinReducer,{ signinSlice } from './SigninSlice'

export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    detail : postDetailReducer,
    sinin: signinReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postSlice.actions,
  ...postDetailSlice.actions,
  ...signinSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
