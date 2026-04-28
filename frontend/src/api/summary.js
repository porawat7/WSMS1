export async function fetchSummary() {
  const res = await fetch('/summary');
  if (!res.ok) throw new Error('Failed to fetch summary');
  return res.json();
}
