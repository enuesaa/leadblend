export default {
	printWidth: 120,
	tabWidth: 2,
	semi: false,
	singleQuote: true,
	endOfLine: 'lf',
	trailingComma: 'none',
	plugins: ['prettier-plugin-svelte'],
	useTabs: true,
	overrides: [
		{
			files: '*.svelte',
			options: {
				parser: 'svelte'
			}
		}
	]
}
