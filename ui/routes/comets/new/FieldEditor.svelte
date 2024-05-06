<script lang="ts">
	import TextInput from '$lib/components/TextInput.svelte'
	import { type CometObject } from '$lib/comet/types'
	import { convertCometDataToJson } from '$lib/comet/tojson'

	export let notice: string
	export let switchEditor: () => void

	let jsondata: string = '{}'
	let cometdata: CometObject = {type: 'object', key: '', values: []}

	function handleConvertToJsonEditor() {
		notice = ''
		jsondata = convertCometDataToJson(cometdata)
		switchEditor()
	}
</script>

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
