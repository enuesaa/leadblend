/** @type {import('tailwindcss').Config} */
export default {
	content: ['./{routes,lib}/**/*.{html,js,svelte,ts}', 'app.html'],
	theme: {
		colors: {
			white: '#fafafa',
			black: '#1a1a1a',
			blackgray: '#3a3a3a',
			gray: '#cccccc',
			graywhite: '#dddddd',
			grayblack: '#aaaaaa',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
	darkMode: 'class',
	plugins: [],
}
