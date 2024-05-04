<script lang="ts">
	import { listIslands } from '$lib/graphql/island'
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet

	const islands = listIslands(planet.id)
</script>

{#each $islands.data as island}
	<div>{island.id}</div>
{/each}

<a href={`/planets/${planet.name}/create`}>+</a>
