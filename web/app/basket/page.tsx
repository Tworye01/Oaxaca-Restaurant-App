'use client';

import React, { useEffect, useState } from 'react';
import { NavigationBarComponent } from "@/components/nav-bar";
import { Item } from '@/lib/types/item';
import { LoadingComponent } from '@/components/loading';
import { deleteBasket, deleteItemFromBasket, getBasket } from '@/lib/utils/basket';
import { Basket, BasketItem } from '@/lib/types/basket';
import { BasketItemComponent } from '@/components/basket';
import axios from 'axios';
import Link from 'next/link';

type Response = {
  items: Item[];
  itemsCount: number;
  cost: number;
}

export default function Page() {
  const [basket, setBasket] = useState<Basket | undefined>();
  const [cost, setCost] = useState<number>(0);

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  const emptyBasket = () => {
    deleteBasket();
    fetchData();
  }

  const deleteItem = (id?: number) => {
    if (id != undefined) {
      deleteItemFromBasket(id)
    }

    fetchData();
  }

  const fetchResponse = async (basketData: Basket): Promise<Response | undefined> => {
    try {
      const response = await axios.post(`http://${apiAddr}/api/v1/baskets/items`, basketData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      console.log('Response:', response.data);
      return response.data;
    } catch (error) {
      console.error('Error:', error);
      return undefined;
    }
  }

  const fetchData = async () => {
    const basketData = getBasket();

    if (!basketData) {
      return;
    }

    console.log(basketData);
    const response = await fetchResponse(basketData);

    if (response === undefined) {
      return;
    }

    setCost(response.cost);

    for (let i = 0; i < basketData?.items.length; i++) {
      for (let j = 0; j < response.items.length; j++) {
        if (basketData.items[i].itemId != response.items[j].id) {
          continue;
        }

        basketData.items[i].item = response.items[j];
        break;
      }
    }

    setBasket(basketData);
  }

  useEffect(() => {
    fetchData();
  }, [])

  return (
    <main className="text-center min-h-screen pb-32 bg-warm-100 text-brown-dark-100">
      <NavigationBarComponent />

      <div className="min-h-screen">
        <div className="flex flex-col justify-center items-center pt-20">
          <h1 className="text-3xl font-bold mb-12">Basket</h1>

          {
            <div className="max-w-md w-full bg-gray-100 rounded-lg shadow-md p-6">
              {
                basket &&
                <div>
                  <h2 className="text-xl font-bold mb-6">Total: £{cost}</h2>
                  {
                    basket.items.map((item: BasketItem, index: number) => (
                      <div key={index} className="mb-4">
                        <BasketItemComponent item={item} onClick={deleteItem} />
                      </div>
                    ))
                  }
                </div>
              }

              <div className="flex justify-between mt-8">
                <button
                  className="px-4 py-2 text-white bg-red-500 hover:bg-red-600 shadow-lg"
                  onClick={() => emptyBasket()}
                >
                  Empty Basket
                </button>

                <Link legacyBehavior href="/menu">
                  <a className="px-4 py-2 text-white bg-pinky-500  hover:bg-pinky-600 shadow-lg">
                    Return to Menu
                  </a>
                </Link>

                <Link legacyBehavior href="/pay">
                  <a className="px-4 py-2 text-white bg-blue-500  hover:bg-blue-600 shadow-lg">
                    Pay
                  </a>
                </Link>
              </div>
            </div>
          }
        </div>
      </div>
    </main >
  );
};
