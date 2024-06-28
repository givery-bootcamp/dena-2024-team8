import { configureStore } from "@reduxjs/toolkit";

import helloReducer, { helloSlice } from "./HelloSlice";
import postReducer, { postSlice } from "./PostsSlice";
import postDetailReducer, { postDetailSlice } from "./PostDetailSlice";
import signinReducer, { signinSlice } from "./SigninSlice";
import userReducer, { userSlice } from "./UserSlice";
import searchPostReducer, { searchPostSlice } from "./searchPostsSlice";
import createPostReducer, { createPostSlice } from "./CreatePostSlice";
import updatePostReducer, { updatePostSlice } from "./UpdatePostSlice";

export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    detail: postDetailReducer,
    searchPostList: searchPostReducer,
    singin: signinReducer,
    user: userReducer,
    createPost: createPostReducer,
    updatePost: updatePostReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postSlice.actions,
  ...postDetailSlice.actions,
  ...searchPostSlice.actions,
  ...signinSlice.actions,
  ...userSlice.actions,
  ...createPostSlice.actions,
  ...updatePostSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
