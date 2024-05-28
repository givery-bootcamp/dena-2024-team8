import { Routes, Route } from 'react-router-dom';

import { HelloWorld } from '../features/helloworld';
import { Home } from '../features/home';
export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/search" element={<HelloWorld />}/>
    </Routes>
  );
};
