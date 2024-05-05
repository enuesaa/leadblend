<script lang="ts">
	import Button from '$lib/components/Button.svelte'
	import TextInput from '$lib/components/TextInput.svelte'
	import { useCreateComet } from '$lib/graphql/comet'
	import { goto } from '$app/navigation'
	import PageTitle from '$lib/components/PageTitle.svelte'

	const createComet = useCreateComet()

	let fielda: string = ''
	let fieldb: string = ''
	let fieldc: string = ''

	async function handleClick() {
		const data: any = {}
		if (fielda !== '') {
			data.a = fielda
		}
		if (fieldb !== '') {
			data.b = fieldb
		}
		if (fieldc !== '') {
			data.c = fieldc
		}
		await $createComet.mutateAsync({ data: JSON.stringify(data) })
		goto('/')
	}
</script>

<PageTitle title="New Comet" />

<TextInput bind:value={fielda} label="$.a" />
<TextInput bind:value={fieldb} label="$.b" />
<TextInput bind:value={fieldc} label="$.c" />
<Button {handleClick} label="Create" />
