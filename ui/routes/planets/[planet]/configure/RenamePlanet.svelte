<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useRenamePlanet } from '$lib/graphql/planet'

	export let planetName: string
	let newPlanetName: string

	const renamePlanet = useRenamePlanet()
	async function handleClick() {
		await $renamePlanet.mutateAsync({ name: planetName, newName: newPlanetName })
		goto('/')
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Rename Planet</h6>
	<TextInput bind:value={newPlanetName} label="New Name" />
	<Button {handleClick} label="Rename" />
</section>
