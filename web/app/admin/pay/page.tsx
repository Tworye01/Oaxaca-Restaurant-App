'use client'

import { LoadingComponent } from '@/components/loading';
import { PaymentItemComponent, PayModalComponent } from '@/components/payment';
import { Basket, BasketItem } from '@/lib/types/basket';
import { Item } from '@/lib/types/item';
import { deleteBasket, getBasket } from '@/lib/utils/basket';
import { deleteStaffBasket, getStaffBasket, setStaffBasket } from '@/lib/utils/staff-basket';
import { useToken } from '@/lib/utils/token';
import axios from 'axios';
import { useRouter } from "next/navigation";
import React, { useEffect, useState } from 'react';

type Response = {
  items: Item[];
  itemsCount: number;
  cost: number;
}

export default function Pay() {
  const [order, setOrder] = useState<boolean>(false);
  const [basket, setStaffBasket] = useState<Basket | undefined>();
  const [table, setTable] = useState<number>();
  const [loading, setLoading] = useState(true);
  const [token] = useToken();
  const router = useRouter();

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  const fetchResponse = async (basketData: Basket): Promise<Response | undefined> => {
    try {
      const response = await axios.post(`http://${apiAddr}/api/v1/baskets/items`, basketData, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
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
    const basketData = getStaffBasket();

    if (!basketData) {
      return;
    }

    const response = await fetchResponse(basketData);

    if (response === undefined) {
      return;
    }

    basketData.total = response.cost;

    for (let i = 0; i < basketData?.items.length; i++) {
      for (let j = 0; j < response.items.length; j++) {
        if (basketData.items[i].itemId != response.items[j].id) {
          continue;
        }

        basketData.items[i].item = response.items[j];
        break;
      }
    }

    setStaffBasket(basketData);
  }

  const postOrder = async () => {
    await axios.post(`http://${apiAddr}/api/v1/baskets/items/${table}`, basket, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    })
      .then(response => {
        console.log('Response:', response.data);
      })
      .catch(error => {
        console.error('Error:', error);
      });

    deleteStaffBasket();
    setOrder(true);
  }

  useEffect(() => {
    if (!token) {
      router.push("/login");
    } else {
      setLoading(false);
      fetchData();
    }
  }, [token, router]);

  useEffect(() => {
    if (order) {
      router.push(`/admin/waiter/${table}`);
    }
    fetchData();
    console.log(table)
  }, [table, order, router])


  if (!basket || loading) return <LoadingComponent />

  return (
    <main>
      <div className="mt-4">

        {
          (basket.items.length > 0) ? (
            <>
              <div className="lg:w-1/3 flex-auto mx-auto items-center justify-center shadow-lg mb-8">
                <h1 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg mb-4">Basket</h1>
                {
                  basket &&
                  basket.items.map((item: BasketItem, index: number) => (
                    <div key={index}>
                      {
                        item.item &&
                        <PaymentItemComponent
                          item={item.item}
                          note={item.note}
                        />
                      }
                    </div>
                  ))
                }
                <h1 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg mt-4">Total £{basket.total}</h1>
              </div>

              <div className="lg:w-1/3 flex-auto mx-auto items-center justify-center">
                <div className="mt-4 max-w-s text-lg">
                  <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Table Number</h2>
                  <input
                    className="mt-2"
                    type="number"
                    placeholder="e.g. 12"
                    onChange={e => setTable(parseInt(e.target.value))}
                  />
                </div>

                <div className="mt-4 max-w-s text-lg items-center justify-center flex-auto">
                  <button
                    className="text-center font-semibold text-gray-900 bg-pinky-300 hover:bg-pinky-600 active:bg-pinky-600 disabled:bg-slate-300 p-4 px-8 shadow-lg mb-4"
                    disabled={table === undefined || table.toString() === "NaN"}
                    onClick={() => postOrder()}
                  >
                    Send Order
                  </button>
                </div>
              </div>
            </>
          ) : (
            <>
              <div className="lg:w-1/3 flex-auto mx-auto items-center justify-center shadow-lg mb-8">
                <h1 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg mb-4">Basket</h1>
                <p className="text-left p-4">No items in basket to be bought.</p>
              </div>
            </>
          )
        }

      </div>
    </main>
  );
};
