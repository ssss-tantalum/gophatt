/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./templates/**/*.{templ,go}'],
  theme: {
    extend: {},
  },
  darkMode: 'class',
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
