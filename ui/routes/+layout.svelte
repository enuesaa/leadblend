<script lang="ts">
	import './app.css'
	import { QueryClient, QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
	import SideMenu from './SideMenu.svelte'
	import Comets from './Comets.svelte'

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
	<main>
		<div class="left">
			<SideMenu />
		</div>
		<div class="right">
			<slot />
		</div>
	</main>
</QueryClientProvider>

<style lang="postcss">
	main {
		min-height: calc(100vh - 200px);
		@apply mx-auto flex gap-5;
	}
	.left {
		@apply flex-none w-48 px-5 pt-1 bg-blackgray text-gray;
	}
	.right {
		@apply flex-auto pt-1 pl-1;
	}
</style>
