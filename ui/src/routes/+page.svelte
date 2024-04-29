<script lang="ts">
	import { queryStore, gql, getContextClient } from '@urql/svelte'
  
	const planets = queryStore({
		client: getContextClient(),
		query: gql`
			query {
				planets {
					id
				}
			}
		`,
	})
</script>
  
{#if $planets.fetching}
	<p>Loading...</p>
{:else if $planets.error}
	<p>Oh no... {$planets.error.message}</p>
{:else}
	<ul>
	{#each $planets.data.planets as planet}
		<li>{planet.id}</li>
	{/each}
	</ul>
{/if}

  