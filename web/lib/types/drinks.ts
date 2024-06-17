import { Item } from "./item";

export type Drink = {
  id: number;
  item: Item;
  alcoholic: boolean;
};

export const getBlankDrink = (): Drink => {
  return {
    id: 0,
    item: {
      id: 0,
      name: "Drink Name",
      description: "Drink Description",
      cost: 0,
      calories: 0,
      allergens: [""],
      available: false,
      image: "http://localhost:8080/images/drinks/blank.jpg",
    },
    alcoholic: false,
  };
};
