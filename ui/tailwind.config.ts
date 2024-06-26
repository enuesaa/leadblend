import { type Config } from 'tailwindcss'

export default {
	content: [
		'routes/**/*.svelte',
		'lib/**/*.svelte',
		'app.html',
	],
	theme: {
		colors: {
			black: '#1a1a1a',
			blackgrayer: '#212121',
			blackgray: '#2a2a2a',
			gray: '#999999',
			graywhite: '#aaaaaa',
			grayblack: '#3a3a3a',
			white: '#cccccc',
		},
		fontFamily: {
			zenkaku: ['Zen Kaku Gothic New', 'sans-serif'],
		},
	},
} satisfies Config
