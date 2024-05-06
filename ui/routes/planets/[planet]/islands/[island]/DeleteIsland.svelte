<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import { useDeleteIsland } from '$lib/graphql/island'
	export let planetName: string
	export let islandId: string

	const deleteIsland = useDeleteIsland()

	async function handleClick() {
		await $deleteIsland.mutateAsync({ id: islandId })
		goto(`/planets/${planetName}`)
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Danger Zone</h6>
	<span class="mr-3">Are you sure to delete this island?</span>
	<Button {handleClick} label="Delete" />
</section>
