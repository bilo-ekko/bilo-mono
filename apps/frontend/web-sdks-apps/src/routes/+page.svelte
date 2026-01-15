<script lang="ts">
</script>

<main class="container">
	<h1>Ekko SDK Demo</h1>
	<p>Choose a page to view:</p>
	<nav class="nav">
		<a href="/checkout" class="nav-link">Checkout Widget</a>
		<a href="/post-purchase" class="nav-link">Post-Purchase Widget</a>
	</nav>
</main>

<style>
	.container {
		min-height: 100vh;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 2rem;
		gap: 1.5rem;
	}

	h1 {
		font-size: 2.5rem;
		font-weight: 700;
		color: var(--color-primary);
	}

	p {
		color: var(--color-text-light);
		font-size: 1.125rem;
	}

	.nav {
		display: flex;
		gap: 1rem;
		margin-top: 1rem;
	}

	.nav-link {
		padding: 1rem 2rem;
		background: var(--color-primary);
		color: white;
		border-radius: var(--radius-md);
		font-weight: 600;
		transition: all 0.2s ease;
		box-shadow: var(--shadow-md);
	}

	.nav-link:hover {
		background: var(--color-primary-light);
		transform: translateY(-2px);
		box-shadow: var(--shadow-lg);
	}
</style>
