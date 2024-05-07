<script lang="ts">
	import { convertCometData } from '$lib/comet/data'
	import type { CometObject } from '$lib/comet/types'
	import SubTitle from '$lib/components/SubTitle.svelte'
	import { getComet } from '$lib/graphql/comet'
	import FieldEditor from '$lib/components/comet/FieldEditor.svelte'
	import DeleteComet from './DeleteComet.svelte'
	import LinkComet from './LinkComet.svelte'

	export let cometId: string
	const comet = getComet(cometId)
	let data: CometObject = { type: 'object', key: '', values: [] }

	$: [data] = convertCometData($comet.data?.data ?? '{}')
</script>

{#if $comet.data !== undefined}
	<SubTitle title="data" />
	{#if $comet.data?.data !== undefined}
		<FieldEditor {data} />
	{/if}

	<SubTitle title="pattern_id" />
	<code>{$comet.data?.patternId ?? ''}</code>
{/if}

<LinkComet {cometId} />
<DeleteComet {cometId} />
