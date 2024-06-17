'use client'

import React, { useEffect, useState } from 'react';
import { v4 as uuidv4 } from 'uuid';
import { NavigationBarComponent } from '@/components/nav-bar';
import { Order, OrdersResponse } from '@/lib/types/order';
import axios from 'axios';
import { LoadingComponent } from '@/components/loading';
import useWebSocket from 'react-use-websocket';
import { OrderComponent } from '@/components/order';
import { AccordionComponent } from '@/components/accordion';
import { Bill, BillsResponse } from '@/lib/types/bill';
import { BillComponent } from '@/components/bills';
import { Notification } from "@/lib/types/notification";
import { ToastComponent } from '@/components/toast';


enum Tab {
  orders = 0,
  bills,
}

type Data = {
  body: string;
}

export default function Table(
  { params }: { params: { id: string } }
) {
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [orders, setOrders] = useState<OrdersResponse>();
  const [bills, setBills] = useState<BillsResponse>();
  const [tab, setTab] = useState<Tab>(Tab.orders);

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  const { sendJsonMessage } = useWebSocket(`ws://${apiAddr}/api/v1/order/${params.id}`, {
    onOpen: () => {
      console.log("connected");
    },
    onMessage: (event: WebSocketEventMap["message"]) => {
      const data: Data = JSON.parse(event.data);
      const newNotification: Notification = { text: data.body };
      setNotifications((prevNotifications) => {
        if (
          prevNotifications.some(
            (notification) => notification.text === newNotification.text
          )
        ) {
          return prevNotifications;
        }
        return [...prevNotifications, newNotification];
      });
      setTimeout(() => {
        setNotifications((prevNotifications) =>
          prevNotifications.filter(
            (notification) => notification.text !== newNotification.text
          )
        );
      }, 10000);
    },
  });
  

  const sendHelpMessage = async () => {
    sendJsonMessage({ type: "PING", body: "ALERT" });
  }

  const fetchData = async () => {
    await axios.get(`http://${apiAddr}/api/v1/table/orders/${params.id}`, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(response => {
        console.log('Response:', response.data);
        setOrders(response.data)
        console.log(orders)
      })
      .catch(error => {
        console.error('Error:', error);
      });

    await axios.get(`http://${apiAddr}/api/v1/table/bills/${params.id}`, {
      headers: {
        'Content-Type': 'application/json'
      }
    })
      .then(response => {
        console.log('Response:', response.data);
        setBills(response.data)
        console.log(orders)
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }

  useEffect(() => {
    fetchData();
  }, [tab])

  useEffect(() => {
    if (orders) {
      sendOrderMessage(orders.orders);
    }
  }, [orders]);

  const sendOrderMessage = async (ordersArray: Order[]) => {
    ordersArray.forEach(order => {
      sendJsonMessage({ type: "ORDER", body: String(order.id) });
    });
  }

  if (!orders || !bills) return <LoadingComponent />

  return (
    <main className="w-full relative text-center bg-warm-100">
      <NavigationBarComponent />

      <div className="min-h-screen">
        <div className="pt-20 pl-2 fixed left-0">
          {notifications.map((notification, index) => (
            <ToastComponent
              key={index}
              style="info"
              onClick={() => console.log("clicked")}
              onClose={() =>
                setNotifications((prevNotifications) =>
                  prevNotifications.filter((_, i) => i !== index)
                )
              }
              text={notification.text}
            />
          ))}
        </div>
        <div className="flex flex-col justify-center items-center pt-20">
          <h1 className="mb-12">Table {parseInt(params.id)}</h1>

          <div className="mb-4">
            <button
              className="w-32 text-center font-semibold text-gray-900 hover:bg-gray-100 active:bg-gray-200 disabled:hover:bg-pinky-100 disabled:active:bg-pinky-200 p-4 px-8 shadow-lg mb-4 border-b-4 border-gray-300 disabled:border-pinky-300"
              disabled={tab === Tab.orders}
              onClick={() => setTab(Tab.orders)}
            >
              Orders
            </button>
            <button
              className="w-32 text-center font-semibold text-gray-900 hover:bg-gray-100 active:bg-gray-200 disabled:hover:bg-pinky-100 disabled:active:bg-pinky-200 p-4 px-8 shadow-lg mb-4 border-b-4 border-gray-300 disabled:border-pinky-300"
              disabled={tab === Tab.bills}
              onClick={() => setTab(Tab.bills)}
            >
              Bills
            </button>
          </div>

          {
            tab === Tab.orders &&
            <div className="flex flex-col justify-center items-center lg:w-1/3 sm:w-3/4">
              {
                orders.orders.map((order: Order, index: number) => (
                  <div key={index} className="mb-4 w-full max-w-md">
                    <AccordionComponent
                      preview={
                        <h1 className="text-left ml-4">{`Order ${order.id} - Bill ${order.bill.id} - ${order.item.name}`}</h1>
                      }
                      contents={
                        <OrderComponent order={order} />
                      }
                    />
                  </div>
                ))
              }
            </div>
          }

          {
            tab === Tab.bills &&
            <div className="flex flex-col justify-center items-center lg:w-1/3 sm:w-3/4">
              {
                bills.bills.map((bill: Bill, index: number) => (
                  <div key={index} className="mb-4 w-full max-w-md">
                    <BillComponent bill={bill} />
                  </div>
                ))
              }
            </div>
          }

          <div className="mt-12">
            <button
              className="text-center font-semibold text-gray-900 bg-pinky-600 hover:bg-pinky-700 active:bg-pinky-700  p-4 px-8 shadow-lg mb-4"
              onClick={() => sendHelpMessage()}
            >
              Get help!
            </button>
          </div>
        </div>
      </div>
    </main >
  );
};
