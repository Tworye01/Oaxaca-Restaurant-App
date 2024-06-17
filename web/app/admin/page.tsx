'use client'

import { NavigationBarStaffComponent } from "@/components/nav-bar-staff";
import { useToken } from "@/lib/utils/token";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { LoadingComponent } from "@/components/loading";

export default function Menu() {
  const [token] = useToken();
  const router = useRouter();
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (!token) {
      router.push("/login");
    } else {
      setLoading(false);
    }
  }, [token, router]);

  if (loading) {
    return <LoadingComponent />;
  }
  const handleKitchenClick = () => {
    router.push("/admin/kitchen");
  };

  const handleWaiterClick = () => {
    router.push("/admin/waiter");
  };

  return (
    <main className="w-full relative items-center text-center bg-warm-100">
      <NavigationBarStaffComponent />
      <div className="min-h-screen flex justify-center items-center">
        <button onClick={handleKitchenClick} className="mr-4 relative">
          <img src="/kitchen.jpg" alt="Kitchen" className="w-96 h-96 rounded-lg opacity-40" />
          <span className="absolute text-3xl inset-0 flex justify-center items-center font-bold text-brown-dark-100">Kitchen</span>
        </button>
        <button onClick={handleWaiterClick} className="relative">
          <img src="/waiter.jpg" alt="Waiter" className="w-96 h-96 rounded-lg opacity-40" />
          <span className="absolute  text-3xl inset-0 flex justify-center items-center font-bold text-brown-dark-100">Waiter</span>
        </button>
      </div>
    </main>
  );
}