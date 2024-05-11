import { defineConfig } from 'vitest/config'
import { sveltekit } from '@sveltejs/kit/vite'
import path from 'node:path'

export default defineConfig({
	plugins: [sveltekit()],
	resolve: {
		alias: {
			$lib: path.join(__dirname, './lib'),
		},
	},
	optimizeDeps: {
		exclude: ['@urql/svelte'],
	},
	test: {
		include: ['lib/**/*.test.ts'],
	},
})
