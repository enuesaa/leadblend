<script lang="ts">
	import { listPatterns } from '$lib/graphql/pattern'

	const patterns = listPatterns()
</script>

{#if $patterns.data !== undefined}
	{#each $patterns.data as pattern}
	<div class="m-2">
		<div class="p-2">{pattern.title}</div>
		<code>{JSON.stringify(pattern.traits)}</code>
	</div>
	{/each}
{/if}

<a href='/pattern/create'>Create Pattern</a>
