<script lang="ts">
	import { planets } from '$lib/graphql/planet'
</script>

{#each $planets as planet}
	<li>{planet.id}</li>
{/each}
