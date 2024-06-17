import { Dish } from "./dish";
import { Drink } from "./drinks";

export type Menu = {
  dishes: Dish[];
  drinks: Drink[];
};
