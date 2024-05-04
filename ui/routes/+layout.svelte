<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
	import Comets from './Comets.svelte'
	import SideMenu from './SideMenu.svelte'

	const queryClient = new QueryClient({
		defaultOptions: {
			queries: {
				staleTime: 0
			}
		}
	})
</script>

<QueryClientProvider client={queryClient}>
	<Header />
	<Comets />
	<main class="container mx-auto flex gap-9">
		<div class="flex-none w-28">
			<SideMenu />
		</div>
		<div class="flex-auto">
			<slot />
		</div>
	</main>
</QueryClientProvider>
