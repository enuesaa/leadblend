<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
</script>

{#if !$listPlanets.fetching && $listPlanets.data !== undefined}
	<ul>
	{#each $listPlanets.data.listPlanets as planet}
		<li>{planet.id}</li>
	{/each}
	</ul>
{/if}
