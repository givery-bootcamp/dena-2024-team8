import { Outlet } from "react-router-dom";
import { Header } from "../features/Header";
import { Sidebar } from "../components/Sidebar";

export const MainLayout = () => {
  return (
    <>
      <Header />
      <div className="flex-1">
        <div className="flex flex-row">
          <div className="col-2 border text-center">
            <Sidebar />
          </div>
          <div className="flex-1 w-full">
            <Outlet />
          </div>
        </div>
      </div>
    </>
  );
};
