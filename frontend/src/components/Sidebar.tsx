export default function Sidebar() {
    return (
        <nav className="nav flex-column">
          <a className="nav-link" href="/">
            <a className="btn flat">
              Home
            </a>
          </a>
          <a className="nav-link" href="/search">Search</a>
        </nav>
    );
}