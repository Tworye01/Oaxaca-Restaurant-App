"use client";

import Link from 'next/link';
import { useState } from 'react';

// Render the navigation bar with a background color, text color, and padding
export const NavigationBarStaffComponent = () => {

  // Create two Variables set both useStates to false
  const [dropDownType, setdropDownType] = useState(String)

  // Function to toggle the dropdown state
  const setdropDown = (title: string) => {
    setdropDownType(title)
  };

  return (
    <nav className="z-50">
      <div className="grid grid-cols-2 bg-pinky-500 text-black w-screen fixed">

        <div id="Oaxaca" className="mt-4 justify-self-start">
          <Link href={"/"} className="px-4 py-5 text-black font-medium hover:bg-slate-100 hover:text-slate-900 ">Oaxaca</Link>
        </div>

        <div id="MenuItems" className="justify-self-end pr-5">
          {[
            ['Home', '/admin'],
            ['Menu', '/admin/menu'],
            ['Kitchen', '/admin/kitchen'],
            ['Waiter', '/admin/waiter'],
            ['Order', '/admin/order'],
          ].map(([title, url]) => (
            <Link
              key={title}
              href={url}
              onMouseEnter={() => setdropDownType(title)} onMouseLeave={() => setdropDownType("null")}
              className={`px-4 py-4 text-black font-medium hover:bg-slate-100 hover:text-slate-900 relative inline-block group `}
            >
              {title}


            </Link>
          ))}
        </div>
      </div>
    </nav>
  );
}
