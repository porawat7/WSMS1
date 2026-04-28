import React from 'react';

export default function MenuList({ coffees }) {
  return (
    <ul>
      {coffees.map(coffee => (
        <li key={coffee.id}>
          {coffee.emoji} {coffee.name} - {coffee.price} บาท
        </li>
      ))}
    </ul>
  );
}
