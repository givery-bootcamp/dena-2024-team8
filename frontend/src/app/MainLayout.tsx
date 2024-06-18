import { Outlet } from "react-router-dom";
import { Header } from "../components/Header";
import { Sidebar } from "../components/Sidebar";

export const MainLayout = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1">
        <div className="flex flex-row">
          <div className="col-2 border text-center">
            <Sidebar />
          </div>
          <div className="flex-1 w-full">
            <Outlet />
          </div>
        </div>
      </main>
    </div>
  );
};