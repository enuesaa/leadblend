<script lang="ts">
	import TextInput from '$lib/components/TextInput.svelte'
	import { type CometValue } from '$lib/comet/types'

	export let data: CometValue
</script>

{#if data.type === 'string'}
	<TextInput bind:value={data.value} label={data.key} />
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
