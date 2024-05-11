import adapter from '@sveltejs/adapter-static'
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

/** @type {import('@sveltejs/kit').Config} */
export default {
	preprocess: vitePreprocess({
		/**
		 * @see https://github.com/sveltejs/vite-plugin-svelte/issues/607#issuecomment-1479111062
		 */
		style: {
			css: {
				postcss: true,
			},
		},
	}),
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
