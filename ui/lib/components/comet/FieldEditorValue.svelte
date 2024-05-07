<script lang="ts">
	import TextInput from '$lib/components/TextInput.svelte'
	import { type CometValue } from '$lib/comet/types'

	export let data: CometValue
</script>

{#if data.type === 'string'}
	<label>
		{data.key}
		<input type="text" bind:value={data.value} />
	</label>
{:else if data.type === 'number'}
	<label>
		{data.key}
		<input type="number" bind:value={data.value} />
	</label>
{:else if data.type === 'boolean'}
	<label>
		{data.key}
		<input type="checkbox" bind:checked={data.value} />
	</label>
{/if}

<style lang="postcss">
	label {
		border: solid 1px;
		@apply border-blackgray block w-3/4 p-3 rounded-lg ml-2 my-1;
	}
	input {
		@apply bg-gray block mt-1 p-2 font-normal rounded text-black;
		@apply outline-none;
	}
</style>
