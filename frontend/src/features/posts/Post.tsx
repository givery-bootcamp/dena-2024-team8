import { useEffect } from "react";
import { useParams } from "react-router-dom"
import { useAppDispatch, useAppSelector } from "../../shared/hooks";
import { APIService } from "../../shared/services";

export const Post = () => {
    const { postId } = useParams();
    const { postDetail } = useAppSelector((state) => state.detail);
    const dispatch = useAppDispatch();

    // ここでpostIdを使ってデータベースから対象のpostの詳細を取得
    useEffect(() => {
        if(postId != undefined){
            dispatch(APIService.getPostDetail(postId));
        }
    },[dispatch,postId])

    return (
        <>
            {postDetail?.UserId}
        </>
    )
}