<script lang="ts">
	export let handleClick: () => void
	export let label: string = 'Create'
</script>

<button on:click|preventDefault={handleClick}>
	{label}
</button>

<style lang="postcss">
	button {
		border: solid 1px #1a1a1a;
		@apply rounded-lg p-2 m-1 ml-3;
	}
</style>
