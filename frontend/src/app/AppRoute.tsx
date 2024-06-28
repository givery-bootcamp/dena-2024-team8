import { Routes, Route } from "react-router-dom";
import { MainLayout } from "./MainLayout";
import { Home } from "../features/home";
import { Post } from "../features/posts/Post";
import { Signin } from "../features/signin";
import { RequiredAuth } from "../components/RequiredAuth";
import { NotRequiredAuth } from "../components/NotRequiredAuth";
export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<RequiredAuth />}>
        <Route element={<MainLayout />}>
          <Route path="/" element={<Home />} />
          <Route path="/search" element={<Home />} />
          <Route path="/posts/:postId" element = {<Post />}/>
        </Route>
      </Route>
      <Route element={<NotRequiredAuth />}>
        <Route path="/signin" element={<Signin />} />
      </Route>
    </Routes>
  );
};