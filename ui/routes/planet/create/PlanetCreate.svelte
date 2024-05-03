<script lang="ts">
	import CreateButton from '$lib/components/CreateButton.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreatePlanet } from '$lib/graphql/planet'

	const createPlanet = useCreatePlanet()

	let name: string = ''
	let comment: string = ''

	async function hanldeClick() {
		const data = { name, comment }
		await $createPlanet.mutateAsync(data)
	}
</script>

<TextInput value={name} label='name' />
<TextInput value={comment} label='comment' />
<CreateButton {hanldeClick} />
