<script lang="ts">
	import { page } from '$app/stores'
	import DataLoader from './DataLoader.svelte'

	let planetName = $page.params.planet
	$: planetName = $page.params.planet
</script>

<DataLoader {planetName} />
