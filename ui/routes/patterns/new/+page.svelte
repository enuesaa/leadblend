<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { useCreatePattern } from '$lib/graphql/pattern'
	import type { TraitInput } from '$lib/graphql/types'
	import TraitForm from './TraitForm.svelte'

	const createPattern = useCreatePattern()

	let title: string = ''
	let traits: TraitInput[] = []

	function handleAddTrait() {
		traits = [...traits, {
			defaultValue: '',
			path: '$',
			required: true,
			type: 'string',
		}]
	}

	async function handleClick() {
		await $createPattern.mutateAsync({ title, traits, color: '#1a1a1a' })
		goto('/')
	}
</script>

<PageTitle title="New Pattern" />
<TextInput bind:value={title} label="title" />

<div class="m-3">
	{#each traits.flat() as trait, i}
		<TraitForm bind:trait={trait} remove={() => traits = traits.splice(i, 1)} />
	{/each}
	<Button handleClick={handleAddTrait} label='add trait' />
</div>

<Button {handleClick} label="Create" />
