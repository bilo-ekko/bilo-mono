<script lang="ts">
	import CheckoutWidget from '$lib/components/CheckoutWidget.svelte';
</script>

<svelte:head>
	<title>Checkout | Ekko SDK</title>
</svelte:head>

<main class="page">
	<div class="container">
		<CheckoutWidget />
	</div>
</main>

<style>
	.page {
		min-height: 100vh;
		background: linear-gradient(135deg, #f0fdf4 0%, #ecfdf5 50%, #f0fdfa 100%);
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 2rem;
	}

	.container {
		width: 100%;
		max-width: 720px;
	}
</style>
