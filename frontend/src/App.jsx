import React from 'react';
import MenuPage from './pages/MenuPage';
import OrderPage from './pages/OrderPage';
import SummaryPage from './pages/SummaryPage';

export default function App() {
  const [page, setPage] = React.useState('menu');
  return (
    <div>
      <nav>
        <button onClick={() => setPage('menu')}>เมนู</button>
        <button onClick={() => setPage('order')}>สั่งกาแฟ</button>
        <button onClick={() => setPage('summary')}>สรุปยอด</button>
      </nav>
      <hr />
      {page === 'menu' && <MenuPage />}
      {page === 'order' && <OrderPage />}
      {page === 'summary' && <SummaryPage />}
    </div>
  );
}
