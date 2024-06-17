import { useState } from "react";

export function Clock() {
  let time = new Date().toLocaleTimeString()

  const [ctime, setTime] = useState(time)
  const UpdateTime = () => {
    time = new Date().toLocaleTimeString()
    setTime(time)
  }
  setInterval(UpdateTime)
  return (
    <div className="fixed bottom-4 right-4 bg-gray-200 p-2 rounded">
      <div className='text-4xl font-bold flex'>
        <span>{ctime}</span>
      </div>
    </div>

  );
}
