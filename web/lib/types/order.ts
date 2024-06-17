import { Bill } from "./bill";
import { Item } from "./item";

export enum Status {
  Unknown = 0,
  Preperation,
  Cooking,
  Waiting,
  Serving,
  Done,
  Cancelled,
}

const statusMap: Map<Status, String> = new Map([
  [Status.Unknown, "unknown"],
  [Status.Preperation, "preparing"],
  [Status.Cooking, "cooking"],
  [Status.Waiting, "waiting"],
  [Status.Serving, "serving"],
  [Status.Done, "completed"],
  [Status.Cancelled, "cancelled"],
]);

export function getStatusString(status: Status): string {
  const courseString = statusMap.get(status);
  if (courseString === undefined) {
    throw new Error("Invalid status value");
  }

  return courseString.toString();
}

export type Order = {
  id: number;
  createdAt: string;
  table: number;
  status: Status;
  cost: number;
  item: Item;
  bill: Bill;
  note: string;
};

export type OrdersResponse = {
  orders: Order[];
  ordersCount: number;
};
