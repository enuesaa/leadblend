<script lang="ts">
	import type { Planet } from '$lib/graphql/types'

	export let planet: Planet
</script>

<a href={`/planets/${planet.name}`} class="inline-block m-2 p-2 rounded-lg font-zenkaku text-2xl font-bold">
	{planet.name}
</a>
