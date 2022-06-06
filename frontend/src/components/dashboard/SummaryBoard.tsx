import React from 'react';
import { useGetLastMonth, useGetToday } from 'data/dashboard/dashboard.hooks';
import Skeleton from 'react-loading-skeleton';

type BoardType = 'lastMonth' | 'today';
interface SummaryBoardProps {
  boardType: BoardType;
}

function BoardBox({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col w-[250px] h-[200px] justify-center px-4 Rounded__Shadow__Border mr-5">
      {children}
    </div>
  );
}

function SummaryBoard({ boardType }: SummaryBoardProps) {
  const isLastMonth = boardType === 'lastMonth';
  const { data, isLoading } = isLastMonth ? useGetLastMonth() : useGetToday();
  const title = isLastMonth ? '지난 달 주문 요약' : '금일 주문 현황';

  if (isLoading) {
    return (
      <BoardBox>
        <Skeleton count={3} />
      </BoardBox>
    );
  }

  return (
    <BoardBox>
      <div className="text-3xl pb-4">{title}</div>
      <div className="text-lg">주문 요청 업체 수: {data?.ordererCount}</div>
      <div className="text-lg">주문 건수: {data?.orderCount}</div>
    </BoardBox>
  );
}

export default SummaryBoard;
