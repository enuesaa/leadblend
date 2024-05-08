<script>
	import { listComets } from '$lib/graphql/comet'
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'
	import { cometVisible } from '$lib/store/comet'

	const comets = listComets()
</script>

{#if $cometVisible}
	<section class="p-3 pt-3 flex gap-3">
		{#each $comets.data ?? [] as comet}
			<CometCard {comet} />
		{/each}
		<CometCreateCard />
	</section>
{:else}
	<section class="p-2" />
{/if}

<style lang="postcss">
	section {
		background-image: radial-gradient(rgba(0, 0, 0, 0.2) 20%, rgba(0, 0, 0, 0) 0%);
		background-size: 6px 6px;
		box-shadow: 0 0 2px rgba(255, 255, 255, 0.3) inset;
	}
</style>
