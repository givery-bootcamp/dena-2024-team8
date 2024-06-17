import { useEffect } from 'react';

import { useAppDispatch } from '../../shared/hooks';
import { APIService } from '../../shared/services';
import Sidebar from '../../components/Sidebar';
import PostForm from '../../components/form/PostForm';
import PostList from '../../features/posts/Posts';

export function Home() {

  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getHello());
  }, [dispatch]);

  return (
    <div className="container-xxl">
    <div className="row">
      <div className="col-2 border text-center">
        <Sidebar />
      </div>
      <div className="col-8 border bg-white vh-100 p-3">
        <PostForm />
        <hr />
        <PostList />
      </div>
      <div className="col-2 border" >
      </div>
    </div>
    </div>
  );
}
