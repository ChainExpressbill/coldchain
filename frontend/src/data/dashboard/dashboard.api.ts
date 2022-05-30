import axios from 'utils/axios';
import {
  ChartsResponse,
  ChartType,
  SummaryCountResponse,
} from './dashboard.model';

export async function getLastMonth() {
  const { data } = await axios.get<SummaryCountResponse>(
    '/dashboard/summary/last-month',
  );
  return data;
}

export async function getToday() {
  const { data } = await axios.get<SummaryCountResponse>(
    '/dashboard/summary/today',
  );
  return data;
}

export async function getCharts(chartType: ChartType) {
  const { data } = await axios.get<ChartsResponse>(
    `/dashboard/charts/${chartType}`,
  );
  return data;
}
