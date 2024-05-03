<script lang="ts">
	import { useCreateIsland } from '$lib/graphql/island'
	const createPlanet = useCreateIsland()

	export let planetId: string
	let title: string = ''
	let content: string = ''
	let comment: string = ''

	async function hanldeClick() {
		const data = { title, content, planetId, comment }
		await $createPlanet.mutateAsync(data)
	}
</script>

<input type="text" bind:value={title} />
<input type="text" bind:value={content} />
<input type="text" bind:value={comment} />
<button on:click|preventDefault={hanldeClick}>submit</button>
