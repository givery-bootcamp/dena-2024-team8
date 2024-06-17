export default function Sidebar() {
    return (
        <nav className="nav flex-column">
          <a className="nav-link btn flat" href="/">
            Home
          </a>
          <a className="nav-link" href="/search">Search</a>
        </nav>
    );
}