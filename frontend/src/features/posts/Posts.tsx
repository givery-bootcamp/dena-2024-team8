import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import PostCardList from '../../components/card/PostCardList';
export default function GetPostList() {
  const { postList } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostList());
  }, [dispatch]);

  return (
    <div>
        <PostCardList itemList={postList ?? []} />
    </div>
  )
}