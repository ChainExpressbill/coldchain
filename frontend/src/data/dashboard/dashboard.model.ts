export const enum ChartType {
  ORDERS = 'orders',
  ORDERERS = 'orderers',
  SUCCESSES = 'successes',
  FAILURES = 'failures',
}

export interface SummaryCountResponse {
  orderCount: number;
  ordererCount: number;
}

export interface ChartsResponse {
  timestamp: string; // UTC string => new Date().toISOString()
  orderer?: string;
  count: number;
}
