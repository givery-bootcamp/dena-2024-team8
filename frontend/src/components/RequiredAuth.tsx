import { Outlet } from "react-router-dom";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector} from "../shared/hooks";
import { APIService } from "../shared/services";
import { useNavigate } from "react-router-dom";

export const RequiredAuth = () => {
    const navigator = useNavigate();
    const dispatch = useAppDispatch();
    const { user, error } = useAppSelector((state) => state.user);
    // ユーザーがログインしているかどうかを判定する
    useEffect(() => {
        const fetchData = async () => {
            try {
                await dispatch(APIService.getUser())
            } catch (error) {
                console.log("Error fetching user:", error);
            }
        };
        fetchData();
    }, [dispatch]);

    useEffect(() => {
        if (user === undefined && error === undefined) {
            console.log("user is undefined");
        } else if (error !== undefined) {
            navigator('/signin');
        } else {
            navigator('/');
        }
    }, [user, error, navigator]);

    return (
        <Outlet />
    );
}