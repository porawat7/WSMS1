import { useState } from "react";
import { useNavigate } from "react-router-dom";
import Input from "../components/Input";
import Button from "../components/Button";

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleLogin = async () => {
    const res = await fetch("http://localhost:3000/api/auth/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ email, password }),
    });

    const data = await res.json();

    if (res.ok) {
      localStorage.setItem("user", JSON.stringify(data));
      navigate("/dashboard");
    } else {
      alert("Login failed");
    }
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>Login</h2>

      <Input placeholder="Email" onChange={(e) => setEmail(e.target.value)} />
      <br />

      <Input type="password" placeholder="Password" onChange={(e) => setPassword(e.target.value)} />
      <br />

      <Button text="Login" onClick={handleLogin} />
    </div>
  );
}

export default Login;