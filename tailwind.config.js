/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./templates/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: ["emerald"],
    darkTheme: "emerald",
  },
};
