<script lang="ts">
	import { planets } from '$lib/graphql/planet'
	import CreatePlanet from './CreatePlanet.svelte'
</script>

<CreatePlanet />

{#each $planets as planet}
	<li>{planet.id}</li>
{/each}
