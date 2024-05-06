<script lang="ts">
	import { type CometObject } from '$lib/comet/types'
	import FieldEditorArray from './FieldEditorArray.svelte'
	import FieldEditorValue from './FieldEditorValue.svelte'

	export let data: CometObject
</script>

{#each data.values as item}
	{#if item.type === 'object'}
		<div>
			{item.key}
			<svelte:self bind:data={item} />
		</div>
	{:else if item.type === 'array'}
		<FieldEditorArray bind:data={item} />
	{:else}
		<FieldEditorValue bind:data={item} />
	{/if}
{/each}

<style lang="postcss">
	div {
		border: solid 1px;
		@apply border-blackgray block w-3/4 p-3 rounded-lg ml-2 my-1;
	}
</style>
