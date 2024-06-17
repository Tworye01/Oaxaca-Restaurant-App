/* eslint-disable @next/next/no-img-element */

import React from 'react';
import { NavigationBarComponent } from '@/components/nav-bar';
import { Carousel } from '@/components/carousel';
import Link from 'next/link';

export default function Home() {
  return (
    <main className="text-center h-screen flex">
      <img src="./HomeCarousel/Tacos.png" alt="background" className="w-screen -translate-y-6 fixed blur-lg"></img>

      <div className="inset-0 h-full w-9/12 mx-auto bg-slate-100 mt-14 fixed">
        <Link href="/menu" type="button" className="text-center mt-8 font-semibold text-gray-900 bg-pinky-300 hover:bg-pinky-400 active:bg-pinky-500  p-4 px-8 shadow-lg mb-4">
          Order Now!
        </Link>

        <div className='mx-auto mt-10 max-h-10 w-9/12'>
          <Carousel />
        </div>
      </div>

      <NavigationBarComponent />
    </main>
  );
};