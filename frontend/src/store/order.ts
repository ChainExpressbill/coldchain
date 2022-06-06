/* eslint-disable @typescript-eslint/no-unused-vars */
import create from 'zustand';
import { devtools } from 'zustand/middleware';

export interface HeaderStore {
  currentTime: number;
  timerId: number;
  setCurrentTime: () => void;
  setTimerId: (id: number) => void;
  resetCurrentTime: () => void;
  clearTimer: () => void;
}

export interface OrderStoreState {
  startDate: Date;
  endDate: Date;
  orderer: string;
  receiver: string;
}

export interface OrderStore extends OrderStoreState {
  setStartDate: (date: Date) => void;
  setEndDate: (date: Date) => void;
  setOrderer: (orderer: string) => void;
  setReceiver: (receiver: string) => void;
  clearFilter: () => void;
}

const initialState: OrderStoreState = {
  startDate: new Date(),
  endDate: new Date(),
  orderer: '',
  receiver: '',
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
    clearFilter: () => {
      set(() => ({ ...initialState }));
    },
  })),
);
