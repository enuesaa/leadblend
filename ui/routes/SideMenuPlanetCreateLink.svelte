<script lang="ts">
	import Icon from '@iconify/svelte'
</script>

<a href="/planets/new">
	<Icon icon="fa-solid:plus" />
</a>

<style lang="postcss">
	a {
		border: solid 1.5px;
		@apply m-1 px-1 font-zenkaku font-bold border-gray;
		@apply rounded-xl w-5 h-5 flex justify-center items-center;
	}
</style>
