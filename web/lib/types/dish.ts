import { Item } from "./item";

export enum Spice {
  Mild = 0,
  Medium,
  Hot,
  ExtraHot,
}

const spiceMap: Map<Spice, String> = new Map([
  [Spice.Mild, "mild"],
  [Spice.Medium, "medium"],
  [Spice.Hot, "hot"],
  [Spice.ExtraHot, "extra hot"],
]);

export function getSpiceString(spice: Spice): string {
  const spiceString = spiceMap.get(spice);
  if (spiceString === undefined) {
    throw new Error("Invalid spice value");
  }

  return spiceString.toString();
}

export enum Course {
  Starter = 0,
  Main,
  Dessert,
  Side,
  Extra,
}

const courseMap: Map<Course, String> = new Map([
  [Course.Starter, "starter"],
  [Course.Main, "main"],
  [Course.Dessert, "dessert"],
  [Course.Side, "side"],
  [Course.Extra, "extra"],
]);

export function getCourseString(course: Course): string {
  const courseString = courseMap.get(course);
  if (courseString === undefined) {
    throw new Error("Invalid spice value");
  }

  return courseString.toString();
}

export type Dish = {
  id: number;
  item: Item;
  course: Course;
  spice: Spice;
  vegetarian: boolean;
};

export const getBlankDish = (): Dish => {
  return {
    id: 0,
    item: {
      id: 0,
      name: "Dish Name",
      description: "Dish Description",
      cost: 0,
      calories: 0,
      allergens: [""],
      available: false,
      image: "http://localhost:8080/images/dishes/blank.jpg",
    },
    course: 0,
    spice: 0,
    vegetarian: false,
  };
};
