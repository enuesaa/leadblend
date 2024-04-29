<script lang="ts">
	import { queryStore, gql, getContextClient } from '@urql/svelte'
  
	const todos = queryStore({
	  client: getContextClient(),
	  query: gql`
		query {
		  todos {
			id
			title
		  }
		}
	  `,
	})
</script>
  
{#if $todos.fetching}
	<p>Loading...</p>
{:else if $todos.error}
	<p>Oh no... {$todos.error.message}</p>
{:else}
	<ul>
	{#each $todos.data.todos as todo}
		<li>{todo.title}</li>
	{/each}
	</ul>
{/if}

  