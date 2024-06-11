import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';

// export function HelloWorld() {
//   const { hello } = useAppSelector((state) => state.hello);
//   const dispatch = useAppDispatch();

//   useEffect(() => {
//     dispatch(APIService.getHello());
//   }, [dispatch]);

//   return <div><p>Now Hello World Page</p>{hello?.message}</div>;
// }

export function HelloWorld() {
  const { postList } = useAppSelector((state) => state.post);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getPostList());
  }, [dispatch]);

  useEffect(() => {
    console.log(postList);
  }, [postList])
  
  return <div><p>Now Hello World Page</p>{postList != undefined && postList[0].CreatedAt}</div>;
}