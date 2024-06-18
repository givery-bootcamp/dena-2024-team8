import { Link } from "react-router-dom";

export const Sidebar = () => {
  return (
    <div className="w-40">
      <nav>
        <ul>
          <li className="mb-2"><Link to ={"/"} className="text-blue-500">Home</Link></li>
          <li><a href="#" className="text-blue-500">Search</a></li>
        </ul>
      </nav>
    </div>
  );
};