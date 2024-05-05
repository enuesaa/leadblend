<script lang="ts">
	import { listStones } from '$lib/graphql/stone'
	import DeleteStone from './DeleteStone.svelte'
	export let islandId: string
	export let planetName: string

	const stones = listStones(islandId)
</script>

{#each $stones.data ?? [] as stone}
	<div>{stone.id}</div>
	<code>{stone.data}</code>
	<DeleteStone stoneId={stone.id} {islandId} {planetName} />
{/each}
