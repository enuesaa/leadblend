<script lang="ts">
	import { page } from '$app/stores'
	import { getPlanet } from '$lib/graphql/planet'
	import CreateIsland from './CreateIsland.svelte'

	let planetName: string
	$: planetName = $page.params.planet

	const planet = getPlanet(planetName)
</script>

<h3>{planetName}: create island</h3>

{#if $planet.data?.id !== undefined}
	<CreateIsland planetId={$planet.data.id} />
{/if}
