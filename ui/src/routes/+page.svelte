<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
</script>

{#if $listPlanets.fetching}
	<p>Loading...</p>
{:else if $listPlanets.error}
	<p>Oh no... {$listPlanets.error.message}</p>
{:else}
	<ul>
	{#each $listPlanets.data.listPlanets as planet}
		<li>{planet.id}</li>
	{/each}
	</ul>
{/if}
