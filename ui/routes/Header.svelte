<script lang="ts">
	import { getCurrentSpace } from '$lib/graphql/space'
	import { cometVisible, showComet, hideComet } from '$lib/store/comet'

	const space = getCurrentSpace()
</script>

<header>
	<div class="flex">
		<a href="/" class="w-52 text-center">
			<h1 class="font-zenkaku text-2xl font-bold">
				{$space.data?.name}
			</h1>
		</a>

		<div class="flex-auto"></div>
		<a href="/patterns" class="w-32 font-zenkaku text-lg font-bold">
			Patterns
		</a>
		{#if $cometVisible}
			<button on:click|preventDefault={hideComet} class="rounded-lg">Comet</button>
		{:else}
			<button on:click|preventDefault={showComet} class="rounded-lg">Comet</button>
		{/if}
	</div>
</header>

<style lang="postcss">
	header {
		box-shadow: 0 0.5px 0.5px rgba(0,0,0,0.3);
		@apply mx-auto px-8 py-1;
	}
</style>
