<script lang="ts">
	export let handleClick: () => void
	export let label: string = 'Create'
</script>

<button on:click|preventDefault={handleClick}>
	{label}
</button>

<style lang="postcss">
	button {
		@apply rounded-lg p-2 m-1 ml-3 bg-graywhite text-black font-semibold;
		@apply hover:bg-gray;
	}
</style>
