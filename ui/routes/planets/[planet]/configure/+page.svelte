<script lang="ts">
	import { page } from '$app/stores'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import DeletePlanet from './DeletePlanet.svelte'
</script>

<PageTitle title="Planet Configuration" />

{#key $page.params}
	<DeletePlanet planetName={$page.params.planet} />
{/key}
