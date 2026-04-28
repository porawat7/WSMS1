import React from 'react';

export default function SummaryView({ summary }) {
  if (!summary) return null;
  return (
    <div>
      <h2>สรุปยอดขาย</h2>
      <div>จำนวนออเดอร์: {summary.total_orders}</div>
      <div>จำนวนแก้ว: {summary.total_cups}</div>
      <div>รายรับรวม: {summary.total_revenue} บาท</div>
      <div>เมนูขายดี: {summary.best_seller} ({summary.best_seller_qty} แก้ว)</div>
    </div>
  );
}
