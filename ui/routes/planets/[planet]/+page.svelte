<script>
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { getPlanet } from '$lib/graphql/planet'
	import { useParams } from '$lib/page'
	import Islands from './Islands.svelte'

	const { planet: planetName } = useParams()
	const planet = getPlanet(planetName)
</script>

<PageTitle title={planetName} />

{#if $planet.data !== undefined}
	<Islands planet={$planet.data} />
{/if}
