import React, { useState } from 'react';

export default function OrderForm({ coffees, onSubmit }) {
  const [items, setItems] = useState([]);
  const [selected, setSelected] = useState('');
  const [qty, setQty] = useState(1);

  const addItem = () => {
    if (!selected || qty < 1) return;
    setItems([...items, { coffee_id: selected, quantity: qty }]);
    setSelected('');
    setQty(1);
  };

  const handleSubmit = e => {
    e.preventDefault();
    if (items.length > 0) onSubmit({ items });
  };

  return (
    <form onSubmit={handleSubmit}>
      <select value={selected} onChange={e => setSelected(e.target.value)}>
        <option value="">เลือกเมนู</option>
        {coffees.map(c => (
          <option key={c.id} value={c.id}>{c.emoji} {c.name}</option>
        ))}
      </select>
      <input type="number" min="1" value={qty} onChange={e => setQty(Number(e.target.value))} />
      <button type="button" onClick={addItem}>เพิ่ม</button>
      <ul>
        {items.map((item, i) => (
          <li key={i}>{coffees.find(c => c.id === item.coffee_id)?.name} x {item.quantity}</li>
        ))}
      </ul>
      <button type="submit">สั่งซื้อ</button>
    </form>
  );
}
