import type { Config } from "tailwindcss";

export default {
  content: ["./src/**/*.{html,js,ts,svelte}"],
  theme: {
    extend: {
      colors: {
        "gray": {
          50: "#f5f5f5",
          100: "#f2f2f2",
          200: "#e6e6e6",
          300: "#d6d6d6",
          400: "#c4c4c4",
          500: "#b3b3b3",
          600: "#9d9d9d",
          700: "#858585",
          800: "#616161",
          900: "#1e1e1e"
        },
        "starter": "#c1c1c1",
        "iron": "#c29f7d",
        "bronze": "#e39067",
        "silver": "#637078",
        "gold": "#fac924",
        "platinum": "#00dd77",
        "diamond": "#674ac9",
        "heroes": "#0081ff"
      },
    },
    spacing: {
      0: "0px",
      "3xs": "0.25rem",
      "2xs": "0.5rem",
      "xs": "1rem",
      "sm": "1.5rem",
      "md": "2rem",
      "lg": "2.5rem",
      "xl": "3rem",
      "2xl": "3.5rem",
      "3xl": "4rem",
      "4xl": "5rem",
      "5xl": "6rem",
      "6xl": "7rem",
      "7xl": "8rem",
      "8xl": "9rem",
      "9xl": "10rem",
      "10xl": "11rem",
      "11xl": "12rem",
      "12xl": "13rem",
      "13xl": "14rem",
      "14xl": "15rem",
      "15xl": "16rem",
      "16xl": "18rem",
      "17xl": "20rem",
      "18xl": "24rem"
    },
    borderRadius: {
      DEFAULT: "16px",
      "lg": "24px",
      "full": "9999px"
    }
  },
  plugins: [],
} satisfies Config;