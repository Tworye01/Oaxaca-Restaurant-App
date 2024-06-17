import { Basket, BasketItem, getEmptyBasket } from "@/lib/types/basket";
import { useEffect } from "react";

// Gets the Basket from local storage, returns a new one if it is undefined.
export function getBasket(): Basket | undefined {
  if (typeof window !== "undefined" && window.localStorage) {
    const basketString = localStorage.getItem("basket");

    if (basketString != undefined) {
      const basket: Basket = JSON.parse(basketString);
      return basket;
    }

    return getEmptyBasket();
  }

  return undefined;
}

// Sets the Basket in local storage.
export function setBasket(basket: Basket) {
  if (typeof window !== "undefined" && window.localStorage) {
    localStorage.setItem("basket", JSON.stringify(basket));
  }
}

// Deletes the entire Basket.
export function deleteBasket() {
  if (typeof window !== "undefined" && window.localStorage) {
    localStorage.removeItem("basket");
  }
}

// Saves a new item to the local storage basket, if it as no id it will generate a new one.
export function saveToBasket(item: BasketItem) {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  if (item.id === undefined) {
    let id =
      basket.items.length != 0 ? basket.items[basket.items.length - 1].id : 0;
    if (id === undefined) {
      id = basket.items.length;
    }

    item.id = id + 1;
  }

  basket.items.push(item);
  setBasket(basket);
}

// If an item exists within the local storage basket it will be updated.
export function updateItemInBasket(item: BasketItem) {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  for (let i: number = 0; i < basket.items.length; i++) {
    if (basket.items[i].id === item.id) {
      basket.items[i] = item;
    }
  }

  setBasket(basket);
}

// Get an item at the given index of the basket in local storage, if the item does not exist it will return undefined.
export function getItemFromBasket(id: number): BasketItem | undefined {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  for (let i: number = 0; i < basket.items.length - 1; i++) {
    if (basket.items[i].id === id) {
      return basket.items[i];
    }
  }

  return undefined;
}

// Get an item at the given index of the basket in local storage, if the item does not exist or it is out of bounds it will return undefined.
export function getIndexFromBasket(index: number): BasketItem | undefined {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  if (index > basket.items.length) {
    return undefined;
  }

  const item: BasketItem = basket.items[index];
  return item;
}

// Delete an item with the given id from the basket in local storage.
export function deleteItemFromBasket(id: number) {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  basket.items = basket.items.filter((item) => item.id !== id);
  setBasket(basket);
}

// Delete an item at the given index from the basket in local storage.
export function deleteIndexFromBasket(index: number) {
  const basket = getBasket();
  if (basket === undefined) {
    return;
  }

  basket.items.splice(index, 1);
  setBasket(basket);
}
