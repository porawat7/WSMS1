CREATE TABLE courses (
  id INT AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  category VARCHAR(100),
  price INT,
  level VARCHAR(20),
  description TEXT,
  platform VARCHAR(100),
  link TEXT,
  startDate DATE,
  time VARCHAR(50)
);

INSERT INTO courses (name, category, price, level, description, platform, link, startDate, time)
VALUES
('Web Dev Online', 'Coding', 1200, 'basic', 'เรียนผ่าน Zoom', 'Zoom', 'https://zoom.us/1', '2026-05-10', '18:00-20:00'),
('UX Advanced', 'Design', 2500, 'silver', 'UX ขั้นสูง', 'Meet', 'https://meet.google.com/1', '2026-05-15', '19:00-21:00'),
('AI Business', 'Business', 5000, 'gold', 'AI สำหรับธุรกิจ', 'Website', 'https://example.com', '2026-05-20', '20:00-22:00');