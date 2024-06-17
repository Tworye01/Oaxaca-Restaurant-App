'use client'

import React, { useState } from 'react';
import Link from 'next/link';
import { NavigationBarStaffComponent } from '@/components/nav-bar-staff';

export default function Table() {
  const [table, setTable] = useState<number>(0);

  return (
    <main className="w-full relative items-center text-center bg-warm-100">
      <NavigationBarStaffComponent />

      <div className="min-h-screen items-center justify-center mb-24 pt-20 w-1/2 mx-auto">
        <h1 className="text-brown-dark-100 p-2 text-2xl font-bold"> Assign yourself to a table </h1>
        <div className="login-wrapper bg-gray-100 p-8 rounded-lg shadow-md">
          <div className="mb-4">
            <label className="block text-sm font-medium text-gray-700">Table Number</label>
            <input
              type="number"
              className="mt-1 p-2 block w-full rounded-md border-gray-300 focus:border-pinky-500 focus:ring focus:ring-ppinky-500 focus:ring-opacity-50"
              placeholder="Table Number"
              onChange={(e) => setTable(parseInt(e.target.value))} />
          </div>

          <div className="flex justify-center">
            <Link
              className="px-4 py-2 bg-pinky-600 text-black rounded-md hover:bg-pinky-700 focus:outline-none shadow-lg focus:bg-pinky-700"
              href={`/admin/waiter/${table}`}
            >
              Submit
            </Link>
          </div>
        </div>
      </div>
    </main>
  );
};
