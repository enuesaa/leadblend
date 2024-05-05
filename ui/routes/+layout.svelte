<script lang="ts">
	import './app.css'
	import { QueryClientProvider } from '@tanstack/svelte-query'
	import Header from './Header.svelte'
	import Comets from './Comets.svelte'
	import LayoutMain from './LayoutMain.svelte'
</script>

<QueryClientProvider>
	<Header />
	<Comets />
	<LayoutMain>
		<slot />
	</LayoutMain>
</QueryClientProvider>
