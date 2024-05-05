<script lang="ts">
	import { listStones } from '$lib/graphql/stone'
	export let islandId: string

	const stones = listStones(islandId)
</script>

{#if $stones.data !== undefined}
	{#each $stones.data as stone}
		<div>{stone.id}</div>
		<code>{stone.data}</code>
	{/each}
{/if}
