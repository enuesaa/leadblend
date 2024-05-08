<script lang="ts">
	import Layout from './Layout.svelte'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { listPlanets } from '$lib/graphql/planet'
	import SideMenuLink from './SideMenuLink.svelte'

	const planets = listPlanets()
</script>

<Layout sidebar={false}>
	<PageTitle title="Planet" />

	{#each $planets.data ?? [] as planet}
		<SideMenuLink {planet} />
	{/each}
</Layout>
