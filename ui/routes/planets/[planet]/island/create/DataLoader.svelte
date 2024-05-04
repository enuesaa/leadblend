<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
import { getPlanet } from '$lib/graphql/planet'
	import CreateIsland from './CreateIsland.svelte'

	export let planetName: string

	const planet = getPlanet(planetName)
</script>

<PageTitle title={`${planetName}: create island`} />

{#if $planet.data?.id !== undefined}
	<CreateIsland planet={$planet.data} />
{/if}
