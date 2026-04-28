export async function fetchCoffees() {
  const res = await fetch('/coffees');
  if (!res.ok) throw new Error('Failed to fetch coffees');
  return res.json();
}
