<script lang="ts">
	import { page } from '$app/stores'

	let planetName: string, islandId: string
	$: planetName = $page.params.planet
	$: islandId = $page.params.island
</script>

planet: {planetName} <br />
island: {islandId}
