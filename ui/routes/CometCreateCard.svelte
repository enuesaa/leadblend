<script lang="ts">
	import Icon from '@iconify/svelte'
</script>

<a href="/comets/new">
	<Icon icon="fa-solid:plus" />
</a>

<style lang="postcss">
	a {
		border: dashed 1px rgba(255, 255, 255, 0.3);
		@apply p-2 rounded-lg w-40 h-40 bg-grayblack;
		@apply flex items-center justify-center;
	}
</style>
