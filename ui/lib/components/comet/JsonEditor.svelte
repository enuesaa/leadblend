<script lang="ts">
	export let data: string = '{}'
</script>

<textarea bind:value={data} />

<style lang="postcss">
	textarea {
		@apply w-3/4 min-h-52;
		@apply bg-gray block m-2 p-2 font-normal rounded text-black;
		@apply outline-none;
	}
</style>
