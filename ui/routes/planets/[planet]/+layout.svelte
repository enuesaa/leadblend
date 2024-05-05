<script lang="ts">
	import { page } from '$app/stores'
	import LayoutPlanetMenu from './LayoutPlanetMenu.svelte'
</script>

{#key $page.params}
	<LayoutPlanetMenu planetName={$page.params.planet} />
{/key}

<div class="w-10/12">
	<slot />
</div>
