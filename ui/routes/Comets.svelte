<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'
	import { cometVisible } from '$lib/store/comet'

	const comets = listComets()
</script>

{#if $cometVisible}
	<section>
		{#each $comets.data ?? [] as comet}
			<CometCard {comet} />
		{/each}
		<CometCreateCard />
	</section>
{/if}

<style lang="postcss">
	section {
		box-shadow: 0 0.5px 0.5px rgba(0, 0, 0, 0.3);
		background: rgba(0, 0, 0, 0.1);
		@apply p-3 pt-3 flex gap-3;
	}
</style>
