<script lang="ts">
	import { currentSpace } from '$lib/graphql/space'
</script>

<header>
	{$currentSpace.name}
</header>
