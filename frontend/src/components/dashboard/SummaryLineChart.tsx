import React from 'react';
import { useGetCharts } from 'data/dashboard/dashboard.hooks';
import { ChartType } from 'data/dashboard/dashboard.model';

interface SummaryLineChartProps {
  chartType: ChartType;
}

function SummaryLineChart({ chartType }: SummaryLineChartProps) {
  const { data } = useGetCharts(chartType);
  console.log(data);
  return <></>;
}

export default SummaryLineChart;
