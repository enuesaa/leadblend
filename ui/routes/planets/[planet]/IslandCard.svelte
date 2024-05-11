<script lang="ts">
	import type { Island } from '$lib/graphql/types'
	import { listStones } from '$lib/graphql/stone'

	export let island: Island
	export let planetName: string

	const stones = listStones(island.id)
</script>

<a href={`/planets/${planetName}/islands/${island.id}`}>
	{island.title}

	<div class="absolute bottom-0 right-0">
		{#each $stones.data ?? [] as _}
			<span class="inline-block w-5 h-5 bg-gray mr-1" />
		{/each}
	</div>
</a>

<style lang="postcss">
	a {
		border-bottom: solid 1px;
		@apply block my-3 p-2 ml-2 font-zenkaku text-xl font-bold border-blackgray;
		@apply relative;
		@apply hover:bg-grayblack;
	}
</style>
