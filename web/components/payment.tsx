/* eslint-disable @next/next/no-img-element */

import { Card, getBlankCard } from "@/lib/types/card";
import { Item } from "@/lib/types/item";
import { validateCard } from "@/lib/utils/card";
import { ReactNode, useEffect, useState } from "react";

export const PayModalComponent = (props: {
  cost: number,
  onSubmit: () => void,
  children?: ReactNode,
}): JSX.Element => {
  const [valid, setValid] = useState<boolean>(false);
  const [card, setCard] = useState<Card>(getBlankCard());

  useEffect(() => {
    if (card === undefined) {
      setValid(false);
      return;
    }

    setValid(validateCard(card));
  }, [card])

  return (
    <div className="p-4">
      <h1 className="text-left font-semibold text-gray-900 bg-pinky-300 p-4 shadow-lg">Total: £{props.cost}</h1>

      <div className="mt-4 max-w-s text-lg">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Card Number</h2>
        <input
          className="mt-4"
          type="number"
          placeholder="e.g. 1234567812345678"
          onChange={e => {
            const newCard = { ...card };
            newCard.cardNumber = e.target.value
            setCard(newCard);
          }}
        />
      </div>

      <div className="mt-4 max-w-s text-lg">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Security Number</h2>
        <input
          className="mt-4"
          type="number"
          placeholder="e.g. 123"
          onChange={e => {
            const newCard = { ...card };
            newCard.securityNumber = e.target.value
            setCard(newCard);
          }}
        />
      </div>

      <div className="mt-4 max-w-s text-lg">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Full Name</h2>
        <input
          className="mt-4"
          type="text"
          placeholder="Full Name"
          onChange={e => {
            const newCard = { ...card };
            newCard.cardName = e.target.value
            setCard(newCard);
          }}
        />
      </div>

      <div className="mt-4 max-w-s text-lg">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Expiration Date</h2>
        <input
          className="mt-4"
          type="month"
          placeholder="month year"
          onChange={e => {
            const newCard = { ...card };
            const newDate = e.target.valueAsDate;
            if (newDate == null) {
              return;
            }

            newCard.expirationDate = newDate;
            setCard(newCard);
          }}
        />
      </div>

      <div className="mt-2">
        {
          valid === true ? (
            <>
              <p className="text-green-600">Your details are valid!</p>
            </>
          ) : (
            <>
              <p className="text-red-600">Current details are invalid</p>
            </>
          )
        }
      </div>

      <button
        className="text-center mt-2 font-semibold text-gray-900 bg-pinky-300 hover:bg-pinky-600 active:bg-pinky-600 disabled:bg-slate-300 p-4 px-8 shadow-lg mb-4"
        disabled={!valid}
        onClick={() => props.onSubmit()}
      >
        Pay
      </button>
    </div>
  );
};

export const PaymentItemComponent = (props: {
  item: Item,
  note: string,
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
        <p className="text-gray-700 text-lg" >Note: {props.note}</p>
      </div>
    </div>
  );
};
