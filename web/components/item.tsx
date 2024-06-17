/* eslint-disable @next/next/no-img-element */

import {
  Dish,
  getCourseString,
  getSpiceString
} from "@/lib/types/dish";
import {
  ReactNode,
  useEffect,
  useState
} from "react";
import { BasketItem } from "@/lib/types/basket";
import { Drink } from "@/lib/types/drinks";
import { Item } from "@/lib/types/item";
import { saveToBasket } from "@/lib/utils/basket";
import { saveToStaffBasket } from "@/lib/utils/staff-basket";

export const AllergensComponent = (props: {
  allergens: string[],
}): JSX.Element => {
  return (
    <div>
      {
        props.allergens.map((allergen: string, index: number) =>
          <div
            className="mt-1 max-w-sm text-gray-600"
            key={index}
          >
            <p className="text-left">{allergen}</p>
          </div>
        )
      }
    </div>
  );
};

export const ItemComponent = (props: {
  item: Item,
  modal: ReactNode,
  children?: ReactNode,
}): JSX.Element => {
  const [visible, setVisible] = useState<boolean>(false);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        setVisible(false);
      }
    };

    if (visible) {
      document.body.style.overflow = 'hidden';
      document.addEventListener('keydown', handleKeyDown);
    } else {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [visible]);

  return (
    <>
      {
        props.item &&
        <button onClick={() => setVisible(true)}>
          <div className="block shadow-lg ">
            <div className="">
              {
                props.item.image &&
                <img
                  className="w-full object-cover h-48 sm:h-64 lg:h-80"
                  alt={`${props.item.name} image`}
                  src={props.item.image}
                />
              }
            </div>

            <div className="p-4">
              <div className="mt-4 text-lg font-bold text-gray-900 sm:text-xl">
                <h2 className="text-left">{props.item.name}</h2>
              </div>

              <div className="mt-2 max-w-sm text-gray-700 line-clamp-2 text-pretty">
                <p className="text-left">{props.item.description}</p>
              </div>
            </div>
          </div>
        </button>
      }

      {
        props.children && props.children
      }

      {
        visible && props.modal &&
        <center>
          <div className="z-10 fixed inset-0 flex justify-center items-center m-auto">
            <div className="overflow-y-auto overflow-x-hidden fixed backdrop-brightness-50 z-50 w-full md:inset-0 max-h-full">
              <div className="shadow-lg mt-24 max-w-md bg-white">
                {props.modal}
              </div>
            </div>
          </div>
        </center>
      }
    </>
  );
};

export const ItemStaffComponent = (props: {
  item: Item,
  modal: ReactNode,
  children?: ReactNode,
}): JSX.Element => {
  const [visible, setVisible] = useState<boolean>(false);

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        setVisible(false);
      }
    };

    if (visible) {
      document.body.style.overflow = 'hidden';
      document.addEventListener('keydown', handleKeyDown);
    } else {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [visible]);

  return (
    <>
      {
        props.item &&
        <button onClick={() => setVisible(true)}>
          <div className="block shadow-lg">

            <div className="p-4">
              <div className="mt-4 text-lg font-bold text-gray-900 sm:text-xl">
                <h2 className="text-left">{props.item.name}</h2>
              </div>

              <div className="mt-2 max-w-sm text-gray-700 line-clamp-2 text-pretty">
                <p className="text-left">{props.item.description}</p>
              </div>
            </div>
          </div>
        </button>
      }

      {
        props.children && props.children
      }

      {
        visible && props.modal &&
        <center>
          <div className="z-10 fixed inset-0 flex justify-center items-center m-auto">
            <div className="overflow-y-auto overflow-x-hidden fixed backdrop-brightness-50 z-50 w-full md:inset-0 max-h-full">
              <div className="shadow-lg mt-24 max-w-md bg-white">
                {props.modal}
              </div>
            </div>
          </div>
        </center>
      }
    </>
  );
};


export const DishModalComponent = (props: {
  dish: Dish,
}): JSX.Element => {
  return (
    <>
      <ModalComponent item={props.dish.item}>
        <div className="mt-4 max-w-s px-4">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Course: </h2>
          <p className="p-4 text-left text-gray-700">This dish is a {getCourseString(props.dish.course)}!</p>
        </div>

        <div className="mt-4 max-w-s px-4">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Spice Level: </h2>
          <p className="p-4 text-left text-gray-700">This dish is {getSpiceString(props.dish.spice)}!</p>
        </div>
      </ModalComponent>
    </>
  );
};

export const DishStaffModalComponent = (props: {
  dish: Dish,
}): JSX.Element => {
  return (
    <>
      <ModalStaffComponent item={props.dish.item}>
        <div className="mt-4 max-w-s px-4">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Course: </h2>
          <p className="p-4 text-left text-gray-700">This dish is a {getCourseString(props.dish.course)}!</p>
        </div>

        <div className="mt-4 max-w-s px-4">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Spice Level: </h2>
          <p className="p-4 text-left text-gray-700">This dish is {getSpiceString(props.dish.spice)}!</p>
        </div>
      </ModalStaffComponent>
    </>
  );
};

export const DrinkModalComponent = (props: {
  drink: Drink,
}): JSX.Element => {
  return (
    <>
      <ModalComponent item={props.drink.item}>
        <div className="mt-4 max-w-s px-4">
          {
            props.drink.alcoholic &&
            <>
              <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Additional Information: </h2>
              <p className="p-4 text-left text-gray-700">This item contains alcohol, you must be 18 or over to purchase this!</p>
            </>
          }
        </div>
      </ModalComponent>
    </>
  );
};

export const DrinkStaffModalComponent = (props: {
  drink: Drink,
}): JSX.Element => {
  return (
    <>
      <ModalStaffComponent item={props.drink.item}>
        <div className="mt-4 max-w-s px-4">
          {
            props.drink.alcoholic &&
            <>
              <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Additional Information: </h2>
              <p className="p-4 text-left text-gray-700">This item contains alcohol, you must be 18 or over to purchase this!</p>
            </>
          }
        </div>
      </ModalStaffComponent>
    </>
  );
};

export const DishEditModalComponent = (props: {
  dish: Dish,
  onSubmit: (dish: Dish, form?: FormData) => void;
  onDelete: (dish: Dish) => void;
}): JSX.Element => {
  const [course, setCourse] = useState<number>(props.dish.course);
  const [spice, setSpice] = useState<number>(props.dish.spice);
  const [vegetarian, setVegetarian] = useState<boolean>(props.dish.vegetarian);
  const [item, setItem] = useState<Item>(props.dish.item);
  const [image, setImage] = useState<FormData>();

  return (
    <>
      <ModalEditComponent
        item={item}
        updateItem={(updatedItem: Item) => setItem(updatedItem)}
        updateImage={(image: FormData) => setImage(image)}
      >
        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Course</h2>
          <input
            className="mt-2"
            type="number"
            defaultValue={props.dish.course}
            onChange={e => setCourse(parseInt(e.target.value))}
            min={0}
            max={3}
          />
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Spice Level</h2>
          <input
            className="mt-2"
            type="number"
            defaultValue={props.dish.spice}
            onChange={e => setSpice(parseInt(e.target.value))}
            min={0}
            max={4}
          />
        </div>

        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Vegetarian</h2>
        <input
          className="mt-2"
          type="checkbox"
          defaultChecked={props.dish.vegetarian}
          onChange={e => setVegetarian(e.target.checked)}
        />
      </ModalEditComponent>

      <div className="flex items-center justify-center gap-4">
        <SubmitButtonComponent onClick={() => {
          let dish: Dish = props.dish
          dish.item = item;
          dish.course = course;
          dish.spice = spice;
          dish.vegetarian = vegetarian;

          props.onSubmit(dish, image);
          window.location.reload();
        }} />

        <DeleteButtonComponent onClick={() => {
          let dish: Dish = props.dish

          props.onDelete(dish);
          window.location.reload();
        }} />
      </div>
    </>
  );
};

export const DrinkEditModalComponent = (props: {
  drink: Drink,
  onSubmit: (drink: Drink, form?: FormData) => void;
  onDelete: (drink: Drink) => void;
}): JSX.Element => {
  const [alcoholic, setAlcoholic] = useState<boolean>(props.drink.alcoholic);
  const [item, setItem] = useState<Item>(props.drink.item);
  const [image, setImage] = useState<FormData>();

  return (
    <>
      <ModalEditComponent
        item={item}
        updateItem={(updatedItem: Item) => setItem(updatedItem)}
        updateImage={(image: FormData) => setImage(image)}
      >
        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Alcoholic</h2>
          <input
            className="mt-2"
            type="checkbox"
            defaultChecked={alcoholic}
            onChange={e => setAlcoholic(e.target.checked)}
          />
        </div>
      </ModalEditComponent>

      <div className="flex items-center justify-center gap-4">
        <SubmitButtonComponent onClick={() => {
          let drink: Drink = props.drink
          drink.item = item;
          drink.alcoholic = alcoholic;

          props.onSubmit(drink, image);
          window.location.reload();
        }} />

        <DeleteButtonComponent onClick={() => {
          let drink: Drink = props.drink

          props.onDelete(drink);
          window.location.reload();
        }} />
      </div>
    </>
  );
};

export const ModalEditComponent = (props: {
  item: Item,
  updateItem: (updatedItem: Item) => void,
  updateImage: (image: FormData) => void,
  children?: ReactNode,
}): JSX.Element => {
  const [item, setItem] = useState<Item>(props.item);

  return (
    <>
      {
        props.item.image &&
        <img
          className="w-full object-cover h-48 sm:h-64 lg:h-80"
          alt={`${props.item.name} image`}
          src={props.item.image}
        />
      }

      <div className="p-4">
        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Name</h2>
          <input
            className="mt-2"
            defaultValue={props.item.name}
            onChange={e => {
              item.name = e.target.value;
              setItem(item);
              props.updateItem(item);
            }}
          />
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Description</h2>
          <textarea
            className="mt-2 flex-auto"
            defaultValue={props.item.description}
            onChange={e => {
              item.description = e.target.value;
              setItem(item);
              props.updateItem(item);
            }}
          />
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Cost</h2>
          <input
            className="mt-2"
            defaultValue={props.item.cost}
            onChange={e => {
              item.cost = parseFloat(e.target.value);
              setItem(item);
              props.updateItem(item);
            }}
          />
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Calories</h2>
          <input
            className="mt-2"
            defaultValue={props.item.calories}
            onChange={e => {
              item.calories = parseInt(e.target.value);
              setItem(item);
              props.updateItem(item);
            }}
          />
        </div>

        <div className="mt-2 max-w-s">
          <div>
            <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">This dish contains: </h2>
            <input
              className="mt-2"
              type="text"
              defaultValue={props.item.allergens}
              onChange={e => {
                item.allergens = e.target.value.split(/\s*,\s*/);
                setItem(item);
                props.updateItem(item);
              }}
            />
          </div>
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Available</h2>
          <input
            className="mt-2"
            type="checkbox"
            defaultChecked={props.item.available}
            onChange={e => {
              item.available = e.target.checked;
              setItem(item);
              props.updateItem(item);
            }}
          />
        </div>

        <div className="mt-4 max-w-s text-lg">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">Image</h2>
          <input
            className="mt-2"
            type="file"
            onChange={e => {
              if (e.target.files === null) {
                return;
              }

              const formData = new FormData();
              formData.append("img", e.target.files[0]);
              props.updateImage(formData);
            }}
          />
        </div>

        {
          props.children && props.children
        }
      </div>
    </>
  );
};

export const ModalStaffComponent = (props: {
  item: Item,
  children?: ReactNode,
  setVisible?: (visible: boolean) => void,
}): JSX.Element => {
  const addItemToBasket = () => {
    let item: BasketItem = {
      itemId: props.item.id,
      note: "",
    }

    saveToStaffBasket(item);
    window.location.reload();
  }

  return (
    <>
      <div className="mt-4 max-w-s text-lg px-4">
        <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">{props.item.name}</h2>
      </div>

      <div className="mt-4 max-w-sm">
        <p className="text-left text-gray-700">{props.item.calories} kcal | £{props.item.cost}</p>
        <p className="text-left text-gray-700">{props.item.description}</p>
      </div>

      <div className="mt-4 max-w-s px-4">
        {
          (props.item.allergens.length >= 1 && props.item.allergens[0] !== '') &&
          <>
            <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">This product contains: </h2>
            <div className="p-4">
              <AllergensComponent allergens={props.item.allergens} />
            </div>
          </>
        }
      </div>

      {props.children && props.children}

      <div className="mt-8">
        <button
          className="text-center font-semibold text-gray-900 bg-pinky-300 hover:bg-pinky-600 active:bg-pinky-600 disabled:bg-slate-300 p-4 px-8 shadow-lg mb-4"
          disabled={!props.item.available}
          onClick={() => addItemToBasket()}
        >
          {props.item.available && "Add to Order"}
          {!props.item.available && "Item Unavailable"}
        </button>
      </div>
    </>
  );
};

export const ModalComponent = (props: {
  item: Item,
  children?: ReactNode,
  setVisible?: (visible: boolean) => void,
}): JSX.Element => {
  const addItemToBasket = () => {
    let item: BasketItem = {
      itemId: props.item.id,
      note: "",
    }

    saveToBasket(item);
    window.location.reload();
  }

  return (
    <>
      {
        props.item.image &&
        <img
          className="w-full object-cover h-48 sm:h-64 lg:h-80"
          alt={`${props.item.name} image`}
          src={props.item.image}
        />
      }

      <>
        <div className="mt-4 max-w-s text-lg px-4">
          <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">{props.item.name}</h2>
        </div>

        <div className="mt-4 max-w-sm">
          <p className="text-left text-gray-700">{props.item.calories} kcal | £{props.item.cost}</p>
          <p className="text-left text-gray-700">{props.item.description}</p>
        </div>

        <div className="mt-4 max-w-s px-4">
          {
            (props.item.allergens.length >= 1 && props.item.allergens[0] !== '') &&
            <>
              <h2 className="text-left font-semibold text-gray-900 bg-slate-300 p-4 shadow-lg">This product contains: </h2>
              <div className="p-4">
                <AllergensComponent allergens={props.item.allergens} />
              </div>
            </>
          }
        </div>

        {props.children && props.children}

        <div className="mt-8">
          <button
            className="text-center font-semibold text-gray-900 bg-pinky-300 hover:bg-pinky-600 active:bg-pinky-600 disabled:bg-slate-300 p-4 px-8 shadow-lg mb-4"
            disabled={!props.item.available}
            onClick={() => addItemToBasket()}
          >
            {props.item.available && "Add to Order"}
            {!props.item.available && "Item Unavailable"}
          </button>
        </div>
      </>
    </>
  );
};

export const SubmitButtonComponent = (props: {
  onClick: () => void,
}) => {
  return (
    <button
      className="text-center font-semibold text-gray-900 bg-pinky-500 hover:bg-pinky-600 active:bg-pinky-700  p-4 px-8 shadow-lg mb-4"
      onClick={() => props.onClick()}
    >
      Submit
    </button>
  );
};

export const DeleteButtonComponent = (props: {
  onClick: () => void,
}) => {
  return (
    <button
      className="text-center font-semibold text-gray-900 bg-red-300 hover:bg-red-400 active:bg-red-500  p-4 px-8 shadow-lg mb-4"
      onClick={() => props.onClick()}
    >
      Delete
    </button>
  );
};
