import React from 'react';
import DatePicker from 'react-datepicker';

import 'react-datepicker/dist/react-datepicker.css';

function OrderSearchFilter() {
  const [startDate, setStartDate] = React.useState(new Date());
  return (
    <form className="flex flex-row">
      <DatePicker
        selected={startDate}
        onChange={(date: Date) => setStartDate(date)}
      />
      <DatePicker
        selected={startDate}
        onChange={(date: Date) => setStartDate(date)}
      />
    </form>
  );
}

export default OrderSearchFilter;
