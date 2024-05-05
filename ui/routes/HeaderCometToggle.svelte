<script lang="ts">
	import { cometVisible, showComet, hideComet } from '$lib/store/comet'
</script>

{#if $cometVisible}
	<button on:click|preventDefault={hideComet}>Comet</button>
{:else}
	<button on:click|preventDefault={showComet}>Comet</button>
{/if}

<style lang="postcss">
	button {
		@apply h-full inline-block w-20 font-zenkaku text-lg font-bold text-center;
	}
</style>
