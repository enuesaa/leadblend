<script lang="ts">
	import { getPlanet } from '$lib/graphql/planet'
	import CreateIsland from './CreateIsland.svelte'

	export let planetName: string

	const planet = getPlanet(planetName)
</script>

<h3>{planetName}: create island</h3>

{#if $planet.data?.id !== undefined}
	<CreateIsland planet={$planet.data} />
{/if}
