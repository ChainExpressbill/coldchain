import React from 'react';
import { useGetCharts } from 'data/dashboard/dashboard.hooks';
import { ChartType } from 'data/dashboard/dashboard.model';
import Chart from 'react-apexcharts';

interface ChartData {
  categories: string[];
  seriesData: number[];
}

const initialState = {
  options: {
    chart: {
      id: 'order-line-chart',
      zoom: {
        enabled: false,
      },
    },
    title: {
      text: '한 달간 주문 건수',
    },
    xaxis: {
      categories: [],
    },
  },
  series: [
    {
      name: '주문 건수',
      data: [],
    },
  ],
};

function SummaryOrderChart() {
  const { data } = useGetCharts(ChartType.ORDERS);
  const [chartState, setChartState] = React.useState(initialState);

  React.useEffect(() => {
    if (data) {
      const { categories, seriesData } = data?.reduce(
        (xy: ChartData, point) => {
          const { timestamp, count } = point;
          const yyyyMMdd = timestamp.split('T')[0];
          xy.categories.push(yyyyMMdd);
          xy.seriesData.push(count);
          return xy;
        },
        {
          categories: [],
          seriesData: [],
        },
      ) as any;

      const series = {
        name: '주문 건수',
        data: seriesData,
      };

      setChartState((prev) => {
        return {
          ...prev,
          options: {
            ...prev.options,
            xaxis: {
              ...prev.options.xaxis,
              categories,
            },
          },
          series: [series],
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
