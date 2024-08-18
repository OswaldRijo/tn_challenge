import { useEffect } from 'react';
import * as React from 'react';
import { useTranslation } from 'react-i18next';
import DeleteIcon from '@mui/icons-material/Delete';
import { IconButton, Table, TableBody, TableCell, TableContainer, TableHead, TablePagination, TableRow, Tooltip } from '@mui/material';

import TableSkeleton from '@/components/operations/TableSkeleton';
import { dispatch, useSelector } from '@/store';
import { deleteOperation, filterOperations } from '@/store/slices/operation';
import { openSnackbar } from '@/store/slices/snackbar';
import { OperatorTypeToSignMap } from '@/types/operation';

export default function OperationsTable() {
  const { t } = useTranslation(['client', 'server']);
  const columns = [
    { field: 'id', headerName: t('table.header.id'), width: 70, valueGetter: (row) => row.id },
    {
      field: 'createdAt',
      headerName: t('table.header.createdAt'),
      sortable: false,
      width: 100,
      valueGetter: (row) => row.createdAt.toLocaleDateString()
    },
    {
      field: 'cost',
      headerName: t('table.header.cost'),
      sortable: false,
      width: 70,
      valueGetter: (row) => `$ ${row.operation.cost.toLocaleString()}`
    },
    {
      field: 'userBalance',
      headerName: t('table.header.userBalance'),
      width: 120,
      valueGetter: (row) => `$ ${row.userBalance.toLocaleString()}`
    },
    {
      field: 'args',
      headerName: t('table.header.args'),
      sortable: false,
      width: 80,
      valueGetter: (row) => {
        const sep = OperatorTypeToSignMap[row.operation.operationType];
        if (row.operation.args.length > 1) {
          return (row.operation.args as number[]).join(` ${sep} `);
        } else if (row.operation.args.length === 1) {
          return `${sep} ${(row.operation.args as number[])[0]}`;
        }
        return sep;
      }
    },
    {
      field: 'operationResponse',
      headerName: t('table.header.operationResponse'),
      type: 'string',
      sortable: false,
      width: 300,
      valueGetter: (row) => row.operationResponse
    },
    {
      field: 'delete'
    }
  ];
  const operationsSelector = useSelector((s) => s.operations);
  const [page, setPage] = React.useState<number>(0);
  const [limit, setLimit] = React.useState<number>(10);

  const filterUserOperations = (page: number, limit: number) => {
    dispatch(filterOperations({ page, limit })).catch((e) => {
      dispatch(
        openSnackbar({
          open: true,
          message: t(e.message),
          anchorOrigin: { vertical: 'top', horizontal: 'right' },
          variant: 'alert',
          closeColor: 'white',
          alert: {
            color: 'error'
          },
          close: true,
          transition: 'SlideLeft'
        })
      );
    });
  };

  const deleteUserOperation = (id: number) => {
    dispatch(deleteOperation(id))
      .then(() => filterUserOperations(page, limit))
      .catch((e) => {
        dispatch(
          openSnackbar({
            open: true,
            message: t(e.message),
            anchorOrigin: { vertical: 'top', horizontal: 'right' },
            variant: 'alert',
            closeColor: 'white',
            alert: {
              color: 'error'
            },
            close: true,
            transition: 'SlideLeft'
          })
        );
      });
  };

  useEffect(() => {
    filterUserOperations(page, limit);
  }, []);

  function handleChangePage(e, p: number) {
    setPage(p);
    filterUserOperations(p, limit);
  }

  function handleChangeRowsPerPage(r) {
    setLimit(r);
    filterUserOperations(page, r);
  }

  function handleDeleteRow(id: number) {
    deleteUserOperation(id);
  }

  if (operationsSelector.isFetching) return <TableSkeleton />;

  return (
    <div style={{ width: '100%' }}>
      <TableContainer sx={{ minWidth: 700 }}>
        <Table stickyHeader aria-label="sticky table">
          <TableHead>
            <TableRow>
              {columns.map((column, i) => {
                return <TableCell key={i}>{column.headerName}</TableCell>;
              })}
            </TableRow>
          </TableHead>
          <TableBody>
            {operationsSelector.records?.map((row, i) => {
              return (
                <TableRow hover tabIndex={-1} key={row.id}>
                  {columns.map((column) => {
                    if (column.field !== 'delete') {
                      const value = column.valueGetter(row);
                      return <TableCell key={column.field}>{value}</TableCell>;
                    } else {
                      return (
                        <TableCell key={row.id}>
                          <Tooltip title={t('table.header.delete')}>
                            <IconButton onClick={() => handleDeleteRow(row.id)}>
                              <DeleteIcon />
                            </IconButton>
                          </Tooltip>
                        </TableCell>
                      );
                    }
                  })}
                </TableRow>
              );
            })}
          </TableBody>
        </Table>
      </TableContainer>
      <TablePagination
        rowsPerPageOptions={[10, 25, 100]}
        component="div"
        count={operationsSelector.totalRecords || 0}
        rowsPerPage={limit}
        page={page}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
      />
    </div>
  );
}
