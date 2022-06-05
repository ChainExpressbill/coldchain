import React from 'react';
import { useGetCharts } from 'data/dashboard/dashboard.hooks';
import { ChartType } from 'data/dashboard/dashboard.model';
import Chart from 'react-apexcharts';

// interface ChartData {
//   categories: string[];
//   seriesData: number[];
// }

const initialState = {
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
    xaxis: {
      categories: [2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016],
    },
    yaxis: [
      {
        axisTicks: {
          show: true,
        },
        axisBorder: {
          show: true,
          color: '#FF1654',
        },
        labels: {
          style: {
            colors: '#FF1654',
          },
        },
        title: {
          text: 'Series A',
          style: {
            color: '#FF1654',
          },
        },
      },
      {
        opposite: true,
        axisTicks: {
          show: true,
        },
        axisBorder: {
          show: true,
          color: '#247BA0',
        },
        labels: {
          style: {
            colors: '#247BA0',
          },
        },
        title: {
          text: 'Series B',
          style: {
            color: '#247BA0',
          },
        },
      },
    ],
  },
  series: [
    {
      name: 'Series A',
      data: [1.4, 2, 2.5, 1.5, 2.5, 2.8, 3.8, 4.6],
    },
    {
      name: 'Series B',
      data: [20, 29, 37, 36, 44, 45, 50, 58],
    },
  ],
};

function SummaryOrderChart() {
  const { data } = useGetCharts(ChartType.ORDERERS);
  const [chartState] = React.useState(initialState);
  // const [chartState, setChartState] = React.useState(initialState);

  React.useEffect(() => {
    console.log(data);
    // if (data) {
    //   const { categories, seriesData } = data?.reduce(
    //     (xy: ChartData, point) => {
    //       const { timestamp, count } = point;
    //       const yyyyMMdd = timestamp.split('T')[0];
    //       xy.categories.push(yyyyMMdd);
    //       xy.seriesData.push(count);
    //       return xy;
    //     },
    //     {
    //       categories: [],
    //       seriesData: [],
    //     },
    //   ) as any;

    //   const series = {
    //     name: '주문 업체 수',
    //     data: seriesData,
    //   };

    //   setChartState((prev) => {
    //     return {
    //       ...prev,
    //       options: {
    //         ...prev.options,
    //         xaxis: {
    //           ...prev.options.xaxis,
    //           categories,
    //         },
    //       },
    //       series: [series],
    //     };
    //   });
    // }
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
