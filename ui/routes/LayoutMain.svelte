<script lang="ts">
	import SideMenu from './SideMenu.svelte'
</script>

<main>
	<div class="left">
		<SideMenu />
	</div>
	<div class="right">
		<slot />
	</div>
</main>

<style lang="postcss">
	main {
		min-height: 100vh;
		@apply mx-auto flex gap-6;
	}
	.left {
		@apply flex-none w-48 px-5 pt-2 bg-blackgray text-gray;
	}
	.right {
		@apply flex-auto pt-2;
	}
</style>
