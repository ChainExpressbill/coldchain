import React, { ChangeEvent } from 'react';
import classNames from 'classnames';
import { FilterDatePicker } from './FilterDatePicker';
import { useOrderStore } from 'store';
import shallow from 'zustand/shallow';

function OrderSearchFilter() {
  const { clearFilter, orderer, setOrderer, receiver, setReceiver } =
    useOrderStore(
      (state) => ({
        orderer: state.orderer,
        receiver: state.receiver,
        setOrderer: state.setOrderer,
        setReceiver: state.setReceiver,
        clearFilter: state.clearFilter,
      }),
      shallow,
    );

  const btnBaseClassnames = classNames('w-16 rounded-lg shadow-basic p-2');
  const resetBtnClassnames = classNames(btnBaseClassnames, 'border-main mr-2');
  const searchBtnClassnames = classNames(
    btnBaseClassnames,
    'bg-main text-white',
  );

  return (
    <form className="flex flex-col relative h-full py-1">
      <div className="mb-2">
        <label className="block text-xl mb-1">주문일자</label>
        <FilterDatePicker />
      </div>
      <div className="mb-2">
        <label className="block text-xl mb-1" htmlFor="orderer">
          주문자명
        </label>
        <input
          className="Form__Input__Border__Box"
          type="text"
          name="orderer"
          value={orderer}
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setOrderer(e.target.value)
          }
        />
      </div>
      <div className="mb-2">
        <label className="block text-xl mb-1" htmlFor="receiver">
          수령자명
        </label>
        <input
          className="Form__Input__Border__Box"
          type="text"
          name="receiver"
          value={receiver}
          onChange={(e: ChangeEvent<HTMLInputElement>) =>
            setReceiver(e.target.value)
          }
        />
      </div>
      <div className="absolute bottom-0 right-0">
        <button
          type="button"
          className={resetBtnClassnames}
          onClick={clearFilter}
        >
          초기화
        </button>
        <button type="button" className={searchBtnClassnames}>
          조회
        </button>
      </div>
    </form>
  );
}

export default OrderSearchFilter;
