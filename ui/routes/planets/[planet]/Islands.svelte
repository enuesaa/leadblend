<script lang="ts">
	import { listIslands } from '$lib/graphql/island'
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet

	const islands = listIslands(planet.id)
</script>

{#if $islands.data !== undefined}
	{#each $islands.data as island}
		<a href={`/planets/${planet.name}/islands/${island.id}`}>{island.title}</a>
	{/each}
{/if}

<style lang="postcss">
	a {
		border: solid 1px #1a1a1a;
		@apply block my-2 p-1 font-zenkaku text-xl font-bold h-10;
	}
</style>
