import { Routes, Route } from "react-router-dom";

import { HelloWorld } from "../features/helloworld";
import { MainLayout } from "./MainLayout";

export const AppRoute = () => {
  return (
    <Routes>
      <Route element={<MainLayout />}>
        <Route path="/" element={<HelloWorld />} />
      </Route>
    </Routes>
  );
};
