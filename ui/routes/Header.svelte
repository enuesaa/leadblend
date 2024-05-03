<script lang="ts">
	import { currentSpace } from '$lib/graphql/space'
</script>

<header class="container mx-auto pt-2 pb-4">
	<a href="/">
		<h1 class="font-zenkaku text-2xl font-bold">
			{$currentSpace.name}
		</h1>
	</a>
</header>
