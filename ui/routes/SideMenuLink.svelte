<script lang="ts">
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet
</script>

<a href={`/planets/${planet.name}`}>
	{planet.name}
</a>

<style lang="postcss">
	a {
		@apply block my-1 p-1 font-zenkaku text-xl font-bold h-10;
		@apply ml-2;
	}
</style>
