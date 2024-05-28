import { AppRoute } from './AppRoute';
import MenuBtn from '../components/button/MenuButton';
import './App.scss';
import { FaBeer } from 'react-icons/fa';

function App() {
  return (
    <div className="app-root">
      <header className="app-header">
        <MenuBtn />
        <FaBeer />
      </header>
      <main className="app-body">
        <AppRoute /> 
      </main>
    </div>
  );
}
export default App;
