import { Outlet } from "react-router-dom";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector} from "../shared/hooks";
import { APIService } from "../shared/services";
import { useNavigate } from "react-router-dom";

export const NotRequiredAuth = () => {
    const navigator = useNavigate();
    const dispatch = useAppDispatch();
    const { user, error } = useAppSelector((state) => state.user);
    // ユーザーがログインしているかどうかを判定する
    useEffect(() => {
        const fetchData = async () => {
            try {
                await dispatch(APIService.getUser());
                console.log("user:", user);
                console.log("error:", error);
                if (user === undefined) {
                    navigator('/signin');
                } else {
                    navigator('/');
                }
            } catch (error) {
                console.log("Error fetching user:", error);
            }
        };
        fetchData();
    }, [dispatch, error]);
    return (
        <Outlet />
    );
}