import create from 'zustand';
import { devtools } from 'zustand/middleware';
import { DATATABLE_VIEW_ROWS } from 'constants/pagination';

export interface OrderStoreState {
  startDate: Date;
  endDate: Date;
  orderer: string;
  receiver: string;
  page: number;
  size: number;
}

export interface OrderStore extends OrderStoreState {
  setStartDate: (date: Date) => void;
  setEndDate: (date: Date) => void;
  setOrderer: (orderer: string) => void;
  setReceiver: (receiver: string) => void;
  setPage: (page: number) => void;
  setSize: (size: number) => void;
  clearFilter: () => void;
}

const initialState: OrderStoreState = {
  startDate: new Date(),
  endDate: new Date(),
  orderer: '',
  receiver: '',
  page: 1,
  size: DATATABLE_VIEW_ROWS.DEFAULT,
};

export const useOrderStore = create<OrderStore>(
  devtools((set) => ({
    ...initialState,
    setStartDate: (date: Date) => {
      set(() => ({ startDate: date }));
    },
    setEndDate: (date: Date) => {
      set(() => ({ endDate: date }));
    },
    setOrderer: (orderer: string) => {
      set(() => ({ orderer }));
    },
    setReceiver: (receiver: string) => {
      set(() => ({ receiver }));
    },
    setPage: (page: number) => {
      set(() => ({ page }));
    },
    setSize: (size: number) => {
      set(() => ({ size, page: 1 }));
    },
    clearFilter: () => {
      set(() => ({ ...initialState }));
    },
  })),
);
