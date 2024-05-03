<script lang="ts">
	import { listPlanets } from '$lib/graphql/planet'
	import SideMenuLink from './SideMenuLink.svelte'

	const planets = listPlanets()
</script>

{#each $planets.data as planet}
	<SideMenuLink planet={planet} />
{/each}
