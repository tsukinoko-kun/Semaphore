const colors = require("tailwindcss/colors")
/** @type {import('tailwindcss').Config} */
export default {
    content: ["./src/**/*.{html,css,tsx}", "./index.html"],
    theme: {
        extend: {
            fontFamily: {
                sans: ["Noto Sans", "Segoe UI", "sans-serif"],
            },
            colors: {
                neutral: {
                    50: "#fafafa",
                    100: "#f5f5f5",
                    200: "#e5e5e5",
                    300: "#d4d4d4",
                    400: "#a3a3a3",
                    500: "#737373",
                    600: "#525252",
                    700: "#404040",
                    800: "#262626",
                    900: "#171717",
                    950: "#0a0a0a",
                    955: "#090909",
                    960: "#080808",
                    965: "#070707",
                    970: "#060606",
                    975: "#050505",
                    980: "#040404",
                    985: "#030303",
                    990: "#020202",
                    995: "#010101",
                },
            },
        },
    },
    plugins: [],
}
