import { Dish } from "../types/dish";
import { Drink } from "../types/drinks";

export function isDish(item: Dish | Drink): item is Dish {
  return (item as Dish).course !== undefined;
}

export function isDrink(item: Dish | Drink): item is Drink {
  return (item as Drink).alcoholic !== undefined;
}
