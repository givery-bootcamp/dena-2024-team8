import { Routes, Route } from "react-router-dom";
import { MainLayout } from "./MainLayout";
import { Home } from "../features/home";
import { Post } from "../features/posts/Post";
import { Signin } from "../features/signin";
import { RequiredAuth } from "../components/RequiredAuth";
import { SearchRouter } from "../features/search/SearchRouter";
import { CreatePost } from "../features/posts/CreatePost";

export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<RequiredAuth />}>
        <Route element={<MainLayout />}>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<SearchRouter />} />
          <Route path="/posts/:postId" element={<Post />} />
          <Route path="/posts/new" element={<CreatePost />} />
        </Route>
      </Route>
      <Route path="/signin" element={<Signin />} />
    </Routes>
  );
};
