function Input({ placeholder, type = "text", onChange }) {
  return (
    <input
      type={type}
      placeholder={placeholder}
      onChange={onChange}
      style={{
        padding: "10px",
        margin: "5px",
        width: "200px"
      }}
    />
  );
}

export default Input;
