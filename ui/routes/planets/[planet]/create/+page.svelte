<script>
	import { page } from '$app/stores'
	import { get } from 'svelte/store'
	import CreateIsland from './CreateIsland.svelte'

	const planetName = get(page).params.planet
</script>

<h3>{planetName}: create island</h3>
<CreateIsland planetId={''} />
