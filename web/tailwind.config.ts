import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
    },
    colors: {
      ...require("tailwindcss/colors"),
      pinky: {
        100: "#ffe0ec",
        200: "#ffb3c6",
        300: "#ff85a1",
        400: "#ff5f86", // Lighter shade
        500: "#ff91af",
        600: "#ff6d98",
        700: "#ff4d7f",
        800: "#ff3266",
        900: "#ff0e44",
      },
      warm: {
        100: `#fff7ed`,
      },
      'brown-dark' : {
        100: `#452410`
      },
    },
  },
  plugins: [],
};
export default config;
