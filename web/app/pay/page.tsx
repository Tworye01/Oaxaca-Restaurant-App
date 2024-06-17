'use client'

import { LoadingComponent } from '@/components/loading';
import { PaymentItemComponent, PayModalComponent } from '@/components/payment';
import { Basket, BasketItem } from '@/lib/types/basket';
import { Item } from '@/lib/types/item';
import { deleteBasket, getBasket } from '@/lib/utils/basket';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import React, { useEffect, useState } from 'react';

type Response = {
  items: Item[];
  itemsCount: number;
  cost: number;
}

export default function Pay() {
  const [payModalVisible, setPayModalVisible] = useState<boolean>(false);
  const [basket, setBasket] = useState<Basket | undefined>();
  const [table, setTable] = useState<number>();
  const [paid, setPaid] = useState<boolean>();
  const router = useRouter();

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

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

    setBasket(basketData);
  }

  const postOrder = async () => {
    await axios.post(`http://${apiAddr}/api/v1/baskets/pay/${table}`, basket, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(response => {
        console.log('Response:', response.data);
      })
      .catch(error => {
        console.error('Error:', error);
      });

    deleteBasket();
    setPaid(true)
  }

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        setPayModalVisible(false);
      }
    };

    if (payModalVisible) {
      document.addEventListener('keydown', handleKeyDown);
    } else {
      document.removeEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [payModalVisible]);

  useEffect(() => {
    if (paid) {
      router.push(`/table/${table}`);
    }
    fetchData();
  }, [table, router, paid])


  if (!basket) return <LoadingComponent />

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
                    onClick={() => setPayModalVisible(true)}
                  >
                    Pay
                  </button>
                </div>
              </div>


              {
                payModalVisible &&
                <center>
                  <div className="fixed inset-0 flex justify-center items-center m-auto">
                    <div className="overflow-y-auto overflow-x-hidden fixed backdrop-brightness-50 z-50 w-full md:inset-0 max-h-full">
                      <div className="shadow-lg mt-24 max-w-md bg-white">
                        <PayModalComponent
                          cost={basket.total}
                          onSubmit={() => postOrder()}
                        />
                      </div>
                    </div>
                  </div>
                </center>
              }
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
