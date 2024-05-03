<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
</script>


<QueryClientProvider>
	<Header />
	<main class="container mx-auto">
		<slot />
	</main>
</QueryClientProvider>
