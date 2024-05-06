<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import JsonEditor from '$lib/components/comet/JsonEditor.svelte'
	import FieldEditor from '$lib/components/comet/FieldEditor.svelte'
	import { convertCometData } from '$lib/comet/data'
	import { convertCometDataToJson } from '$lib/comet/tojson'
	import { type CometObject } from '$lib/comet/types'
	import SwitchEditorButton from '$lib/components/comet/SwitchEditorButton.svelte'
	import SubTitle from '$lib/components/SubTitle.svelte'

	const createComet = useCreateComet()

	let notice: string = ''
	let data: CometObject = { type: 'object', key: '', values: [] }
	let jsondata: string = '{}'
	let useJsonEditor = false

	function handleSwitchToFieldEditor() {
		;[data, notice] = convertCometData(jsondata)
		useJsonEditor = false
	}
	function handleSwitchToJsonEditor() {
		jsondata = convertCometDataToJson(data)
		useJsonEditor = true
	}

	async function handleClick() {
		jsondata = convertCometDataToJson(data)
		await $createComet.mutateAsync({ data: jsondata })
		goto('/')
	}
</script>

<PageTitle title="New Comet" />

<SubTitle title="Data">
	{#if useJsonEditor}
		<SwitchEditorButton label="Use Field Editor" handle={handleSwitchToFieldEditor} />
	{:else}
		<SwitchEditorButton label="Use JSON Editor" handle={handleSwitchToJsonEditor} />
	{/if}
</SubTitle>

{#if useJsonEditor}
	<JsonEditor bind:data={jsondata} />
{:else}
	<FieldEditor bind:data />
{/if}

<div>
	{notice}
</div>

<section class="mt-5 p-3">
	<Button {handleClick} label="Create" />
</section>
