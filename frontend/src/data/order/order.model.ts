export interface OrderSearchParams {
  orderer?: string;
  receiver?: string;
  startDate: string;
  endDate: string;
  page: number;
  size: number;
}

export interface Order {
  id: number; // 주문에 대한 db column id
  oid: string; // UUID -> 주문 번호
  orderer: string;
  receiver: string;
  drugName: string;
  drugStandard: string;
  quantity: number;
  registerName: string;
  storageCondition: string;
  createdAt: string; // UTC string => new Date().toISOString()
  updatedAt: string; // UTC string => new Date().toISOString()
}

export interface GetOrdersResponse {
  orderList: Order[];
  page: number;
  size: number;
  totalPage: number;
  totalCount: number;
}

export interface OrderRequestBody {
  // jwt 적용 전 임시로 id를 받아서 처리
  accountId: string;
  id: number;
  oid: string;
  orderer: string;
  receiver: string;
  drugName: string;
  drugStandard: string;
  quantity: number;
  registerName: string;
  storageCondition: string;
}
