import {
  Order,
  Status,
  getStatusString
} from "@/lib/types/order";
import { Item } from "@/lib/types/item";
import axios from "axios";
import { ReactNode, useState } from "react";
import { getToken, useToken } from "@/lib/utils/token";
import { LoadingComponent } from "./loading";

export const PaidButtonComponent = (props: { order: Order }) => {
  const [paidClicked, setPaidClicked] = useState(false);

  const handlePaidClick = async () => {
    if (props.order.bill.paid) {
      return;
    } else {
      const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;
      props.order.bill.paid = true;
      await axios
        .patch(`http://${apiAddr}/api/v1/bill/${props.order.bill.id}`, props.order.bill, {
          headers: {
            'Content-Type': 'application/json',
          },
        })
        .then((response) => {
          console.log('Response:', response.data);
          setPaidClicked(true);
        })
        .catch((error) => {
          console.error('Error:', error);
        });
    }
  };

  return (
    <button
      className={`px-4 py-2 text-white ${props.order.bill.paid ? 'bg-green-500' : 'bg-red-500'} shadow-lg`}
      onClick={handlePaidClick}
      disabled={props.order.bill.paid}
    >
      {props.order.bill.paid ? 'Paid' : 'Pay'}
    </button>
  );
};

export const OrderComponent = (props: {
  order: Order,
  showId?: boolean,
  children?: ReactNode
}): JSX.Element => {
  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;
  const [token] = useToken();

  const patchStatus = async () => {
    await axios.patch(`http://${apiAddr}/api/v1/orders/${props.order.id}/6`, {}, {
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
  }

  if (!token) return <LoadingComponent />

  return (
    <div className="shadow-lg items-center justify-center p-2">
      <div className="justify-center items-center mb-2">
        <OrderItemComponent
          item={props.order.item}
          status={
            <div className="mt-2">
              <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Status</h2>
              <p className="text-left p-2 mt-2 text-gray-700">{getStatusString(props.order.status)}</p>
            </div>
          }
        />
      </div>

      {
        props.showId &&
        <div className=" justify-center items-center mb-2">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Order number</h2>
          <p className="text-left text-gray-700 p-2 mt-2">{props.order.id}</p>
        </div>
      }

      <div className=" justify-center items-center mb-2">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Time of order</h2>
        <p className="text-left text-gray-700 p-2 mt-2">Time of order: {new Date(props.order.createdAt).toLocaleTimeString()}</p>
      </div>
      <div className=" justify-center items-center mb-2">
        <OrderInformationComponent order={props.order} />
      </div>

      <div className="justify-center items-center mb-1">
        <CancelButtonComponent onClick={() => patchStatus()} />
      </div>

      {
        props.children && props.children
      }
    </div>
  );
};

export const OrderKitchenComponent = (props: {
  order: Order,
  status: number,
  setStatus: (status: number) => void
}): JSX.Element => {
  const { order, status, setStatus } = props;
  const [apiAddr] = useState<string>(process.env.NEXT_PUBLIC_API_ADDR || '');

  const patchStatus = async () => {
    try {
      const response = await axios.patch(`http://${apiAddr}/api/v1/orders/${order.id}/${status}`, {}, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      console.log('Response:', response.data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const deleteOrder = async () => {
    try {
      const response = await axios.delete(`http://${apiAddr}/api/v1/orders/${order.id}`, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      console.log('Response:', response.data);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const statusClassName = status === 6 ? "bg-red-200" : "";

  return (
    <div className={`mt-20 shadow-lg p-4 ${statusClassName}`}>
      <div className="flex justify-end">
        <button
          className="text-gray-600 hover:text-red-500"
          onClick={() => deleteOrder()}
        >
          X
        </button>
      </div>

      <div className="items-center mb-2">
        <h1 className="text-2xl font-bold">Order {order.id}</h1>
      </div>

      <div className="items-center mb-2">
        <h1 className="text-xl font-bold">Time of order: {new Date(order.createdAt).toLocaleTimeString()}</h1>
      </div>

      <div className="items-center mb-2">
        <p className="text-gray-600">Table {order.table}</p>
      </div>

      <div className="items-center mb-2">
        <OrderKitchenItemComponent item={order.item} order={order} status={
          <div className="flex flex-wrap items-center mb-2">
            <label className="text-gray-600">Status: </label>
            <input
              className="w-36"
              type="number"
              id="status"
              value={status}
              onChange={e => setStatus(parseInt(e.target.value))}
            />
          </div>
        } />
      </div>

      <button className="justify-center text-center bg-pinky-500 hover:bg-pinky-600 active:bg-pinky-600 p-1 px-4 shadow-lg mb-2" onClick={patchStatus}>Update</button>
    </div>
  );
};

export const OrderInformationComponent = (props: {
  order: Order,
}): JSX.Element => {
  return (
    <>
      <div className="mt-4 max-w-s">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg">Cost</h2>
        <p className="text-left text-gray-700 p-2 mt-2">£{props.order.cost}</p>
      </div>

      <div className="mt-4 max-w-s">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg">Notes</h2>
        <p className="text-left text-gray-700 p-2 mt-2">{props.order.note}</p>
      </div>
    </>
  );
}

export const OrderKitchenItemComponent = (props: {
  item: Item,
  order: Order,
  status: ReactNode,
}): JSX.Element => {
  return (
    <>
      <div className="mt-4 max-w-s">
        <h2 className="mb-2 text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg">{props.item.name}</h2>
        <p className="text-lg mb-2">Note: {props.order.note}</p>
      </div>

      {
        props.status && props.status
      }
    </>
  );
};

export const OrderItemComponent = (props: {
  item: Item,
  status: ReactNode,
}): JSX.Element => {
  return (
    <>
      <div className="mt-4 max-w-s">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg text-lg">{props.item.name}</h2>
        <p className="text-left text-gray-700 p-2 mt-2">{props.item.calories} kcal</p>
        <p className="text-left text-gray-700 p-2">{props.item.description}</p>
      </div>

      {
        props.status && props.status
      }
    </>
  );
};

export const CancelButtonComponent = (props: {
  onClick: () => void,
}): JSX.Element => {
  return (
    <button
      className="text-center justify-center font-semibold text-gray-900 bg-pinky-500 hover:bg-pinky-600 active:bg-pinky-600  p-4 px-8 shadow-lg mb-4"
      onClick={() => props.onClick()}
    >
      Cancel
    </button>
  );
};
