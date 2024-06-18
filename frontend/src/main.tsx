import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { Provider } from 'react-redux';

import './main.scss';
import { store } from './shared/store';
import { AppRoute } from './app/AppRoute';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <BrowserRouter>
    <Provider store={store}>
      <AppRoute />
    </Provider>
  </BrowserRouter>
);