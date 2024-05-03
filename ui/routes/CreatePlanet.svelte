<script lang="ts">
	let name: string
	let comment: string

	async function hanldeClick() {
		const data = {name, comment}
		console.log(data)
		// const res = await fetch('https://example.com', {
		// 	method: 'POST',
		// 	body: JSON.stringify(formdata),
		// })
	}
</script>

<input type="text" bind:value={name} />
<input type="text" bind:value={comment} />
<button on:click|preventDefault={hanldeClick}>submit</button>
