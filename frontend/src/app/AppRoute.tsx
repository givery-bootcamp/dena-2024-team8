import { Routes, Route } from "react-router-dom";

import { MainLayout } from "./MainLayout";
import { Home } from "../features/home";
import { Signin } from "../features/signin";
export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/" element={<Home />} />
        <Route path="/signin" element={<Signin />} />
      </Route>
    </Routes>
  );
};
