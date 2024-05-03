<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
	import CreatePlanet from './CreatePlanet.svelte'

	const planets = listPlanets()
</script>

<CreatePlanet />

{#each $planets.data as planet}
	<li>{planet.id}</li>
{/each}
