import { useNavigate } from "react-router-dom";
import SignoutButton from "./button/SignoutButton";
import { useMemo } from "react";
import { useAppDispatch} from "../shared/hooks";
import { APIService } from "../shared/services";

export const Header = () => {
  const navigator = useNavigate();
  const dispatch = useAppDispatch();
  // ユーザーがログインしているかどうかを判定する
  const getUser = () => {
    console.log("getUser");
    dispatch(APIService.getUser());
  }
  const user = useMemo(() => getUser(), []); 
  console.log(user);
  
  return (
    <header className="bg-gray-100">
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
