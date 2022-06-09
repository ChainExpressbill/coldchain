import React from 'react';
import DatePicker, { registerLocale } from 'react-datepicker';
import classnames from 'classnames';
import { useOrderStore } from 'store';
import shallow from 'zustand/shallow';
import { subDays } from 'date-fns/esm';
import ko from 'date-fns/locale/ko';
import 'react-datepicker/dist/react-datepicker.css';
import { DATE_TYPE } from 'constants/date';

registerLocale('ko', ko);

function FilterDatePicker() {
  const { startDate, setStartDate, endDate, setEndDate } = useOrderStore(
    (state) => ({
      startDate: state.startDate,
      setStartDate: state.setStartDate,
      endDate: state.endDate,
      setEndDate: state.setEndDate,
    }),
    shallow,
  );

  const datePickerButtonClassnames = classnames(
    'bg-main rounded-lg shadow-basic text-white p-2',
  );

  const datePickerBoxClassnames = classnames(
    'flex flex-row mb-4 Range__Date__Picker',
  );

  const datePickerClassnames = classnames(
    'border-main rounded-lg shadow-basic border-2 w-[98%] p-1',
  );

  const onDatePickerButtonClick = (type: keyof typeof DATE_TYPE) => {
    switch (type) {
      case DATE_TYPE.TODAY:
        setStartDate(new Date());
        setEndDate(new Date());
        break;
      case DATE_TYPE.WEEK:
        setStartDate(subDays(endDate, 7));
        break;
      case DATE_TYPE.MONTH:
        setStartDate(subDays(endDate, 30));
        break;
    }
  };

  return (
    <>
      <div className="grid grid-cols-3 gap-3 mb-4">
        <button
          type="button"
          className={datePickerButtonClassnames}
          onClick={() => onDatePickerButtonClick(DATE_TYPE.TODAY)}
        >
          오늘
        </button>
        <button
          type="button"
          className={datePickerButtonClassnames}
          onClick={() => onDatePickerButtonClick(DATE_TYPE.WEEK)}
        >
          일주일
        </button>
        <button
          type="button"
          className={datePickerButtonClassnames}
          onClick={() => onDatePickerButtonClick(DATE_TYPE.MONTH)}
        >
          1개월
        </button>
      </div>
      <div className={datePickerBoxClassnames}>
        <DatePicker
          locale="ko"
          dateFormat="yyyy-MM-dd"
          todayButton="오늘"
          className={datePickerClassnames}
          selectsStart
          selected={startDate}
          startDate={startDate}
          endDate={endDate}
          maxDate={endDate}
          onChange={(date: Date) => setStartDate(date)}
          placeholderText="시작 일자"
        />
        <DatePicker
          locale="ko"
          dateFormat="yyyy-MM-dd"
          todayButton="오늘"
          className={datePickerClassnames}
          selectsEnd
          selected={endDate}
          startDate={startDate}
          endDate={endDate}
          minDate={startDate}
          onChange={(date: Date) => setEndDate(date)}
          placeholderText="종료 일자"
        />
      </div>
    </>
  );
}

export default FilterDatePicker;
