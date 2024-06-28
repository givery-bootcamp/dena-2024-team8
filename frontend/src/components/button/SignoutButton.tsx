import { useAppDispatch } from "../../shared/hooks";
import { APIService } from "../../shared/services";

export default function SignoutButton() {
  const dispatch = useAppDispatch();

  const handleSignOut = () => {
    dispatch(APIService.signout());
  };
  return (
    <button
      onClick={() => {
        handleSignOut();
      }}
      className="lg:justify-end bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
    >
      Logout
    </button>
  );
}
