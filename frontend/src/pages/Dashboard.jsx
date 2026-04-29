import { useEffect, useState } from "react";
import CourseCard from "../components/CourseCard";

function Dashboard() {
  const [courses, setCourses] = useState([]);

  useEffect(() => {
    const user = localStorage.getItem("user");

    if (!user) {
      window.location.href = "/login";
      return;
    }

    fetch("http://localhost:3000/api/courses")
      .then((res) => res.json())
      .then((data) => setCourses(data));
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h2>Courses</h2>

      {courses.map((c) => (
        <CourseCard key={c.id} course={c} />
      ))}
    </div>
  );
}

export default Dashboard;