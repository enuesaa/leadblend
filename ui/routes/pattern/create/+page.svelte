<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { useCreatePattern } from '$lib/graphql/pattern'

	const createPattern = useCreatePattern()

	let title: string = ''

	async function handleClick() {
		await $createPattern.mutateAsync({ title })
		goto('/')
	}
</script>

<PageTitle title="create pattern" />
a
<TextInput bind:value={title} label="title" />
<Button {handleClick} label="Create" />
