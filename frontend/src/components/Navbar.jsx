import { Link } from "react-router-dom";

function Navbar() {
  const user = localStorage.getItem("user");

  return (
    <div style={{ display: "flex", gap: "20px", padding: "10px", borderBottom: "1px solid #ccc" }}>
      <Link to="/">Home</Link>

      {!user && <Link to="/login">Login</Link>}

      {user && (
        <>
          <Link to="/dashboard">Dashboard</Link>
          <button
            onClick={() => {
              localStorage.removeItem("user");
              window.location.href = "/";
            }}
          >
            Logout
          </button>
        </>
      )}
    </div>
  );
}

export default Navbar;