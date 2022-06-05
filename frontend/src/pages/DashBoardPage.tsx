import React from 'react';
import { SummaryBoard, SummaryLineChart } from 'components/dashboard';
import { ChartType } from 'data/dashboard/dashboard.model';

function DashBoard() {
  return (
    <div className="dashboard h-full grid grid__Two__Row">
      <div className="Summary__Board flex flex-row justify-evenly items-center">
        <SummaryBoard boardType="lastMonth" />
        <SummaryBoard boardType="today" />
      </div>
      <div className="Summary__Chart">
        <SummaryLineChart chartType={ChartType.ORDERS} />
      </div>
    </div>
  );
}

export default DashBoard;
