import { Routes, Route } from 'react-router-dom';

import { Home } from '../features/home';
export const AppRoute = () => {
  return (
    <Routes>
      <Route path="/" element={<Home />} />
    </Routes>
  );
};
