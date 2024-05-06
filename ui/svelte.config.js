import adapter from '@sveltejs/adapter-static'
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: [
		// https://github.com/sveltejs/vite-plugin-svelte/issues/607#issuecomment-1479111062
		vitePreprocess({
			style: {
				css: {
					postcss: true,
				},
			},
		}),
	],
	kit: {
		adapter: adapter({
			fallback: 'index.html',
			pages: 'dist',
		}),
		files: {
			lib: 'lib',
			routes: 'routes',
			appTemplate: 'app.html',
		},
	},
}

export default config
