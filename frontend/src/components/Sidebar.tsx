export default function Sidebar() {
    return (
        <nav className="nav flex-column">
          <ul>
            <li>
              <a className="nav-link btn flat" href="/">
                Home
              </a>
            </li>
            <li>
              <a className="nav-link" href="/search">Search</a>
            </li>
          </ul>
        </nav>
    );
}