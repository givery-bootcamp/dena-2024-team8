import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
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
      <div className="col-2 border text-center" style={{ height: '100vh' }}>
        <Sidebar />
      </div>
      <div className="col-8 border" style={{ backgroundColor: 'white', height: '100vh', padding: '1%' }}>
        <PostForm />
        <hr />
        <PostList />
      </div>
      <div className="col-2 border" style={{ height: '100vh' }}>
      </div>
    </div>
    </div>
  );
}
