<script lang="ts">
	import PostPurchaseWidget from '$lib/components/PostPurchaseWidget.svelte';
</script>

<svelte:head>
	<title>Post-Purchase | Ekko SDK</title>
</svelte:head>

<main class="page">
	<div class="container">
		<PostPurchaseWidget />
	</div>
</main>

<style>
	.page {
		min-height: 100vh;
		background: linear-gradient(180deg, #e5e7eb 0%, #d1d5db 100%);
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 4rem 2rem;
	}

	.container {
		width: 100%;
		max-width: 720px;
	}
</style>
