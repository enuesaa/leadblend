<script lang="ts">
	import { page } from '$app/stores'
	import DataLoader from './DataLoader.svelte'
</script>

{#key $page.params}
	<DataLoader planetName={$page.params.planet} />
{/key}
