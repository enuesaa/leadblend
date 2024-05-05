<script lang="ts">
	import SubTitle from '$lib/components/SubTitle.svelte'
	import { listPlanets } from '$lib/graphql/planet'
	import SideMenuLink from './SideMenuLink.svelte'
	import SideMenuPlanetCreateLink from './SideMenuPlanetCreateLink.svelte'

	const planets = listPlanets()
</script>

<SubTitle title="Planet">
	<SideMenuPlanetCreateLink />
</SubTitle>

{#each $planets.data ?? [] as planet}
	<SideMenuLink {planet} />
{/each}
