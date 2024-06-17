import { AppRoute } from './AppRoute';
import MenuButton from '../components/button/MenuButton';
import './App.scss';
import { FaBeer } from 'react-icons/fa';

function App() {
  return (
    <div className="app-root">
      <header className="app-header">
        <MenuButton />
        <FaBeer />
      </header>
      <main className="app-body">
        <AppRoute /> 
      </main>
    </div>
  );
}
export default App;
