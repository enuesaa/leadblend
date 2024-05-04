<script lang="ts">
	import { page } from '$app/stores'
	import DataLoader from './DataLoader.svelte'
</script>

{#key $page.params}
	<DataLoader cometId={$page.params.comet} />
{/key}
