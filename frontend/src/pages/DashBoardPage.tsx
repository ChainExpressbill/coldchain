import React from 'react';
import {
  SummaryBoard,
  SummaryOrderChart,
  SummaryOrdererChart,
} from 'components/dashboard';
import { DATE_TYPE } from 'constants/date';

function DashBoard() {
  return (
    <div className="dashboard h-full grid grid__Two__Row">
      <div className="Summary__Board flex flex-row items-center">
        <SummaryBoard boardType={DATE_TYPE.LAST_MONTH} />
        <SummaryBoard boardType={DATE_TYPE.TODAY} />
      </div>
      <div className="Summary__Chart grid grid-cols-2">
        <SummaryOrderChart />
        <SummaryOrdererChart />
      </div>
    </div>
  );
}

export default DashBoard;
