import React from 'react';
import {
  SummaryBoard,
  SummaryOrderChart,
  SummaryOrdererChart,
} from 'components/dashboard';

function DashBoard() {
  return (
    <div className="dashboard h-full grid grid__Two__Row">
      <div className="Summary__Board flex flex-row items-center">
        <SummaryBoard boardType="lastMonth" />
        <SummaryBoard boardType="today" />
      </div>
      <div className="Summary__Chart grid grid-cols-2">
        <SummaryOrderChart />
        <SummaryOrdererChart />
      </div>
    </div>
  );
}

export default DashBoard;
