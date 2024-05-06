<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import { useDeleteStone } from '$lib/graphql/stone'
	import type { Stone } from '$lib/graphql/types'

	export let stone: Stone
	export let planetName: string
	export let islandId: string

	const deleteStone = useDeleteStone()

	async function handleClick() {
		await $deleteStone.mutateAsync({ id: stone.id })
		goto(`/planets/${planetName}/islands/${islandId}`)
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Stone</h6>
	<div>{stone.id}</div>
	<code class="block">{stone.data}</code>
	<Button {handleClick} label="Delete" />
</section>
