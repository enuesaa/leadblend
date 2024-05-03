<script>
	import { useParams } from '$lib/page'

	const { planet: planetName } = useParams()
</script>

<h3>{planetName}</h3>

<a href={`/planets/${planetName}/create`}>create</a>
