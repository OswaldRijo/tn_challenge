import { useState } from 'react';

const usePaginator = ({ initialPage = 0, limit = 10 } = {}) => {
  const [paginator, setPaginator] = useState({
    page: initialPage,
    limit
  });

  const setPage = (page: number) => setPaginator((p) => ({ ...p, page }));
  const setLimit = (limit: number) => setPaginator((p) => ({ ...p, limit }));
  const nextPage = () => setPaginator((p) => ({ ...p, page: p.page + 1 }));
  const prevPage = () => setPaginator((p) => ({ ...p, page: p.page - 1 }));

  return {
    page: paginator.page,
    pagePlusOne: paginator.page + 1,
    limit: paginator.limit,
    setPage,
    setLimit,
    nextPage,
    prevPage
  };
};

export default usePaginator;
