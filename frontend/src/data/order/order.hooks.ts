import {
  useQuery,
  UseQueryOptions,
  useMutation,
  UseMutationOptions,
} from 'react-query';
import { AxiosError } from 'axios';
import {
  GetOrdersResponse,
  Order,
  OrderRequestBody,
  OrderSearchParams,
} from './order.model';
import {
  getOrders,
  getOrderDetail,
  orderCreate,
  orderUpdate,
} from './order.api';

export function useGetOrders(
  params: OrderSearchParams,
  options:
    | UseQueryOptions<GetOrdersResponse, AxiosError, GetOrdersResponse>
    | undefined = {},
) {
  return useQuery<GetOrdersResponse, AxiosError, GetOrdersResponse>(
    ['/orders', params],
    () => getOrders(params),
    {
      retry: 2,
      ...options,
    },
  );
}

export function useOrderDetail(
  id: number,
  options: UseQueryOptions<Order, AxiosError, Order> | undefined = {},
) {
  return useQuery<Order, AxiosError, Order>(
    ['/orderDetail'],
    () => getOrderDetail(id),
    {
      retry: 2,
      ...options,
    },
  );
}

export function useOrderCreate(
  options: UseMutationOptions<unknown, AxiosError, OrderRequestBody>,
) {
  return useMutation<unknown, AxiosError, OrderRequestBody>(
    '/orderCreate',
    orderCreate,
    {
      retry: false,
      ...options,
    },
  );
}

export function useOrderUpdate(
  options: UseMutationOptions<unknown, AxiosError, OrderRequestBody>,
) {
  return useMutation<unknown, AxiosError, OrderRequestBody>(
    '/orderUpdate',
    orderUpdate,
    {
      retry: false,
      ...options,
    },
  );
}
