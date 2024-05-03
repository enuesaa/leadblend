<script>
	import { page } from '$app/stores'
	import { get } from 'svelte/store'

	const planetName = get(page).params.planet
</script>
