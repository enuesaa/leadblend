<script lang="ts">
	export let value: string = ''
	export let placeholder: string = ''
	export let label: string = 'テキスト'
</script>

<label class="w-11/12 block my-3 mt-4 font-bold">
	<span class="pl-2">{label}</span>
	<input bind:value {placeholder} />
</label>

<style lang="postcss">
	input {
		@apply w-full bg-grayer block m-2 p-2 font-normal rounded text-black;
		@apply outline-none;
	}
</style>
