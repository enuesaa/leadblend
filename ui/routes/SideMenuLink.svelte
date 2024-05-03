<script lang="ts">
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet
</script>

<a href={`/planets/${planet.name}`} class="block m-1 px-1 font-zenkaku text-xl font-bold">
	{planet.name}
</a>
