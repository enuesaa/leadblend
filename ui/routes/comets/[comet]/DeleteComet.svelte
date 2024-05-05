<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import { useDeleteComet } from '$lib/graphql/comet'

	export let cometId: string
	const deleteComet = useDeleteComet()

	async function handleClick() {
		await $deleteComet.mutateAsync({ id: cometId })
		goto('/')
	}
</script>

<section class="my-6 p-4 bg-grayblack w-fit rounded-lg">
	<h6 class="text-lg font-bold mb-2">Danger Zone</h6>
	<span class="mr-3">Are you sure to delete this comet?</span>
	<Button {handleClick} label="Delete" />
</section>
