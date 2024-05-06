<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { convertCometData, type CometObject } from '$lib/comet/data'
	import { convertCometDataToJson } from '$lib/comet/tojson'

	const createComet = useCreateComet()

	let notice: string = ''

	let useJsonEditor = false
	let jsondata: string = '{}'
	let cometdata: CometObject = {type: 'object', key: '', values: []}

	function handleConvertToFieldEditor() {
		[cometdata, notice] = convertCometData(jsondata)
		useJsonEditor = false
	}

	function handleConvertToJsonEditor() {
		notice = ''
		jsondata = convertCometDataToJson(cometdata)
		useJsonEditor = true
	}

	async function handleClick() {
		await $createComet.mutateAsync({ data: jsondata })
		goto('/')
	}
</script>

<PageTitle title="New Comet" />

<div>
	{notice}
</div>

{#if useJsonEditor}
	<button on:click|preventDefault={handleConvertToFieldEditor}>use field editor</button>
	<textarea bind:value={jsondata} />
{:else}
	<button on:click|preventDefault={handleConvertToJsonEditor}>use json editor</button>
	{#each cometdata.values as field}
		{#if field.type === 'string'}
			<TextInput bind:value={field.value} label={field.key} />
		{:else if field.type === 'number'}
			<label>
				{field.key}
				<input type="number" bind:value={field.value} />
			</label>
		{:else if field.type === 'boolean'}
			<label>
				{field.key}
				<input type="checkbox" bind:checked={field.value} />
			</label>
		{/if}
	{/each}
{/if}
<Button {handleClick} label="Create" />
