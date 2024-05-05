<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import { useLinkComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import { listPlanets } from '$lib/graphql/planet'
	import { listIslands } from '$lib/graphql/island'
	export let cometId: string

	const linkComet = useLinkComet()

	let planetId: string = ''
	const planets = listPlanets()

	let islandId: string = ''
	let islands = listIslands(planetId)
	$: islands = listIslands(planetId)

	async function handleClick() {
		await $linkComet.mutateAsync({ id: cometId, islandId })
		goto('/')
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Link Island</h6>
	<label>
		<span>Planet</span>
		<select bind:value={planetId}>
			{#each $planets.data ?? [] as planet}
				<option value={planet.id}>{planet.name}</option>
			{/each}
		</select>
	</label>
	<label>
		<span>Island</span>
		<select bind:value={islandId}>
			{#each $islands.data ?? [] as island}
				<option value={island.id}>{island.title}</option>
			{/each}
		</select>
	</label>
	<Button {handleClick} label="OK" />	
</section>

<style lang="postcss">
	label {
		@apply inline-block mx-1 py-1;
	}
	label span {
		@apply ml-2;
	}
	select {
		@apply w-28 bg-graywhite block m-2 p-2 font-normal rounded text-black outline-none;
	}
</style>
