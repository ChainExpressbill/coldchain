import React from 'react';

type PaginationProps = {
  totalPages: number;
  displayAmount: number;
  currentPage: number;
  setPage: (page: number) => void;
};

function Pagination({
  totalPages,
  displayAmount,
  currentPage,
  setPage,
}: PaginationProps) {
  const pagesStart =
    Math.floor((currentPage - 1) / displayAmount) * displayAmount + 1;

  const pagesArray = () => {
    if (
      Math.floor(totalPages / displayAmount) ===
      Math.floor((currentPage - 1) / displayAmount)
    ) {
      const lastArraySize = Math.floor(totalPages % displayAmount);
      return [
        ...Array(lastArraySize)
          .fill(0)
          .map((_, i) => pagesStart + i),
      ];
    } else {
      return [
        ...Array(displayAmount)
          .fill(0)
          .map((_, i) => pagesStart + i),
      ];
    }
  };

  return (
    <div>
      <button onClick={() => setPage(1)} disabled={currentPage === 1}>
        &lt;&lt;
      </button>
      <button
        onClick={() => setPage(currentPage - 1)}
        disabled={currentPage === 1}
      >
        &lt;
      </button>
      {pagesArray().map((page) => (
        <button
          key={page}
          onClick={() => setPage(page)}
          className={page === currentPage ? 'current-page' : ''}
        >
          {page}
        </button>
      ))}
      <button
        onClick={() => setPage(currentPage + 1)}
        disabled={currentPage === totalPages}
      >
        &gt;
      </button>
      <button
        onClick={() => setPage(totalPages)}
        disabled={currentPage === totalPages}
      >
        &gt;&gt;
      </button>
    </div>
  );
}

export default Pagination;
