<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import JsonEditor from './JsonEditor.svelte'
	import FieldEditor from './FieldEditor.svelte'

	const createComet = useCreateComet()

	let notice: string = ''
	let useJsonEditor = false

	async function handleClick() {
		await $createComet.mutateAsync({ data: '{}' })
		goto('/')
	}
</script>

<PageTitle title="New Comet" />

<div>{notice}</div>

{#if useJsonEditor}
	<JsonEditor bind:notice={notice} switchEditor={() => useJsonEditor = false } />
{:else}
	<FieldEditor bind:notice={notice} switchEditor={() => useJsonEditor = true } />
{/if}

<Button {handleClick} label="Create" />
