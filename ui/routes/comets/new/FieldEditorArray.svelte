<script lang="ts">
	import { type CometArray } from '$lib/comet/types'
	import FieldEditor from './FieldEditor.svelte'
	import FieldEditorValue from './FieldEditorValue.svelte'

	export let data: CometArray
</script>

<div>
	{data.key}
	{#each data.values as item}
		{#if item.type === 'object'}
			<FieldEditor bind:data={item} />
		{:else if item.type === 'array'}
			<svelte:self bind:data={item} />
		{:else}
			<FieldEditorValue bind:data={item} />
		{/if}
	{/each}
</div>

<style lang="postcss">
	div {
		border: solid 1px;
		@apply border-blackgray block w-3/4 p-3 rounded-lg ml-2 my-1;
	}
</style>
