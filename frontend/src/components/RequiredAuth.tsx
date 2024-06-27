import { Outlet } from "react-router-dom";
import { useEffect } from "react";
import { useAppDispatch, useAppSelector} from "../shared/hooks";
import { APIService } from "../shared/services";
import { useNavigate } from "react-router-dom";

export const RequiredAuth = () => {
    const navigator = useNavigate();
    const dispatch = useAppDispatch();
    const { error } = useAppSelector((state) => state.user);
    // ユーザーがログインしているかどうかを判定する
    useEffect(() => {
        dispatch(APIService.getUser())
        if (error) {
        navigator('/signin');
        }
    }, [dispatch, error]);

    return (
        <Outlet />
    );
}