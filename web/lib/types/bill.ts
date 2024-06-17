export type Bill = {
  id: number;
  table: number;
  cost: number;
  paid: boolean;
};

export type BillsResponse = {
  bills: Bill[];
  billsCount: number;
};
