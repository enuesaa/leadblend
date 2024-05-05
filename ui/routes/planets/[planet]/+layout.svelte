<script lang="ts">
	import { page } from '$app/stores'
	import PlanetMenu from './PlanetMenu.svelte'
</script>

{#key $page.params}
	<PlanetMenu planetName={$page.params.planet} />
{/key}

<slot />
