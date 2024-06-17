import React from 'react';
import { NavigationBarComponent } from '@/components/nav-bar';
import { Login } from '@/components/login';

export default function Home() {
  return (
    <main className="text-center min-h-screen bg-warm-100">
      <NavigationBarComponent />
      <div>
        <Login />
      </div>
    </main>
  );
};
