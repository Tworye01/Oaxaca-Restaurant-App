import { Bill } from "@/lib/types/bill";
import { Order } from "@/lib/types/order";
import axios from "axios";
import { useEffect, useState } from "react";
import { LoadingComponent } from "./loading";
import { AccordionComponent } from "./accordion";
import { OrderComponent } from "./order";
import { Item } from "@/lib/types/item";

type Response = {
  bill: Bill;
  orders: Order[];
  ordersCount: number;
}

export const BillOrderComponent = (props: {
  item: Item,
}): JSX.Element => {
  return (
    <>

    </>
  );
};

export const BillComponent = (props: {
  bill: Bill,
}): JSX.Element => {
  const [response, setResponse] = useState<Response>();

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  useEffect(() => {
    const fetchData = async () => {
      await axios.get(`http://${apiAddr}/api/v1/bill/orders/${props.bill.id}`, {
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => {
          console.log('Response:', response.data);
          setResponse(response.data)
        })
        .catch(error => {
          console.error('Error:', error);
        });
    }

    fetchData();
  }, [])

  if (!response) return <LoadingComponent />

  return (
    <div>
      <div className="mt-4 max-w-s px-4">

      </div>

      <AccordionComponent
        preview={
          <h1 className="text-left ml-4">{`Bill ${props.bill.id} - £${props.bill.cost}`}</h1>
        }
        contents={
          <div className="mt-4 max-w-s px-4">
            {
              response.orders.map((order: Order, index: number) => (
                <div key={index} className="mb-4 w-full max-w-md">
                  <AccordionComponent
                    preview={
                      <h1 className="text-left ml-4">{`Order ${order.id} - ${order.item.name}`}</h1>
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
      />
    </div>
  );
};
