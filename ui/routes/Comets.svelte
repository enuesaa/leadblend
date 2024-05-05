<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'

	const comets = listComets()
</script>

<section class="p-3 mt-2 pt-1 flex gap-3">
	{#if $comets.data !== undefined}
		{#each $comets.data as comet}
			<CometCard {comet} />
		{/each}
	{/if}
	<CometCreateCard />
</section>

<style>
	section {
		box-shadow: 0 1px 1px rgba(0,0,0,0.3);
	}
</style>
