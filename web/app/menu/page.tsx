/* eslint-disable @next/next/no-img-element */
'use client'

import { DividerComponent } from "@/components/divider";
import { NavigationBarComponent } from "@/components/nav-bar";
import { Dish } from "@/lib/types/dish"
import { Drink } from "@/lib/types/drinks";
import { useEffect, useState } from "react";
import axios from "axios";
import {
  ItemComponent,
  DishModalComponent,
  DrinkModalComponent
} from "@/components/item";
import { FilterComponent } from "@/components/filter";
import { buildFilterQuery } from "@/lib/utils/query";
import { LoadingComponent } from "@/components/loading";
import { BasketIcon } from "@/components/basket";

export default function Menu() {
  const [dishes, setDishes] = useState<Dish[]>();
  const [drinks, setDrinks] = useState<Drink[]>();
  const [selectedFilters, setSelectedFilters] = useState<string[]>([]);
  const [caloriesRange, setCaloriesRange] = useState<number[]>([0, 1500]);

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR;

  const fetchData = async () => {
    const filter = buildFilterQuery(selectedFilters, caloriesRange);
    console.log(filter)
    await axios.get(`http://${apiAddr}/api/v1/menu/dishes${filter}`, {
      headers: {
        'Content-Type': 'application/json'
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
        'Content-Type': 'application/json'
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
    fetchData();
  }, [selectedFilters, caloriesRange])

  if (!dishes || !drinks) return <LoadingComponent />
  return (
    <main className="text-brown-dark-100 text-center flex">
      <NavigationBarComponent />
      <img src="./HomeCarousel/Tacos.png" alt="background" className="w-screen -translate-y-6 fixed blur-lg" />

      <div className="relative inset-0 min-h-screen w-9/12 mx-auto bg-warm-100 mt-14 pb-16">
        <BasketIcon />


        <h1 className="w-full font-bold text-3xl mb-8 pt-10">Starters</h1>
        <FilterComponent
          selectedFilters={selectedFilters}
          setSelectedFilters={setSelectedFilters}
          caloriesRange={caloriesRange}
          setCaloriesRange={setCaloriesRange}
        />
        {
          dishes.filter(dish => dish.course === 0).length === 0 ? (
            <div className="text-center items-center justify-center">
              <p className="font-bold text-2xl mb-8 pt-20">Starters not found</p>
            </div>
          ) : null
        }
        <div className="z-10 grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            dishes.map((dish: Dish, index: number) => (
              dish.course === 0 ? (
                <div key={index}>
                  <ItemComponent
                    item={dish.item}
                    modal={<DishModalComponent dish={dish} />}
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
          dishes.filter(dish => dish.course === 1).length === 0 ? (
            <div className="text-center items-center justify-center">
              <p className="font-bold text-2xl mb-8 pt-20">Mains not found</p>
            </div>
          ) : null
        }
        <div className="z-10 grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            dishes.map((dish: Dish, index: number) => (
              dish.course === 1 ? (
                <div key={index}>
                  <ItemComponent
                    item={dish.item}
                    modal={<DishModalComponent dish={dish} />}
                  />
                </div>
              ) : null
            ))
          }
        </div>

        <div className="z-0 min-w-full">
          <DividerComponent />
        </div>

        <h1 className="z-10 font-bold text-3xl mb-8 pt-20">Desserts</h1>
        {
          dishes.filter(dish => dish.course === 2).length === 0 ? (
            <div className="text-center  items-center justify-center">
              <p className="font-bold text-2xl mb-8 pt-20">Desserts not found</p>
            </div>
          ) : null
        }
        <div className="z-10 grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            dishes.map((dish: Dish, index: number) => (
              dish.course === 2 ? (
                <div key={index}>
                  <ItemComponent
                    item={dish.item}
                    modal={<DishModalComponent dish={dish} />}
                  />
                </div>
              ) : null
            ))
          }
        </div>
        <div className="z-0 min-w-full">
          <DividerComponent />
        </div>

        <h1 className="z-10 font-bold text-3xl mb-8 pt-20">Sides</h1>
        {
          dishes.filter(dish => dish.course === 3).length === 0 ? (
            <div className="text-center items-center justify-center">
              <p className="font-bold text-2xl mb-8 pt-20">Sides not found</p>
            </div>
          ) : null
        }

        <div className="z-10 grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            dishes.map((dish: Dish, index: number) => (
              dish.course === 3 ? (
                <div key={index}>
                  <ItemComponent
                    item={dish.item}
                    modal={<DishModalComponent dish={dish} />}
                  />
                </div>
              ) : null
            ))
          }
        </div>

        <div className="z-0 min-w-full">
          <DividerComponent />
        </div>

        <h1 className="z-10 font-bold text-3xl mb-8">Drinks</h1>
        <div className="z-10 grid grid-cols-2 w-1/2 lg:grid-cols-3 lg:gap-8 mx-auto">
          {
            drinks.map((drink: Drink, index: number) => (
              <div key={index}>
                <ItemComponent
                  item={drink.item}
                  modal={<DrinkModalComponent drink={drink} />}
                />
              </div>
            ))
          }
        </div>
      </div>
    </main>
  );
};
