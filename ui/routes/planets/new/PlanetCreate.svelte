<script lang="ts">
	import { goto } from '$app/navigation'
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreatePlanet } from '$lib/graphql/planet'

	const createPlanet = useCreatePlanet()

	let name: string = ''
	let comment: string = ''

	async function handleClick() {
		const data = { name, comment }
		await $createPlanet.mutateAsync(data)
		goto('/')
	}
</script>

<TextInput bind:value={name} label="name" />
<TextInput bind:value={comment} label="comment" />
<Button {handleClick} label="Create" />
