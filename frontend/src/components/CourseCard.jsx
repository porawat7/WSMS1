function CourseCard({ course }) {
  return (
    <div style={{
      border: "1px solid #ddd",
      padding: "15px",
      borderRadius: "8px",
      marginBottom: "10px"
    }}>
      <h3>{course.name}</h3>
      <p>Category: {course.category}</p>
      <p>Price: {course.price} บาท</p>
    </div>
  );
}

export default CourseCard;