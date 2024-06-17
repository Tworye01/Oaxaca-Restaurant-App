"use client";
import { useEffect, useState } from "react";

export const Carousel = (): JSX.Element => {
  const [current, setCurrent] = useState(0);
  const images = [
    './HomeCarousel/Chicken-Quesadillas.jpg',
    './HomeCarousel/chickpea-aubergine-and-pepper-enchiladas.png',
    './HomeCarousel/Chilli-Cheese-Nachos.jpg',
    './HomeCarousel/Tacos.png',
    './HomeCarousel/churros.jpg'
  ];

  const previousSlide = () => {
    setCurrent((current - 1 + images.length) % images.length);
  };

  const nextSlide = () => {
    setCurrent((current + 1) % images.length);
  };

  useEffect(() => {
    const interval = setInterval(() => { nextSlide(); }, 4000);

    return () => clearInterval(interval);
  }, [current]);

  return (
    <div className="overflow-hidden relative" >
      <div className="flex transition ease-out duration-500" style={{ transform: `translateX(-${current * 100}%)` }}>
        {
          images.map((url, index) => (
            <img src={url} key={index} className="w-full" alt={`Slide ${index}`} />
          ))
        }
      </div>

      <div className="absolute top-0 h-full w-full flex justify-between items-center text-black px-10 text-3xl">
        <button className="bg-white border-black border-2 rounded pb-1 bg-opacity-80" onClick={previousSlide}>{'<'}</button>
        <button className="bg-white border-black border-2 rounded pb-1 bg-opacity-80" onClick={nextSlide}>{'>'}</button>
      </div>

      <div className="absolute bottom-0 py-4 flex justify-center gap-3 w-full">
        {images.map((url, index) => (
          <div
            onClick={() => setCurrent(index)}
            key={`circle-${index}`}
            className={`rounded-full w-5 h-5 cursor-pointer ${index === current ? "bg-white" : "bg-gray-500"}`}
          ></div>
        ))}
      </div>
    </div>
  );
};
