<script lang="ts">
	import { getCurrentSpace } from '$lib/graphql/space'
	import HeaderCometToggle from './HeaderCometToggle.svelte'

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
		<HeaderCometToggle />
		<a href="/patterns" class="inline-block w-20 font-zenkaku text-lg font-bold text-center">Patterns</a>
	</div>
</header>

<style lang="postcss">
	header {
		box-shadow: 0 0.5px 0.5px rgba(0,0,0,0.3);
		@apply mx-auto px-8 py-1;
	}
</style>
