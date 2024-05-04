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

<Button {handleClick} label="Delete" />
