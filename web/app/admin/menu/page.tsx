'use client'

import { DividerComponent } from "@/components/divider";
import { Dish, getBlankDish } from "@/lib/types/dish"
import { Drink, getBlankDrink } from "@/lib/types/drinks";
import { useToken } from "@/lib/utils/token";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import {
  ItemComponent,
  DishEditModalComponent,
  DrinkEditModalComponent,
} from "@/components/item";
import { isDish } from "@/lib/utils/utils";
import { LoadingComponent } from "@/components/loading";
import { NavigationBarStaffComponent } from "@/components/nav-bar-staff";

export default function Menu() {
  const [dishes, setDishes] = useState<Dish[]>();
  const [drinks, setDrinks] = useState<Drink[]>();
  const [loading, setLoading] = useState(true);

  const [token] = useToken();
  const router = useRouter();

  const blankDish: Dish = getBlankDish();
  const blankDrink: Drink = getBlankDrink();

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  const fetchData = async () => {
    await axios.get(`http://${apiAddr}/api/v1/menu/dishes`, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    })
      .then(response => {
        console.log('Response:', response.data);
        setDishes(response.data)
      })
      .catch(error => {
        console.error('Error:', error);
      });

    await axios.get(`http://${apiAddr}/api/v1/menu/drinks`, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      }
    })
      .then(response => {
        console.log('Response:', response.data);
        setDrinks(response.data)
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };

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

  if (loading) {
    return <LoadingComponent />;
  }

  const patchRequest = async (item: Dish | Drink, form?: FormData) => {
    let type: string;
    if (isDish(item)) {
      type = "dishes";
    } else {
      type = "drinks";
    }

    if (form == null) {
      form = new FormData();
    }

    form.append("body", JSON.stringify(item))

    axios.patch(`http://${apiAddr}/api/v1/menu/${type}`, form, {
      headers: {
        'Content-Type': 'multipart/form-data',
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

  const postRequest = async (item: Dish | Drink, form?: FormData) => {
    let type: string;
    if (isDish(item)) {
      type = "dishes";
    } else {
      type = "drinks";
    }

    if (form == null) {
      form = new FormData();
    }

    form.append("body", JSON.stringify(item))

    axios.post(`http://${apiAddr}/api/v1/menu/${type}`, form, {
      headers: {
        'Content-Type': 'multipart/form-data',
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

  const deleteRequest = async (item: Dish | Drink) => {
    let type: string;
    if (isDish(item)) {
      type = "dishes";
    } else {
      type = "drinks";
    }

    axios.delete(`http://${apiAddr}/api/v1/menu/${type}/${item.id}`, {
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

  return (
    <main>
      <div className="w-full relative items-center text-center bg-warm-100 text-brown-dark-100 pb-16">
        <NavigationBarStaffComponent />

        <div className="min-h-screen mb-24">
          <h1 className="font-bold text-3xl mb-8 pt-20">Starters</h1>
          {
            dishes && dishes.filter(dish => dish.course === 0).length === 0 ? (
              <div className="text-center  items-center justify-center">
                <p className="font-bold text-2xl mb-8 pt-20">Starters not found</p>
              </div>
            ) : null
          }
          <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
            {
              dishes && dishes.map((dish: Dish, index: number) => (
                dish.course === 0 ? (
                  <div key={index}>
                    <ItemComponent
                      item={dish.item}
                      modal={
                        <DishEditModalComponent
                          dish={dish}
                          onSubmit={(dish: Dish, form?: FormData) => patchRequest(dish, form)}
                          onDelete={(dish: Dish) => deleteRequest(dish)}
                        />
                      }
                    />
                  </div>
                ) : null
              ))
            }
          </div>

          <div className="z-0 min-w-full">
            <DividerComponent />
          </div>


          <h1 className="font-bold text-3xl mb-8 pt-20">Mains</h1>
          {
            dishes && dishes.filter(dish => dish.course === 1).length === 0 ? (
              <div className="text-center  items-center justify-center">
                <p className="font-bold text-2xl mb-8 pt-20">Mains not found</p>
              </div>
            ) : null
          }
          <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
            {
              dishes && dishes.map((dish: Dish, index: number) => (
                dish.course === 1 ? (
                  <div key={index}>
                    <ItemComponent
                      item={dish.item}
                      modal={<DishEditModalComponent
                        dish={dish}
                        onSubmit={(dish: Dish, form?: FormData) => patchRequest(dish, form)}
                        onDelete={(dish: Dish) => deleteRequest(dish)}
                      />}
                    />
                  </div>
                ) : null
              ))
            }
          </div>

          <div className="z-0 min-w-full">
            <DividerComponent />
          </div>

          <h1 className="font-bold text-3xl mb-8 pt-20">Desserts</h1>
          {
            dishes && dishes.filter(dish => dish.course === 2).length === 0 ? (
              <div className="text-center  items-center justify-center">
                <p className="font-bold text-2xl mb-8 pt-20">Desserts not found</p>
              </div>
            ) : null
          }
          <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
            {
              dishes && dishes.map((dish: Dish, index: number) => (
                dish.course === 2 ? (
                  <div key={index}>
                    <ItemComponent
                      item={dish.item}
                      modal={<DishEditModalComponent
                        dish={dish}
                        onSubmit={(dish: Dish, form?: FormData) => patchRequest(dish, form)}
                        onDelete={(dish: Dish) => deleteRequest(dish)}
                      />}
                    />
                  </div>
                ) : null
              ))
            }
          </div>

          <div className="z-0 min-w-full">
            <DividerComponent />
          </div>

          <h1 className="font-bold text-3xl mb-8 pt-20">Sides</h1>
          {
            dishes && dishes.filter(dish => dish.course === 3).length === 0 ? (
              <div className="text-center  items-center justify-center">
                <p className="font-bold text-2xl mb-8 pt-20">Sides not found</p>
              </div>
            ) : null
          }
          <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
            {
              dishes && dishes.map((dish: Dish, index: number) => (
                dish.course === 3 ? (
                  <div key={index}>
                    <ItemComponent
                      item={dish.item}
                      modal={
                        <DishEditModalComponent
                          dish={dish}
                          onSubmit={(dish: Dish, form?: FormData) => patchRequest(dish, form)}
                          onDelete={(dish: Dish) => deleteRequest(dish)}
                        />
                      }
                    />
                  </div>
                ) : null
              ))
            }
          </div>

          <div className="z-0 min-w-full">
            <DividerComponent />
          </div>

          <ItemComponent
            item={blankDish.item}
            modal={
              <DishEditModalComponent
                dish={blankDish}
                onSubmit={(dish: Dish, form?: FormData) => postRequest(dish, form)}
                onDelete={(dish: Dish) => deleteRequest(dish)}
              />
            }
          />
        </div>

        <div className="min-w-full">
          <DividerComponent />
        </div>

        <h1 className="font-bold text-xl mb-8">Drinks</h1>
        <div className="grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            drinks && drinks.map((drink: Drink, index: number) => (
              <div key={index}>
                <ItemComponent
                  item={drink.item}
                  modal={
                    <DrinkEditModalComponent
                      drink={drink}
                      onSubmit={(drink: Drink, form?: FormData) => patchRequest(drink, form)}
                      onDelete={(drink: Drink) => deleteRequest(drink)}
                    />
                  }
                />
              </div>
            ))
          }
          <ItemComponent
            item={blankDrink.item}
            modal={
              <DrinkEditModalComponent
                drink={blankDrink}
                onSubmit={(drink: Drink, form?: FormData) => postRequest(drink, form)}
                onDelete={(drink: Drink) => deleteRequest(drink)}
              />
            }
          />
        </div>
      </div>
    </main >
  );
};
