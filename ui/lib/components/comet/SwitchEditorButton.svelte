<script lang="ts">
	export let label: string
	export let handle: () => void
</script>

<button on:click|preventDefault={handle}>
	{label}
</button>

<style lang="postcss">
	button {
		border: solid 1px;
		@apply p-1 px-2 border-blackgray rounded-lg ml-3 mb-5 inline-block;
		@apply text-base font-zenkaku font-bold;
	}
</style>
