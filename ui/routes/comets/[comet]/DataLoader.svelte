<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
	import { getComet } from '$lib/graphql/comet'
	import DeleteComet from './DeleteComet.svelte'

	export let cometId: string
	const comet = getComet(cometId)
</script>

<PageTitle title={`Comet`} />

{#if $comet.data?.data !== undefined}
	<div>{$comet.data.data}</div>
{/if}

<DeleteComet {cometId} />
