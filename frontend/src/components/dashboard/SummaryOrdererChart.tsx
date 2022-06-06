/* eslint-disable @typescript-eslint/no-unused-vars */
import React from 'react';
import { useGetCharts } from 'data/dashboard/dashboard.hooks';
import { ChartType } from 'data/dashboard/dashboard.model';
import Chart from 'react-apexcharts';

interface SeriesData {
  x: string; // yyyy-MM-dd
  y: number;
}
interface Series {
  name: string;
  data: SeriesData[];
}

const initialState: {
  options: any;
  series: Series[];
} = {
  options: {
    chart: {
      id: 'orderer-multi-axis-line-chart',
      zoom: {
        enabled: false,
      },
    },
    title: {
      text: '한 달간 주문 업체 수',
    },
  },
  series: [],
};

function SummaryOrderChart() {
  const { data } = useGetCharts(ChartType.ORDERERS);
  const [chartState, setChartState] = React.useState(initialState);

  React.useEffect(() => {
    if (data) {
      const initialSeries = data?.reduce(
        (seriesData: Series[], { orderer }) => {
          const seriesIndex = seriesData.findIndex(
            (s: Series) => s.name === orderer,
          );

          if (seriesIndex < 0) {
            seriesData.push({
              name: orderer ?? '',
              data: [],
            });
          }
          return seriesData;
        },
        [],
      );

      const series = initialSeries.reduce(
        (seriesData: Series[], current, index) => {
          for (const { orderer, timestamp, count } of data) {
            const yyyyMMdd = timestamp.split('T')[0];

            const duplicatedTimestampIndex = seriesData[index].data.findIndex(
              (xy) => xy.x === yyyyMMdd,
            );

            const y = current.name === orderer ? count : 0;
            if (duplicatedTimestampIndex >= 0) {
              seriesData[index].data[duplicatedTimestampIndex].y += y;
            } else {
              seriesData[index].data.push({
                x: yyyyMMdd,
                y,
              });
            }
          }
          return seriesData;
        },
        initialSeries,
      );

      setChartState((prev) => {
        return {
          ...prev,
          series,
        };
      });
    }
  }, [data]);

  return (
    <Chart
      options={chartState.options}
      series={chartState.series}
      type="line"
      width="800"
    />
  );
}

export default SummaryOrderChart;
