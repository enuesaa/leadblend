<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
	import SideMenuLink from './SideMenuLink.svelte'
	import SideMenuPlanetCreateLink from './SideMenuPlanetCreateLink.svelte'

	const planets = listPlanets()
</script>

{#each $planets.data as planet}
	<SideMenuLink {planet} />
{/each}
<SideMenuPlanetCreateLink />
