<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { getPlanet } from '$lib/graphql/planet'
	import Islands from './Islands.svelte'

	export let planetName: string
	const planet = getPlanet(planetName)
</script>

<PageTitle title={planetName} />

{#if $planet.data !== undefined}
	<Islands planet={$planet.data} />
{/if}
