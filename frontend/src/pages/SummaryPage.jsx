import React, { useEffect, useState } from 'react';
import { fetchSummary } from '../api/summary';
import SummaryView from '../components/SummaryView';

export default function SummaryPage() {
  const [summary, setSummary] = useState(null);
  useEffect(() => {
    fetchSummary().then(setSummary);
  }, []);
  return <SummaryView summary={summary} />;
}
