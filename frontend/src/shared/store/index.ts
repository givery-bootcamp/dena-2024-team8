import { configureStore } from "@reduxjs/toolkit";

import helloReducer, { helloSlice } from "./HelloSlice";
import postReducer, { postSlice } from "./PostsSlice";
import postDetailReducer, { postDetailSlice } from "./PostDetailSlice";
import signinReducer, { signinSlice } from "./SigninSlice";
import deletePostReducer, { deletePostSlice } from "./DeletePostSlice";
import userReducer, { userSlice } from "./UserSlice";
import searchPostReducer, { searchPostSlice } from "./searchPostsSlice";
import createPostReducer, { createPostSlice } from "./CreatePostSlice";

export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    detail: postDetailReducer,
    deletePost: deletePostReducer,
    searchPostList: searchPostReducer,
    singin: signinReducer,
    user: userReducer,
    createPost: createPostReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postSlice.actions,
  ...postDetailSlice.actions,
  ...deletePostSlice.actions,
  ...searchPostSlice.actions,
  ...signinSlice.actions,
  ...userSlice.actions,
  ...createPostSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
