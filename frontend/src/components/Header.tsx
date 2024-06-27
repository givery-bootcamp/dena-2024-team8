import { FaBars } from "react-icons/fa";
import { useNavigate } from "react-router-dom";
import SignoutButton from "./button/SignoutButton";

export const Header = () => {
  const navigator = useNavigate();
  return (
    <header className="bg-blue-300">
      <nav
        className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8"
        aria-label="Global"
      >
        <div className="flex lg:flex-1">
          <a href="/" className="-m-1.5 p-1.5">
            <span className="sr-only">Your Company</span>
            <img
              className="h-8 w-auto"
              src="https://tailwindui.com/img/logos/mark.svg?color=indigo&shade=600"
              alt=""
            />
          </a>
        </div>

        <div className="lg:flex lg:flex-1 lg:justify-end">
          <SignoutButton />
          <button onClick={() => {navigator('/signin')}} className="lg:justify-end bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
            Login
          </button>
        </div>
      </nav>
    </header>
  );
};
