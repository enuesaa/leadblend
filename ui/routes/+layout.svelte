<script lang="ts">
	import './app.css'
	import { QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
	import Comets from './Comets.svelte'
	import SideMenu from './SideMenu.svelte'
</script>

<QueryClientProvider>
	<Header />
	<Comets />
	<main class="container mx-auto flex">
		<div class="flex-none w-20">
			<SideMenu />
		</div>
		<div class="flex-auto">
			<slot />
		</div>
	</main>
</QueryClientProvider>
