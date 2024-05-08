<script lang="ts">
	import { getCurrentSpace } from '$lib/graphql/space'

	const space = getCurrentSpace()
</script>

<header class="mx-auto px-8 py-1 bg-blackgrayer">
	<div class="flex">
		<a href="/" class="w-52 text-center">
			<h1 class="font-zenkaku text-2xl font-bold">
				{$space.data?.name}
			</h1>
		</a>

		<div class="flex-auto"></div>
		<a href="/patterns" class="inline-block w-20 font-zenkaku text-lg font-bold text-center">Patterns</a>
	</div>
</header>
