<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import { useDeleteStone } from '$lib/graphql/stone'
	export let stoneId: string
	export let planetName: string
	export let islandId: string

	const deleteStone = useDeleteStone()

	async function handleClick() {
		await $deleteStone.mutateAsync({ id: stoneId })
		goto(`/planets/${planetName}/islands/${islandId}`)
	}
</script>

<Button {handleClick} label="Delete" />
