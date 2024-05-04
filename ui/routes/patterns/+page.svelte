<script lang="ts">
	import { listPatterns } from '$lib/graphql/pattern'

	const patterns = listPatterns()
</script>

{#if $patterns.data !== undefined}
	{#each $patterns.data as pattern}
		<div>{pattern.id}</div>
	{/each}
{/if}

<a href='/pattern/create'>Create Pattern</a>
