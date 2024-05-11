<script lang="ts">
	import type { TraitInput } from '$lib/graphql/types'
	import TextInput from '$lib/components/TextInput.svelte'
	import Button from '$lib/components/Button.svelte'

	export let trait: TraitInput
	export let remove: () => void
</script>

<section>
	<TextInput bind:value={trait.path} label="path" />
	<label>
		type
		<select bind:value={trait.type}>
			<option value="string">string</option>
			<option value="boolean">boolean</option>
			<option value="number">number</option>
			<option value="array">array</option>
			<option value="object">object</option>
			<option value="null">null</option>
		</select>
	</label>
	<TextInput bind:value={trait.defaultValue} label="defaultValue" />

	<label>
		required
		<input type="checkbox" bind:value={trait.required} />
	</label>

	<div>
		<Button handleClick={remove} label="remove" />
	</div>
</section>

<style lang="postcss">
	section {
		border: solid 1px;
		@apply border-gray px-5 py-2;
	}
	select {
		@apply w-28 block m-2 p-2 font-normal rounded text-black outline-none;
	}
</style>
