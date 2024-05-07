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
		@apply w-full bg-gray block m-2 p-2;
		@apply outline-none rounded text-black text-xl font-zenkaku font-bold;
		@apply hover:bg-graywhite;
	}
</style>
