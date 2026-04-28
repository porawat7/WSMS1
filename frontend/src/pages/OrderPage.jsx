import React, { useEffect, useState } from 'react';
import { fetchCoffees } from '../api/coffee';
import { createOrder } from '../api/order';
import OrderForm from '../components/OrderForm';

export default function OrderPage() {
  const [coffees, setCoffees] = useState([]);
  const [result, setResult] = useState(null);
  useEffect(() => {
    fetchCoffees().then(setCoffees);
  }, []);
  const handleSubmit = async (order) => {
    try {
      const res = await createOrder(order);
      setResult(res);
    } catch (e) {
      setResult({ error: e.message });
    }
  };
  return (
    <div>
      <OrderForm coffees={coffees} onSubmit={handleSubmit} />
      {result && <pre>{JSON.stringify(result, null, 2)}</pre>}
    </div>
  );
}
