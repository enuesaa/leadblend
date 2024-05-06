<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'

	const createComet = useCreateComet()

	let notice: string = ''

	let useJsonEditor = false
	let jsondata: string = '{}'
	type Field = {path: string, value: any}
	let fields: Field[] = [] // convert from jsondata

	function toggleEditor() {
		try {
			const parseddata = JSON.parse(jsondata)
			notice = ''
			for (const [key, value] of Object.entries(parseddata)) {
				fields = [...fields, {
					path: `$.${key}`,
					value,
				}]
			}
		} catch (err) {
			notice = 'invalid json format'
			jsondata = '{}'
		}
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
	{#each fields as field}
		<TextInput bind:value={field.value} label={field.path} />		
	{/each}
{/if}
<Button {handleClick} label="Create" />
