<script lang="ts">
	import Icon from '@iconify/svelte'
</script>

<a href="/planet/create">
	<Icon icon="fa-solid:plus" />
</a>

<style lang="postcss">
	a {
		border: solid 2px #1a1a1a;
		@apply m-1 px-1 font-zenkaku font-bold;
		@apply rounded-xl w-6 h-6 flex justify-center items-center;
	}
</style>
