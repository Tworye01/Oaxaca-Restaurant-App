'use client'

import { useState, useEffect } from 'react';
import axios from 'axios';
import { Order } from '@/lib/types/order';
import { OrderKitchenComponent } from '@/components/order'; // Assuming correct path to OrderKitchenComponent
import { LoadingComponent } from '@/components/loading';
import { Clock } from '@/components/clock';
import { NavigationBarStaffComponent } from '@/components/nav-bar-staff';
import { useToken } from '@/lib/utils/token';
import { useRouter } from 'next/navigation';

type Response = {
  orders: Order[];
  ordersCount: number;
};

export default function Kitchen() {
  const [orders, setOrders] = useState<Response | null>(null);
  const [loading, setLoading] = useState(true);
  const [token] = useToken();
  const router = useRouter();

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR || '';

  const fetchData = async () => {
    try {
      const response = await axios.get(`http://${apiAddr}/api/v1/orders`, {
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`
        }
      });
      console.log('Response:', response.data);
      setOrders(response.data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  useEffect(() => {
    if (!token) {
      router.push('/login');
    } else {
      setLoading(false);
      fetchData();
    }
  }, [token, router]);


  useEffect(() => {
    const intervalId = setInterval(() => {
      fetchData();
    }, 5000);

    return () => clearInterval(intervalId);
  }, []);

  const sortByCreatedAt = (a: Order, b: Order): number => {
    const dateA = new Date(a.createdAt).getTime();
    const dateB = new Date(b.createdAt).getTime();
    return dateA - dateB;
  };

  if (loading) {
    return <LoadingComponent />;
  }

  const handleSetStatus = (orderId: number, newStatus: number) => {
    setOrders(prevState => {
      if (!prevState) return null;
      const updatedOrders = prevState.orders.map(order => {
        if (order.id === orderId) {
          return {
            ...order,
            status: newStatus
          };
        }
        return order;
      });
      return {
        ...prevState,
        orders: updatedOrders
      };
    });
  };

  return (
    <main className="min-h-screen bg-warm-100">
      <NavigationBarStaffComponent />
      <div>
        <Clock />
      </div>

      <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
        {orders &&
          orders.orders.slice().sort(sortByCreatedAt).map((order, index) => (
            <div key={index}>
              <OrderKitchenComponent
                key={index}
                order={order}
                status={order.status}
                setStatus={(status) => handleSetStatus(order.id, status)} // Pass setStatus function
              />
            </div>
          ))}
      </div>
    </main>
  );
}
