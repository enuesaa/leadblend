<script lang="ts">
	import type { Comet } from '$lib/graphql/types'

	export let comet: Comet
</script>

<a href={`/comets/${comet.id}`}>
	{comet.data}
</a>

<style lang="postcss">
	a {
		border: solid 1px #1a1a1a;
		@apply p-2 rounded-lg w-40 h-40;
	}
</style>
