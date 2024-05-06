<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { convertCometData, type CometObject } from '$lib/comet/data'

	const createComet = useCreateComet()

	let notice: string = ''

	let useJsonEditor = false
	let jsondata: string = '{}'
	let cometdata: CometObject = {type: 'object', values: []}

	function toggleEditor() {
		cometdata = convertCometData(jsondata)
		useJsonEditor = !useJsonEditor
	}

	let fielda: string = ''
	async function handleClick() {
		const data: any = {}
		if (fielda !== '') {
			data.a = fielda
		}
		await $createComet.mutateAsync({ data: JSON.stringify(data) })
		goto('/')
	}
</script>

<PageTitle title="New Comet" />

<div>
	{notice}
</div>

<button on:click|preventDefault={toggleEditor}>toggle editor</button>

{#if useJsonEditor}
	<textarea bind:value={jsondata} />
{:else}
	{#each cometdata.values as field}
		{#if field.type === 'string'}
			<TextInput bind:value={field.value} />
		{:else if field.type === 'number'}
			<input type="number" bind:value={field.value} />
		{:else if field.type === 'boolean'}
			<input type="checkbox" bind:checked={field.value} />
		{/if}
	{/each}
{/if}
<Button {handleClick} label="Create" />
