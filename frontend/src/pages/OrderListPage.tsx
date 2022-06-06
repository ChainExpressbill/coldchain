import React from 'react';
import { OrderSearchFilter } from 'components/filter';

function OrderListPage() {
  // optional: orderer, receiver
  // required: page, size, startDate, endDate
  return (
    <div className="Order__List h-full grid grid__Two__Row">
      <div>
        <OrderSearchFilter />
      </div>
      <div>data table</div>
    </div>
  );
}

export default OrderListPage;
