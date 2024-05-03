<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
	import CreatePlanet from './CreatePlanet.svelte'
	import { get } from 'svelte/store'

	const planets = listPlanets()
	console.log(get(planets))
</script>

<CreatePlanet />

{#each $planets as planet}
	<li>{planet.id}</li>
{/each}
