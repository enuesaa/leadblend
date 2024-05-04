<script lang="ts">
	import { page } from '$app/stores'

	let planetName = $page.params.planet
	let islandId = $page.params.island
	$: planetName = $page.params.planet
	$: islandId = $page.params.island
</script>

planet: {planetName} <br />
island: {islandId}
