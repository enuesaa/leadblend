<script lang="ts">
	import { fetchNotes } from '$lib/api'
	import { createQuery } from '@tanstack/svelte-query'

	const notes = createQuery({
		queryKey: ['notes'],
		queryFn: fetchNotes
	})
</script>

{#if $notes.isLoading}
	<p>読み込み中... </p>
{:else if $notes.isError}
	<p>エラーが起きました</p>
{:else if $notes.isSuccess}
	{#each $notes.data as note}
		<p>タイトル: {note.name}</p>
	{/each}
{/if}
