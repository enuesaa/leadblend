<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import { useDeletePlanet } from '$lib/graphql/planet'

	export let planetName: string

	const deletePlanet = useDeletePlanet()
	async function handleClick() {
		await $deletePlanet.mutateAsync({ name: planetName })
		goto('/')
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Danger Zone</h6>
	<span class="mr-3">Are you sure to delete this planet?</span>
	<Button {handleClick} label="Delete" />
</section>
