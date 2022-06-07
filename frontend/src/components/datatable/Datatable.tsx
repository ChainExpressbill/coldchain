import React from 'react';
import { useGetOrders } from 'data/order/order.hooks';
import { useOrderStore } from 'store';
import shallow from 'zustand/shallow';

function Datatable() {
  const { startDate, endDate, orderer, receiver, page, size } = useOrderStore(
    (state) => ({
      startDate: state.startDate,
      endDate: state.endDate,
      orderer: state.orderer,
      receiver: state.receiver,
      page: state.page,
      size: state.size,
    }),
    shallow,
  );
  useGetOrders({
    startDate: startDate.toISOString(),
    endDate: endDate.toISOString(),
    orderer,
    receiver,
    page,
    size,
  });
  return <div>data table</div>;
}

export default Datatable;
