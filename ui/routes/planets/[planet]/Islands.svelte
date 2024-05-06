<script lang="ts">
	import { listIslands } from '$lib/graphql/island'
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet

	const islands = listIslands(planet.id)
</script>

<section class="pt-3">
	{#each $islands.data ?? [] as island}
		<a href={`/planets/${planet.name}/islands/${island.id}`}>{island.title}</a>
	{/each}
</section>

<style lang="postcss">
	a {
		border-bottom: solid 1px;
		@apply block my-3 p-2 ml-2 font-zenkaku text-xl font-bold border-blackgray;
		@apply hover:bg-grayblack;
	}
</style>
