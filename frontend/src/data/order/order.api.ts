import axios from 'utils/axios';
import {
  OrderRequestBody,
  OrderSearchParams,
  Order,
  GetOrdersResponse,
} from './order.model';

export async function getOrders(params: OrderSearchParams) {
  const { data } = await axios.get<GetOrdersResponse>('/orders', {
    params,
  });
  return data;
}

export async function getOrderDetail(id: number) {
  const { data } = await axios.get<Order>(`/orders/${id}`);
  return data;
}

export async function orderCreate(orderRequestBody: OrderRequestBody) {
  const { data } = await axios.post('/orders', orderRequestBody);
  return data;
}

export async function orderUpdate(orderRequestBody: OrderRequestBody) {
  const { data } = await axios.put('/orders', orderRequestBody);
  return data;
}
