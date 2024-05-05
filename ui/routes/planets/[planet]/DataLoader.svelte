<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { getPlanet } from '$lib/graphql/planet'
	import Icon from '@iconify/svelte'
	import Islands from './Islands.svelte'

	export let planetName: string
	let planet = getPlanet(planetName)
</script>

<PageTitle title={planetName} />

{#if $planet.data !== undefined}
	<Islands planet={$planet.data} />
	<a href={`/planets/${$planet.data.name}/island/create`}>
		<Icon icon="fa-solid:plus" />
	</a>
{/if}
