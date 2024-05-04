<script lang="ts">
	import { page } from '$app/stores'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { getPlanet } from '$lib/graphql/planet'
	import Islands from './Islands.svelte'

	const planetName = $page.params.planet
	const planet = getPlanet(planetName)
</script>

<PageTitle title={planetName} />

{#if $planet.data !== undefined}
	<Islands planet={$planet.data} />
{/if}
