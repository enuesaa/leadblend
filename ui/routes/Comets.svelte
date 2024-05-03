<script>
	import CometCard from './CometCard.svelte'
	import CometCreateCard from './CometCreateCard.svelte'
</script>

<section class="p-2 m-2 flex gap-3">
	<CometCard />
	<CometCard />
	<CometCreateCard />
</section>

<style>
	section {
		border-top: solid 1px #1a1a1a;
		border-bottom: solid 1px #1a1a1a;
	}
</style>
