/** @type {import('tailwindcss').Config} */
export default {
  content: ['./{routes,lib}/**/*.{html,js,svelte,ts}', 'app.html'],
  theme: {
    extend: {},
    colors: {
      white: '#fafafa',
      black: '#1a1a1a',
      gray: '#dddddd',
      grayer: '#eeeeee'
    },
    fontFamily: {
      zenkaku: ['Zen Kaku Gothic New', 'sans-serif']
    }
  },
  darkMode: 'class',
  plugins: []
}
