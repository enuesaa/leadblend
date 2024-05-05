<script lang="ts">
	import { page } from '$app/stores'
	import PageTitle from '$lib/components/PageTitle.svelte'
	import DataLoader from './DataLoader.svelte'
</script>

<PageTitle title={`Comet`} />

{#key $page.params}
	<DataLoader cometId={$page.params.comet} />
{/key}
