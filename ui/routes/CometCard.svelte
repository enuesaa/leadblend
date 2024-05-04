<script lang="ts">
	import type { Comet } from '$lib/graphql/types'

	export let comet: Comet
</script>

<div>
	{comet.data}
</div>

<style lang="postcss">
	div {
		border: solid 1px #1a1a1a;
		@apply p-2 rounded-lg w-40 h-40;
	}
</style>
