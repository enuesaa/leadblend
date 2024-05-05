import adapter from '@sveltejs/adapter-static'
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

/** @type {import('@sveltejs/kit').Config} */
const config = {
	preprocess: vitePreprocess(),
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
