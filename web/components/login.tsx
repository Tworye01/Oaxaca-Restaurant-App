'use client'

import React, { useState } from "react";
import axios from "axios";
import { useToken } from "../lib/utils/token";
import { useRouter } from 'next/navigation'

type Creds = {
  username: string;
  password: string;
}

export const Login = (): JSX.Element => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false); // Track login success
  const [token, setToken] = useToken();
  const router = useRouter()

  const apiAddr = process.env.NEXT_PUBLIC_API_ADDR

  const handleSubmit = async () => {
    if (!username || !password) { // Check if both username and password are filled
      setError("Please fill in both username and password");
      return;
    }

    const creds: Creds = {
      username: username,
      password: password
    };

    axios.post(`http://${apiAddr}/api/v1/login`, creds)
      .then(response => {
        console.log("Login successful. Token:", response.data);
        if (!token) {
          setToken(response.data)
        }
        setSuccess(true);
        setError("")
        router.push("/admin");
      })
      .catch(error => {
        console.log("Error:", error);
        setError("Invalid login details");
        setSuccess(false);
        setUsername("");
        setPassword("");
      });
  };

  return (
    <div className="flex items-center justify-center ">
      <div className="login-wrapper bg-gray-100 p-10 rounded-lg shadow-md  mt-60">
        <h1 className="text-2xl text-brown-dark-100 font-bold mb-4">Please Log In</h1>
        <form>
          <div className="mb-4">
            <label
              htmlFor="username"
              className="block text-sm font-medium text-gray-700"
            >
              Username
            </label>
            <input
              id="username"
              type="text"
              className="mt-1 p-2 block w-full rounded-md border-gray-300 focus:border-pinky-500 focus:ring focus:ring-pinky-500 focus:ring-opacity-50"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
            />
          </div>
          <div className="mb-4">
            <label
              htmlFor="password"
              className="block text-sm font-medium text-gray-700"
            >
              Password
            </label>
            <input
              id="password"
              type="password"
              className="mt-1 p-2 block w-full rounded-md border-gray-200 focus:border-pinky-500 focus:ring focus:ring-pinky-500 focus:ring-opacity-50"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          {error && <div className="text-red-500 mb-4">{error}</div>}
        </form>
        <div className="flex justify-center">
          <button
            type="button"
            className="px-4 py-2 rounded-md bg-pinky-600 text-brown-dark-100  hover:bg-pinky-700 focus:outline-none shadow-lg"
            onClick={handleSubmit}
          >
            Submit
          </button>
        </div>
        {success && (
          <p className="text-green-500 mt-4">Login successful!</p>
        )}
      </div>
    </div>
  );
};
