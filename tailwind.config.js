/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['**/*.{templ,go}'],
  theme: {
    extend: {},
  },
  darkMode: 'class',
  plugins: [require('@tailwindcss/forms'), require('@tailwindcss/typography')],
}
