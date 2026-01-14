<script lang="ts">
	import '../app.css';
	
	let { children } = $props();
</script>

<svelte:head>
	<title>Ekko SDK</title>
	<meta name="description" content="Support climate action with every purchase" />
</svelte:head>

{@render children()}
