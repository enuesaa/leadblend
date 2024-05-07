<script lang="ts">
	import type { Comet } from '$lib/graphql/types'

	export let comet: Comet
</script>

<a href={`/comets/${comet.id}`}>
	{comet.data}
</a>

<style lang="postcss">
	a {
		border: solid 1px rgba(255, 255, 255, 0.1);
		@apply p-2 rounded-lg w-40 h-40 bg-grayblack overflow-hidden;
	}
</style>
