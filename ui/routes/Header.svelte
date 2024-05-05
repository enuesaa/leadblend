<script lang="ts">
	import { getCurrentSpace } from '$lib/graphql/space'

	const space = getCurrentSpace()
</script>

<header class="container mx-auto px-2 pt-1">
	<div class="flex">
		<a href="/" class="w-32">
			<h1 class="font-zenkaku text-2xl font-bold">
				{$space.data?.name}
			</h1>
		</a>

		<div class="flex-auto"></div>
		<a href="/patterns" class="w-32 font-zenkaku text-lg font-bold">
			Patterns
		</a>
	</div>
</header>
