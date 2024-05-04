<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'

	const createComet = useCreateComet()

	let data: string = ''

	async function handleClick() {
		await $createComet.mutateAsync({data})
		goto('/')
	}
</script>

<PageTitle title="create comet" />

<TextInput bind:value={data} label='data' />
<Button {handleClick} label='Create' />
