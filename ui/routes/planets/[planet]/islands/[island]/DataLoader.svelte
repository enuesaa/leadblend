<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { useDeleteIsland } from '$lib/graphql/island'

	export let planetName: string
	export let islandId: string

	const deleteIsland = useDeleteIsland()

	async function handleClick() {
		await $deleteIsland.mutateAsync({ id: islandId })
		goto(`/planets/${planetName}`)
	}
</script>

<PageTitle title={planetName} />
island: {islandId}

<Button {handleClick} label="Delete" />
