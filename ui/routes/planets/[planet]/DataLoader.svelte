<script lang="ts">
	import { getPlanet } from '$lib/graphql/planet'
	import DeletePlanet from './DeletePlanet.svelte'
	import Islands from './Islands.svelte'

	export let planetName: string
	let planet = getPlanet(planetName)
</script>

{#if $planet.data !== undefined}
	<Islands planet={$planet.data} />
	<DeletePlanet {planetName} />
{/if}
