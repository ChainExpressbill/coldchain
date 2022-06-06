import React from 'react';
import { OrderSearchFilter } from 'components/filter';

function OrderListPage() {
  return (
    <div className="Order__List h-full grid grid__Two__Column__Filter__Layout">
      <div className="Rounded__Shadow__Border p-4">
        <OrderSearchFilter />
      </div>
      <div className="Rounded__Shadow__Border p-4 ml-6">data table</div>
    </div>
  );
}

export default OrderListPage;
