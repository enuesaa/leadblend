<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'

	const comets = listComets()
</script>

<section class="p-2 m-2 flex gap-3">
	{#each $comets.data as comet}
		<CometCard {comet} />
	{/each}
	<CometCreateCard />
</section>

<style>
	section {
		border-top: solid 1px #1a1a1a;
		border-bottom: solid 1px #1a1a1a;
	}
</style>
