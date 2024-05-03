<script>
	import { page } from '$app/stores'
	import { get } from 'svelte/store'
	import { getPlanet } from '$lib/graphql/planet'
	import CreateIsland from './CreateIsland.svelte'

	const planetName = get(page).params.planet
	const planet = getPlanet(planetName)
</script>

<h3>{planetName}: create island</h3>

{#if $planet.data?.id !== undefined}
	<CreateIsland planetId={$planet.data.id} />	
{/if}
