import { useState } from "react";

export function useToken(): [string | null, (token: string) => void] {
  const getToken = (): string | null => {
    if (typeof window !== "undefined") {
      const tokenString = localStorage.getItem("token");
      if (tokenString != null) {
        return tokenString;
      }
    }

    return null;
  };

  const [token, setToken] = useState<string | null>(getToken());

  const saveToken = (token: string) => {
    localStorage.setItem("token", token);
    setToken(token);
  };

  return [token, saveToken];
}

export function getToken(): string | undefined {
  if (typeof window !== "undefined") {
    const token = localStorage.getItem("token");
    if (token == null) {
      return undefined;
    }

    return token;
  }

  return undefined;
}
