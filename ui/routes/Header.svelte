<script lang="ts">
	import { getCurrentSpace } from '$lib/graphql/space'
	const space = getCurrentSpace()
</script>

<header class="container mx-auto pt-2 pb-4">
	<a href="/">
		<h1 class="font-zenkaku text-2xl font-bold">
			{$space.data?.name}
		</h1>
	</a>
</header>
