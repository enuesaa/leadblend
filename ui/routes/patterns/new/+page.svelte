<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { useCreatePattern } from '$lib/graphql/pattern'
	import type { TraitInput } from '$lib/graphql/types'

	const createPattern = useCreatePattern()

	let title: string = ''

	async function handleClick() {
		const traits = [
			{ path: '$.a', type: 'string', defaultValue: 'a', required: true },
			{ path: '$.b', type: 'string', defaultValue: 'b', required: true },
			{ path: '$.c', type: 'string', defaultValue: 'c', required: true },
		] satisfies TraitInput[]
		await $createPattern.mutateAsync({ title, traits })
		goto('/')
	}
</script>

<PageTitle title="New Pattern" />
<TextInput bind:value={title} label="title" />
<Button {handleClick} label="Create" />
