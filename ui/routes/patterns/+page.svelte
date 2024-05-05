<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { listPatterns } from '$lib/graphql/pattern'

	const patterns = listPatterns()
</script>

<PageTitle title="Patterns" />

{#each $patterns.data ?? [] as pattern}
	<div class="m-2">
		<div class="p-2">{pattern.title}</div>
		<code>{JSON.stringify(pattern.traits)}</code>
	</div>
{/each}

<a href="/patterns/new">Create Pattern</a>
