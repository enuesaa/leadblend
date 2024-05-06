<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import SubTitle from '$lib/components/SubTitle.svelte'
	import { useDeleteIsland } from '$lib/graphql/island'
	import DeleteIsland from './DeleteIsland.svelte'
	import Stones from './Stones.svelte'

	export let planetName: string
	export let islandId: string
</script>

<PageTitle title={planetName} />

<SubTitle title="IslandId" />
{islandId}

<Stones {planetName} {islandId} />
<DeleteIsland {planetName} {islandId} />
