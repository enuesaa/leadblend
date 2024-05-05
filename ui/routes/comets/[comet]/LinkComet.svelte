<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useLinkComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	export let cometId: string

	const linkComet = useLinkComet()

	let islandId: string = ''

	async function handleClick() {
		await $linkComet.mutateAsync({ id: cometId, islandId })
		goto('/')
	}
</script>

<TextInput bind:value={islandId} label='Island Id' />
<Button {handleClick} label='Link' />
