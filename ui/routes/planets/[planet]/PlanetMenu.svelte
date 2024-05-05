<script lang="ts">
	import PageTitle from '$lib/components/PageTitle.svelte'
	import Icon from '@iconify/svelte'

	export let planetName: string
</script>

<PageTitle title={planetName} />

<nav>
	<a href={`/planets/${planetName}`}>islands</a>
	<a href={`/planets/${planetName}/island/create`}>
		<Icon icon="fa-solid:plus" />
	</a>
</nav>

<style lang="postcss">
	nav {
		box-shadow: 0 1px 1px 1px rgba(0,0,0,0.3);
		@apply my-2 py-2 w-10/12 bg-blackgray text-gray;
	}
	a {
		@apply inline-block px-2 ml-1 rounded-lg;
	}
</style>
