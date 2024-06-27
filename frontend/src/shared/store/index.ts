import { configureStore } from '@reduxjs/toolkit';

import helloReducer, { helloSlice } from './HelloSlice';
import postReducer, { postSlice } from './PostsSlice';
import postDetailReducer,{ postDetailSlice } from './PostDetailSlice'
import signinReducer,{ signinSlice } from './SigninSlice'
import userReducer,{ userSlice } from './UserSlice'
export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    detail : postDetailReducer,
    singin: signinReducer,
    user: userReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postSlice.actions,
  ...postDetailSlice.actions,
  ...signinSlice.actions,
  ...userSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
