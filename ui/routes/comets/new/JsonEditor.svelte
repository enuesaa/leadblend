<script lang="ts">
	import { convertCometData } from '$lib/comet/data'
	import { type CometObject } from '$lib/comet/types'

	export let notice: string
	export let switchEditor: () => void

	let useJsonEditor = false
	let jsondata: string = '{}'
	let cometdata: CometObject = {type: 'object', key: '', values: []}

	function handleConvertToFieldEditor() {
		[cometdata, notice] = convertCometData(jsondata)
		switchEditor()
	}
</script>

<button on:click|preventDefault={handleConvertToFieldEditor}>use field editor</button>
<textarea bind:value={jsondata} />
