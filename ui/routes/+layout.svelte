<script lang="ts">
	import './app.css'
	import { QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
</script>

<QueryClientProvider>
	<Header />
	<slot />
</QueryClientProvider>
