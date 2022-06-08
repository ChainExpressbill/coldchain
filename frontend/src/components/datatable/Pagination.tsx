import React from 'react';
import classnames from 'classnames';
import './pagination.scss';

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
    <div className="pagination">
      <button
        className="pagination__chevron"
        onClick={() => setPage(1)}
        disabled={currentPage === 1}
      >
        &lt;&lt;
      </button>
      <button
        className="pagination__chevron"
        onClick={() => setPage(currentPage - 1)}
        disabled={currentPage === 1}
      >
        &lt;
      </button>
      {pagesArray().map((page) => (
        <button
          key={page}
          onClick={() => setPage(page)}
          className={classnames('pagination__page', {
            pagination__current__page: page === currentPage,
          })}
        >
          {page}
        </button>
      ))}
      <button
        className="pagination__chevron"
        onClick={() => setPage(currentPage + 1)}
        disabled={currentPage === totalPages}
      >
        &gt;
      </button>
      <button
        className="pagination__chevron"
        onClick={() => setPage(totalPages)}
        disabled={currentPage === totalPages}
      >
        &gt;&gt;
      </button>
    </div>
  );
}

export default Pagination;
