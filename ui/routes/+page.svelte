<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
	import CreatePlanet from './CreatePlanet.svelte'
	import PlanetLink from './PlanetLink.svelte'

	const planets = listPlanets()
</script>

<div>
	{#each $planets.data as planet}
		<PlanetLink planet={planet} />
	{/each}
</div>

<CreatePlanet />
