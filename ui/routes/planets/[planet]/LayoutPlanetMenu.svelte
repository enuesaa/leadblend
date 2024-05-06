<script lang="ts">
	export let planetName: string
</script>

<nav>
	<a href={`/planets/${planetName}`}>
		<h2 class="text-3xl font-zenkaku font-bold mx-3">{planetName}</h2>
	</a>
	<a href={`/planets/${planetName}`}>Islands</a>
	<a href={`/planets/${planetName}/islands/new`}>New Island</a>
	<a href={`/planets/${planetName}/configure`}>Configure</a>
</nav>

<style lang="postcss">
	nav {
		box-shadow: 0 1px 1px 1px rgba(0, 0, 0, 0.3) inset;
		@apply bg-blackgray text-gray w-10/12 py-2 pb-3 rounded-md mb-2;
	}
	a {
		@apply inline-block px-2 ml-1 rounded-lg;
	}
</style>
