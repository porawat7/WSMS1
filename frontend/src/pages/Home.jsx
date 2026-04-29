import { Link } from "react-router-dom";
import Navbar from "../components/Navbar";
import Input from "../components/Input";
import Button from "../components/Button";

function Home() {
  return (
    <div style={{ textAlign: "center" }}>
       <Navbar />
      <h1>Course Platform</h1>
      <p>ยินดีต้อนรับ</p>

      <Link to="/login">
        <Button text="Login" />
      </Link>
    </div>
  );
}

export default Home;