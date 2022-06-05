import React from 'react';
import { useGetLastMonth, useGetToday } from 'data/dashboard/dashboard.hooks';
import Skeleton from 'react-loading-skeleton';

type BoardType = 'lastMonth' | 'today';
interface SummaryBoardProps {
  boardType: BoardType;
}

function SummaryBoard({ boardType }: SummaryBoardProps) {
  const isLastMonth = boardType === 'lastMonth';
  const { data, isLoading } = isLastMonth ? useGetLastMonth() : useGetToday();
  const title = isLastMonth ? '지난 달 주문 요약' : '금일 주문 현황';

  if (isLoading) {
    return <Skeleton count={3} />;
  }

  return (
    <div>
      <div>{title}</div>
      <div>주문 요청 업체 수: {data?.ordererCount}</div>
      <div>주문 건수: {data?.orderCount}</div>
    </div>
  );
}

export default SummaryBoard;
