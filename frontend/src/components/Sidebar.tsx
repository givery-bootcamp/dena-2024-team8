import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";

export const Sidebar = () => {
  const navigator = useNavigate();
  const location = useLocation();
  const [activePath, setActivePath] = useState<string>("");

  useEffect(() => {
    setActivePath(location.pathname);
    console.log(location.pathname);
  }, [location]);

  return (
    <div className="h-full min-w-48 border-r pr-2">
      <div className="py-4 border-b-4 border-gray-300">
        <b>メニュー</b>
      </div>
      <nav>
        <ul className="mx-4 py-4 flex flex-col gap-4">
          <li>
            <button
              onClick={() => navigator("/")}
              className={`rounded py-3 border w-full active:bg-blue-200 active:shadow-inner active:shadow-black ${
                activePath === "/"
                  ? "text-blue-700 shadow-black shadow-inner bg-blue-200"
                  : "text-blue-500 shadow shadow-gray-400 hover:bg-blue-100"
              }`}
            >
              Home
            </button>
          </li>
          <li>
            <button
              onClick={() => navigator("/search")}
              className={`rounded py-3 border w-full active:bg-blue-200 active:shadow-inner active:shadow-black ${
                activePath === "/search"
                  ? "text-blue-700 shadow-black shadow-inner bg-blue-200"
                  : "text-blue-500 shadow shadow-gray-400 hover:bg-blue-100"
              }`}
            >
              Search
            </button>
          </li>
          <li>
            <button
              onClick={() => navigator("/signin")}
              className={`rounded py-3 border w-full active:bg-blue-200 active:shadow-inner active:shadow-black ${
                activePath === "/signin"
                  ? "text-blue-700 shadow-black shadow-inner bg-blue-200"
                  : "text-blue-500 shadow shadow-gray-400 hover:bg-blue-100"
              }`}
            >
              test
            </button>
          </li>
        </ul>
      </nav>
    </div>
  );
};
