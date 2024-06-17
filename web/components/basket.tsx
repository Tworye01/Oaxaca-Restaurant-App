/* eslint-disable @next/next/no-img-element */

import
React, {
  useState
} from "react";
import Link from 'next/link';
import { BasketItem } from "@/lib/types/basket";
import { Item } from "@/lib/types/item";
import { updateItemInBasket } from "@/lib/utils/basket";
import { updateItemInStaffBasket } from "@/lib/utils/staff-basket";

export const BasketIcon = () => {
  return (
    <Link href="/basket" className='bg-pinky-600 hover:bg-slate-100 hover:text-slate-900 w-16 px-4 py-4 flex mt-24 fixed shadow-lg mr-4 right-64' >
      <img className="" src="/Basket.png" alt="basket-icon" />
    </Link>
  );
};

export const BasketIconWaiter = () => {
  return (
    <Link href="/admin/basket" className='bg-pinky-600 hover:bg-slate-100 hover:text-slate-900 w-16 px-4 py-4 flex mt-24 fixed shadow-lg mr-4 right-64' >
      <img className="" src="/Basket.png" alt="basket-icon" />
    </Link>
  );
};

export const BasketItemComponent = (props: {
  item: BasketItem,
  onClick: (id?: number) => void,
}): JSX.Element => {
  const [note, setNote] = useState("");

  const setNoteHandle = (value: string): void => {
    // Update the local state
    setNote(value);

    // Prepare the updated item object with the new note value
    const updatedItem: BasketItem = { ...props.item, note: value };

    // Update the item in the basket (assuming this function updates the item in an external data source)
    updateItemInBasket(updatedItem);
  }

  return (
    <div className="flex items-center justify-between p-4 border bg-slate-300 shadow-lg">
      <div className="flex flex-col flex-1">
        <div className="px-2">
          <input
            type="text"
            className="w-full px-4 py-1 border border-gray-300 focus:outline-none focus:border-indigo-500"
            placeholder="Add a note..."
            value={note}
            onChange={(e) => setNoteHandle(e.target.value)}
          />
        </div>
        {
          props.item.item &&
          <BasketMenuItemComponent item={props.item.item} />
        }
      </div>
      <button
        className=" mt-8 px-4 py-2 bg-red-500 text-black hover:bg-red-700  shadow-lg"
        onClick={() => props.onClick(props.item.id)}
      >
        Remove
      </button>
    </div>
  );
};

export const BasketMenuItemComponent = (props: {
  item: Item,
}): JSX.Element => {
  return (
    <div className="mt-2 flex items-center">
      <img
        className="ml-2 w-24 h-24 object-cover"
        alt={`${props.item.name} image`}
        src={props.item.image}
      />
      <div className="flex-1 text-left pl-2">
        <p className="text-gray-700 text-lg">Name: {props.item.name}</p>
        <p className="text-gray-700 text-lg">Cost: £{props.item.cost.toFixed(2)}</p>
        <p className="text-gray-700 text-lg" >Calories: {props.item.calories}</p>
      </div>
    </div>
  );
};

export const BasketItemStaffComponent = (props: {
  item: BasketItem,
  onClick: (id?: number) => void,
}): JSX.Element => {
  const [note, setNote] = useState("");

  const setNoteHandle = (value: string): void => {
    // Update the local state
    setNote(value);

    // Prepare the updated item object with the new note value
    const updatedItem: BasketItem = { ...props.item, note: value };

    // Update the item in the basket (assuming this function updates the item in an external data source)
    updateItemInStaffBasket(updatedItem);
  }

  return (
    <div className="flex items-center justify-between p-4 border bg-slate-300 shadow-lg">
      <div className="flex flex-col flex-1">
        <div className="px-2">
          <input
            type="text"
            className="w-full px-4 py-1 border border-gray-300 focus:outline-none focus:border-indigo-500"
            placeholder="Add a note..."
            value={note}
            onChange={(e) => setNoteHandle(e.target.value)}
          />
        </div>
        {
          props.item.item &&
          <BasketMenuItemComponent item={props.item.item} />
        }
      </div>
      <button
        className=" mt-8 px-4 py-2 bg-red-500 text-black hover:bg-red-700  shadow-lg"
        onClick={() => props.onClick(props.item.id)}
      >
        Remove
      </button>
    </div>
  );
};
