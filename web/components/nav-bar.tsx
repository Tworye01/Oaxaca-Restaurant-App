"use client";

import Link from 'next/link';
import { useState } from 'react';
import { BasketIcon } from "@/components/basket";


// Render the navigation bar with a background color, text color, and padding
export const NavigationBarComponent = () => {

  // Create two Variables set both useStates to false
  const [dropDownType, setdropDownType] = useState(String)

  // Function to toggle the dropdown state
  const setdropDown = (title: string) => {
    setdropDownType(title)
  };

  return (
    <nav className="grid grid-cols-2  w-screen z-10 bg-pinky-500 text-black fixed top-0">
      <div id="Oaxaca" className="mt-4 justify-self-start">
        <Link href={"/"} className="px-4 py-5 text-brown-dark-100 font-medium hover:bg-slate-100 hover:text-brown-dark-100 ">Oaxaca</Link>
      </div>

      <div id="MenuItems" className="justify-self-end pr-5">


        {[
          ['Home', '/'],
          ['Menu', '/menu'],
          ['About', '/about'],
          ['Services', '/services'],
          ['Contact', '/contact'],
          ['Login', '/login'],
          ['Table', '/table']
        ].map(([title, url]) => (
          <a
            key={title}
            href={url}
            onMouseEnter={() => setdropDownType(title)} onMouseLeave={() => setdropDownType("null")}
            className={`px-4 py-4 text-black font-medium hover:bg-slate-100 hover:text-slate-900 relative inline-block group `}> {title}
            {dropDownType === "Contact" && (


              <div id="menu-dropDownContent" className="px-6 py-4 left-0 bg-slate-100 absolute hidden group-hover:block">
                {[
                  ['FAQs', '/faqs', "FAQ"],].map(([title, url, id]) => (
                    <a
                      key={title}
                      href={url}
                      type="button"
                      className={`hover:text-gray-400 cursor-pointer block`}> {title}

                    </a>
                  ))}

              </div>

            )}
          </a>

        ))}

      </div>
    </nav>
  );
}
