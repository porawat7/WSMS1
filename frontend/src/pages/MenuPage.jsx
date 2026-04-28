import React, { useEffect, useState } from 'react';
import { fetchCoffees } from '../api/coffee';
import MenuList from '../components/MenuList';

export default function MenuPage() {
  const [coffees, setCoffees] = useState([]);
  useEffect(() => {
    fetchCoffees().then(setCoffees).catch(() => setCoffees([]));
  }, []);
  return <MenuList coffees={coffees} />;
}
