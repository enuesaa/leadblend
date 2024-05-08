<script lang="ts">
	import { page } from '$app/stores'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import DeletePlanet from './DeletePlanet.svelte'
	import RenamePlanet from './RenamePlanet.svelte'
</script>

<PageTitle title="Planet Configuration" />

{#key $page.params}
	<RenamePlanet planetName={$page.params.planet} />
	<DeletePlanet planetName={$page.params.planet} />
{/key}
