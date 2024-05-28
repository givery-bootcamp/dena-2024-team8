import { AppRoute } from './AppRoute';
import MenuBtn from '../components/button/menuButton';
import './App.scss';

function App() {
  return (
    <div className="app-root">
      <header className="app-header">
        <MenuBtn />
        サンプルアプリケーション
      </header>
      <main className="app-body">
        <div className="container-xxl">
          <div className="row">
            <div className="col-2 border text-center" style={{ height: '100vh' }}>
              {/* Twitterライクなメニューを作成 */}
              <nav className="nav flex-column">
                <a className="nav-link" href="/">
                  <a className="btn flat" >
                    Home
                  </a>
                </a>
                <a className="nav-link" href="/search">Search</a>
              </nav>
            </div>
            <div className="col-8 border" style={{ backgroundColor: 'white', height: '100vh', padding: '1%' }}>
              <div className="mb-3" style={{ margin: '1%' }}>
                <label htmlFor="exampleFormControlTextarea1" className="form-label">今の気持ちは？？</label>
                <textarea className="form-control" id="exampleFormControlTextarea1">

                </textarea>
                <button type="button" className="btn btn-primary">Button</button>
              </div>
              <hr />
              <AppRoute />
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">ツイートのタイトル</h5>
                  <p className="card-text">ツイートの本文</p>
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">ツイートのタイトル</h5>
                  <p className="card-text">ツイートの本文</p>
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">ツイートのタイトル</h5>
                  <p className="card-text">ツイートの本文</p>
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">ツイートのタイトル</h5>
                  <p className="card-text">ツイートの本文</p>
                </div>
              </div>
              <div className="card">
                <div className="card-body">
                  <h5 className="card-title">ツイートのタイトル</h5>
                  <p className="card-text">ツイートの本文</p>
                </div>
              </div>
            </div>
            <div className="col-2 border" style={{ height: '100vh' }}>
            </div>
          </div>
        </div>

      </main>
    </div>
  );
}

export default App;
