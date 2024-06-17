import { Outlet } from "react-router-dom";
import { Header } from "../features/Header";

export const MainLayout = () => {
  return (
    <>
      <Header />
      <div className="flex-1">
        <Outlet />
      </div>
    </>
  );
};
