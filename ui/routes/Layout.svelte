<script lang="ts">
	import Comets from './Comets.svelte'
	import SideMenu from './SideMenu.svelte'
	import { hideComet } from '$lib/store/comet'

	export let comet: boolean = true
	export let sidebar: boolean = true
	$: if (!comet) {
		hideComet()
	}
</script>

<Comets />
<main>
	{#if sidebar}
		<div class="left">
			<SideMenu />
		</div>
		<div class="right">
			<slot />
		</div>
	{:else}
		<div class="rightSidebarHidden">
			<slot />
		</div>
	{/if}
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
	.rightSidebarHidden {
		@apply flex-auto pt-3 pl-7 max-w-4xl mx-auto;
	}
</style>
