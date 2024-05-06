<script lang="ts">
	import { listStones } from '$lib/graphql/stone'
	import Stone from './Stone.svelte'
	export let islandId: string
	export let planetName: string

	const stones = listStones(islandId)
</script>

{#each $stones.data ?? [] as stone}
	<Stone {stone} {planetName} {islandId} />
{/each}
