import { Item } from "./item";

export type BasketItem = {
  id?: number;
  itemId: number;
  item?: Item;
  note: string;
};

export type Basket = {
  items: BasketItem[];
  total: number;
};

export const getEmptyBasket = (): Basket => {
  return {
    items: [],
    total: 0,
  };
};
