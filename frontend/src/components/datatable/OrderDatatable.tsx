import React from 'react';
import { useGetOrders } from 'data/order/order.hooks';
import { useOrderStore } from 'store';
import shallow from 'zustand/shallow';
import Skeleton from 'react-loading-skeleton';
import { Order } from 'data/order/order.model';
import { format } from 'date-fns/esm';
import { useNavigate } from 'react-router-dom';
import Pagination from './Pagination';
import { DATATABLE_VIEW_ROWS } from 'constants/pagination';
import Select from 'react-select';

export type OrderKeys = keyof Order;
export interface OrderDatatableHeader<T = OrderKeys> {
  column: T;
  text: string;
  width?: string;
  align?: string;
}

interface OrderDatatableProps {
  header: OrderDatatableHeader[];
}

function OrderDatatable({ header }: OrderDatatableProps) {
  const navigator = useNavigate();
  const {
    startDate,
    endDate,
    orderer,
    receiver,
    page,
    size,
    setPage,
    setSize,
  } = useOrderStore(
    (state) => ({
      startDate: state.startDate,
      endDate: state.endDate,
      orderer: state.orderer,
      receiver: state.receiver,
      page: state.page,
      size: state.size,
      setPage: state.setPage,
      setSize: state.setSize,
    }),
    shallow,
  );
  const { data, isLoading, isError, refetch } = useGetOrders(
    {
      startDate: startDate.toISOString(),
      endDate: endDate.toISOString(),
      orderer,
      receiver,
      page,
      size,
    },
    {
      keepPreviousData: true,
      enabled: false,
    },
  );
  const options = [
    { value: DATATABLE_VIEW_ROWS.DEFAULT, label: DATATABLE_VIEW_ROWS.DEFAULT },
    { value: DATATABLE_VIEW_ROWS.MIDDLE, label: DATATABLE_VIEW_ROWS.MIDDLE },
    { value: DATATABLE_VIEW_ROWS.MAX, label: DATATABLE_VIEW_ROWS.MAX },
  ];

  React.useEffect(() => {
    refetch();
  }, [refetch, page, size]);

  const handleClick = (id: number) => {
    navigator(`${id}`);
  };

  const onSelectChange = (option: any) => {
    setSize(option.value);
  };

  const tableHeaders = React.useMemo(
    () =>
      header.map((h: OrderDatatableHeader) => (
        <th key={h.column}>
          <div>{h.text}</div>
        </th>
      )),
    [header],
  );

  if (isLoading) {
    return (
      <table className="w-full h-full table-fixed">
        <thead>
          <tr>{tableHeaders}</tr>
        </thead>
        <tbody>
          <tr>
            {header.map((h: OrderDatatableHeader) => (
              <td key={h.column}>
                <Skeleton height={30} count={20} />
              </td>
            ))}
          </tr>
        </tbody>
      </table>
    );
  }

  if (isError) {
    return (
      <div className="flex items-center justify-center">
        데이터를 불러오는데 실패했습니다.
      </div>
    );
  }

  const tableRows = data?.orderList.map((order: Order) => (
    <tr
      key={order.id}
      className="grid grid-cols-order border-b-[1px] border-b-main pl-2 py-4 hover:bg-gray-200 hover:cursor-pointer"
    >
      {header.map((h: OrderDatatableHeader) => {
        const textAlign = `${h.align ?? 'center'}` as any;

        return (
          <td
            key={h.column}
            style={{ textAlign }}
            onClick={() => handleClick(order.id)}
          >
            {h.column === 'createdAt'
              ? format(new Date(order[h.column]), 'yyyy-MM-dd')
              : order[h.column]}
          </td>
        );
      })}
    </tr>
  ));

  function renderPagination() {
    if (data) {
      return (
        <Pagination
          totalPages={data.totalPage}
          displayAmount={10}
          currentPage={data.page}
          setPage={setPage}
        />
      );
    }
  }

  return (
    <div className="w-full">
      <div className="flex flex-row items-center justify-between mb-2">
        <div>총 개수: {data?.totalCount ?? 0}</div>
        <div>
          <Select
            options={options}
            className="w-[8rem]"
            defaultValue={options[0]}
            onChange={onSelectChange}
          />
        </div>
      </div>
      <table
        className="w-full table-fixed border-0 border-collapse"
        cellSpacing="0"
        cellPadding="0"
      >
        <thead className="block">
          <tr className="grid grid-cols-order border-b-[1px] border-b-main py-4">
            {tableHeaders}
          </tr>
        </thead>
        <tbody className="block h-[43rem] overflow-auto">{tableRows}</tbody>
      </table>
      {renderPagination()}
    </div>
  );
}

export default OrderDatatable;
