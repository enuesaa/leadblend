<script lang="ts">
	import SubTitle from '$lib/components/SubTitle.svelte'
	import { getComet } from '$lib/graphql/comet'
	import DeleteComet from './DeleteComet.svelte'
	import LinkComet from './LinkComet.svelte'

	export let cometId: string
	const comet = getComet(cometId)
</script>

{#if $comet.data !== undefined}
	<SubTitle title="data" />
	<code>{$comet.data?.data ?? ''}</code>

	<SubTitle title="pattern_id" />
	<code>{$comet.data?.patternId ?? ''}</code>
{/if}

<LinkComet {cometId} />
<DeleteComet {cometId} />
