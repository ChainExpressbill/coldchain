import { useQuery, UseQueryOptions } from 'react-query';
import { AxiosError } from 'axios';
import {
  SummaryCountResponse,
  ChartType,
  ChartsResponse,
} from './dashboard.model';
import { getLastMonth, getToday, getCharts } from './dashboard.api';

export function useGetLastMonth(
  options:
    | UseQueryOptions<SummaryCountResponse, AxiosError, SummaryCountResponse>
    | undefined = {},
) {
  return useQuery<SummaryCountResponse, AxiosError, SummaryCountResponse>(
    ['/dashboard/summary/last-month'],
    getLastMonth,
    {
      retry: 2,
      ...options,
    },
  );
}

export function useGetToday(
  options:
    | UseQueryOptions<SummaryCountResponse, AxiosError, SummaryCountResponse>
    | undefined = {},
) {
  return useQuery<SummaryCountResponse, AxiosError, SummaryCountResponse>(
    ['/dashboard/summary/today'],
    getToday,
    {
      retry: 2,
      ...options,
    },
  );
}

export function useGetCharts(
  chartType: ChartType,
  options:
    | UseQueryOptions<ChartsResponse[], AxiosError, ChartsResponse[]>
    | undefined = {},
) {
  return useQuery<ChartsResponse[], AxiosError, ChartsResponse[]>(
    ['/dashboard/summary/charts'],
    () => getCharts(chartType),
    {
      retry: 2,
      ...options,
    },
  );
}
