import { Outlet } from "react-router-dom";
import { Header } from "../components/Header";
import { Sidebar } from "../components/Sidebar";

export const MainLayout = () => {
  return (
    <div className="min-h-screen flex flex-col">
      <Header />
      <main className="flex-1 h-full flex pt-4 container mx-auto">
        <div className="text-center">
          <Sidebar />
        </div>
        <div className="w-full">
          <Outlet />
        </div>
      </main>
    </div>
  );
};
