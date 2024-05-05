<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'

	const comets = listComets()
</script>

<section class="p-2 m-2 flex gap-3">
	{#if $comets.data !== undefined}
		{#each $comets.data as comet}
			<CometCard {comet} />
		{/each}
	{/if}
	<CometCreateCard />
</section>

<style>
	section {
		border-top: solid 1.5px #1a1a1a;
		border-bottom: solid 1.5px #1a1a1a;
	}
</style>
