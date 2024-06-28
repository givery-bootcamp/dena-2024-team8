import { Routes, Route } from "react-router-dom";
import { MainLayout } from "./MainLayout";
import { Home } from "../features/home";
import { PostDetailRouter } from "../features/posts/PostDetailRouter";
import { Signin } from "../features/signin";
import { SearchRouter } from "../features/search/SearchRouter";

export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/" element={<Home />} />
        <Route path="/search" element={<SearchRouter />} />
        <Route path="/posts/:postId" element = {<PostDetailRouter />}/>
        <Route path="/signin" element={<Signin />} />
      </Route>
    </Routes>
  );
};