import { configureStore } from "@reduxjs/toolkit";

import helloReducer, { helloSlice } from "./HelloSlice";
import postReducer, { postSlice } from "./PostsSlice";
import postDetailReducer, { postDetailSlice } from "./PostDetailSlice";
import searchPostReducer, { searchPostSlice } from "./searchPostsSlice";
import signinReducer, { signinSlice } from "./SigninSlice";
import createPostReducer,{createPostSlice} from "./CreatePostSlice";

export const store = configureStore({
  reducer: {
    hello: helloReducer,
    post: postReducer,
    detail: postDetailReducer,
    searchPostList: searchPostReducer,
    singin: signinReducer,
    createPost: createPostReducer,
  },
});

export const actions = {
  ...helloSlice.actions,
  ...postSlice.actions,
  ...postDetailSlice.actions,
  ...searchPostSlice.actions,
  ...signinSlice.actions,
  ...createPostSlice.actions,
};

export type RootState = ReturnType<typeof store.getState>;

export type AppDispatch = typeof store.dispatch;
