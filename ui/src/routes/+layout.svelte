<script lang="ts">
	import { Client, setContextClient, cacheExchange, fetchExchange } from '@urql/svelte'
	import './app.css'

	const client = new Client({
		url: 'http://localhost:3000/graphql',
		exchanges: [cacheExchange, fetchExchange],
	});

	setContextClient(client);
</script>

<slot />
