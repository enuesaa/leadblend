<script lang="ts">
	import CreateButton from '$lib/components/CreateButton.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateIsland } from '$lib/graphql/island'
	import { goto } from '$app/navigation'
	import type { Planet } from '$lib/graphql/types'

	const createPlanet = useCreateIsland()

	export let planet: Planet
	let title: string = ''
	let content: string = ''
	let comment: string = ''

	async function hanldeClick() {
		const data = { title, content, comment, planetId: planet.id }
		await $createPlanet.mutateAsync(data)
		goto(`/planets/${planet.name}`)
	}
</script>

<TextInput value={title} label='title' />
<TextInput value={content} label='content' />
<TextInput value={comment} label='comment' />
<CreateButton {hanldeClick} />
