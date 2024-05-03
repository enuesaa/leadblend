<script lang="ts">
	import { useCreatePlanet } from '$lib/graphql/planet'

	const createPlanet = useCreatePlanet()

	let name: string
	let comment: string

	async function hanldeClick() {
		const data = { name, comment }
		await $createPlanet.mutateAsync(data)
	}
</script>

<input type="text" bind:value={name} />
<input type="text" bind:value={comment} />
<button on:click|preventDefault={hanldeClick}>submit</button>
