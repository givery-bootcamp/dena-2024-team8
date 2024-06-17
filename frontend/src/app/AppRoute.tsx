import { Routes, Route } from "react-router-dom";

import { MainLayout } from "./MainLayout";
import { Home } from "../features/home";

export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/" element={<Home />} />
      </Route>
    </Routes>
  );
};
