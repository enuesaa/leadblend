<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'
	import { cometVisible, showComet } from '$lib/store/comet'

	const comets = listComets()
</script>

{#if $cometVisible}
	<section>
		{#each $comets.data ?? [] as comet}
			<CometCard {comet} />
		{/each}
		<CometCreateCard />
	</section>
{:else}
	<section on:click|stopPropagation={showComet} />
{/if}

<style lang="postcss">
	section {
		box-shadow: 0 0.5px 0.5px rgba(0, 0, 0, 0.3);
		@apply p-3 pt-3 flex gap-3 bg-grayblack;
	}
</style>
