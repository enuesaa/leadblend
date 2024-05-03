<script lang="ts">
	export let hanldeClick: () => void
</script>

<button on:click|preventDefault={hanldeClick}>
	Create
</button>

<style lang="postcss">
	button {
		border: solid 1px #1a1a1a;
		@apply rounded-lg p-2 m-1 ml-3;
	}
</style>
