import React from 'react';
import { OrderSearchFilter } from 'components/filter';
import { OrderDatatable, OrderDatatableHeader } from 'components/datatable';

function OrderListPage() {
  const header: OrderDatatableHeader[] = [
    {
      column: 'oid',
      text: '주문번호',
      align: 'left',
    },
    {
      column: 'orderer',
      text: '주문업체',
    },
    {
      column: 'receiver',
      text: '수령업체',
    },
    {
      column: 'drugName',
      text: '제품명',
    },
    {
      column: 'quantity',
      text: '제품수량',
      align: 'right',
    },
    {
      column: 'createdAt',
      text: '등록일자',
    },
  ];

  return (
    <div className="Order__List h-full grid grid__Two__Column__Filter__Layout">
      <div className="Rounded__Shadow__Border p-4">
        <OrderSearchFilter />
      </div>
      <div className="Rounded__Shadow__Border p-4 ml-6">
        <OrderDatatable header={header} />
      </div>
    </div>
  );
}

export default OrderListPage;
