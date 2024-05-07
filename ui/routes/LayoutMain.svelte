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
		@apply mx-auto flex;
	}
	.left {
		box-shadow: 0 0 2px rgba(255, 255, 255, 0.3) inset;
		@apply flex-none w-48 px-5 pt-2 bg-blackgrayer;
	}
	.right {
		box-shadow: 0 0 2px rgba(255, 255, 255, 0.3) inset;
		@apply flex-auto pt-2 pl-7;
	}
</style>
