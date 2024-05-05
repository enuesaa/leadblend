<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateIsland } from '$lib/graphql/island'
	import { goto } from '$app/navigation'
	import type { Planet } from '$lib/graphql/types'

	const createPlanet = useCreateIsland()

	export let planet: Planet
	let title: string = ''
	let content: string = ''
	let comment: string = ''

	async function handleClick() {
		const data = { title, content, comment, planetId: planet.id }
		await $createPlanet.mutateAsync(data)
		goto(`/planets/${planet.name}`)
	}
</script>

<TextInput bind:value={title} label="title" />
<TextInput bind:value={content} label="content" />
<TextInput bind:value={comment} label="comment" />
<Button {handleClick} label="Create" />
