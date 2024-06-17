'use client'

import { NavigationBarStaffComponent } from "@/components/nav-bar-staff";
import { OrderComponent, PaidButtonComponent } from "@/components/order";
import { ToastComponent } from "@/components/toast";
import { Notification } from "@/lib/types/notification";
import { Order } from "@/lib/types/order";
import { useEffect, useState } from "react";
import useWebSocket from "react-use-websocket";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useToken } from "@/lib/utils/token";
import { LoadingComponent } from "@/components/loading";
import Link from "next/link";

type Response = {
  orders: Order[];
  ordersCount: number;
}

type Data = {
  body: string;
}

export default function Staff(
  { params }: { params: { id: string } }
) {
  const [orders, setOrders] = useState<Response>();
  const [notifications, setNotifications] = useState<Notification[]>([]);
  const [loading, setLoading] = useState(true);
  const [token] = useToken();
  const router = useRouter();

  useWebSocket(`ws://localhost:8080/api/v1/staff/${params.id}`, {
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

  const fetchData = async () => {
    const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

    await axios.get(`http://${apiAddr}/api/v1/table/orders/${params.id}`, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    })
      .then(response => {
        console.log('Response:', response.data);
        setOrders(response.data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
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
    if (!loading) {
      fetchData();
    }
  }, [loading]);

  useEffect(() => {
    const intervalId = setInterval(() => {
      fetchData();
    }, 5000);

    return () => clearInterval(intervalId);
  }, []);

  if (loading) {
    return <LoadingComponent />;
  }

  return (
    <main className="w-full relative items-center text-center bg-warm-100 min-h-screen">
      <NavigationBarStaffComponent />
      <div className="flex flex-col justify-start items-start flex-grow pt-20 pl-2 overflow-y-auto">
        <div className="fixed top-20 right-0 mr-4 mt-4 z-50">
          <Link legacyBehavior href="/admin/order">
            <a className="px-4 py-2 text-black bg-pinky-500 hover:bg-pinky-600 shadow-lg">
              Order
            </a>
          </Link>
        </div>
        <div className="mx-auto w-full pl-2 fixed left-0 max-w-screen-md">
          {
            notifications.map((notification, index) => (
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
            ))
          }
        </div>
        <div className="mx-auto w-1/3">
          {
            orders && orders.orders.map((order: Order, index: number) => (
              <div key={index}>
                <OrderComponent order={order} showId={true}>
                  <div className="justify-center items-center mb-1">
                    <PaidButtonComponent order={order} />
                  </div>
                </OrderComponent>
              </div>
            ))
          }
        </div>
      </div>
    </main>
  );
}
