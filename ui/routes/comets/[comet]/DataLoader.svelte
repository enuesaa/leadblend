<script lang="ts">
	import { getComet } from '$lib/graphql/comet'
	import DeleteComet from './DeleteComet.svelte'
	import LinkComet from './LinkComet.svelte'

	export let cometId: string
	const comet = getComet(cometId)
</script>

{#if $comet.data !== undefined}
	{$comet.data?.data ?? ''}
{/if}

<LinkComet {cometId} />
<DeleteComet {cometId} />
