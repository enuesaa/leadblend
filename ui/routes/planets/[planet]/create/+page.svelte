<script>
	import { getPlanet } from '$lib/graphql/planet'
	import CreateIsland from './CreateIsland.svelte'
	import { useParams } from '$lib/page'

	const { planet: planetName } = useParams()
	const planet = getPlanet(planetName)
</script>

<h3>{planetName}: create island</h3>

{#if $planet.data?.id !== undefined}
	<CreateIsland planetId={$planet.data.id} />	
{/if}
