<script>
	import { useParams } from '$lib/page'

	const { planet: planetName, island: islandId } = useParams()
</script>

planet: {planetName} <br />
island: {islandId}
