<script>
	import { page } from '$app/stores'
	import { get } from 'svelte/store'

	const params = get(page).params
	const planetName = params.planet
	const islandId = params.island
</script>

planet: {planetName} <br />
island: {islandId}
