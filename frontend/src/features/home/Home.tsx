import { useEffect } from 'react';

import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';

export function Home() {
  const { hello } = useAppSelector((state) => state.hello);
  const dispatch = useAppDispatch();

  useEffect(() => {
    dispatch(APIService.getHello());
  }, [dispatch]);

  return <div><p>Now Home Page</p>{hello?.message}</div>;
}
