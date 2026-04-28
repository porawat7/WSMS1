export async function createOrder(order) {
  const res = await fetch('/orders', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(order),
  });
  if (!res.ok) throw new Error('Failed to create order');
  return res.json();
}

export async function fetchOrders() {
  const res = await fetch('/orders');
  if (!res.ok) throw new Error('Failed to fetch orders');
  return res.json();
}
