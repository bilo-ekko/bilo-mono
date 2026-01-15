<script lang="ts">
	import EkkoLogo from './EkkoLogo.svelte';
	
	const carbonFootprint = 21;
	
	function handleFindOutMore() {
		// Handle click
	}
	
	function handleInfoClick() {
		// Handle info click
	}
</script>

<article class="widget">
	<div class="background-image">
		<img
			src="https://images.unsplash.com/photo-1448375240586-882707db888b?w=1200&h=600&fit=crop&q=80"
			alt="Dense forest canopy from above"
			loading="lazy"
		/>
	</div>
	
	<div class="widget-content">
		<div class="card">
			<div class="card-content">
				<h2 class="title">Give a little. Change a lot.</h2>
				<p class="description">
					Support climate projects and act on the carbon
					footprint of this purchase (~{carbonFootprint} kgCO2e).
				</p>
			</div>
		</div>
		
		<button class="info-button" onclick={handleInfoClick} aria-label="More information">
			<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
				<path d="M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18Z" stroke="currentColor" stroke-width="1.5"/>
				<path d="M10 14V10M10 6H10.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
			</svg>
		</button>
	</div>
	
	<footer class="footer">
		<div class="powered-by">
			<span>Powered by</span>
			<EkkoLogo size="md" />
		</div>
		
		<button class="cta-button" onclick={handleFindOutMore}>
			<span>Find out more</span>
			<svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg">
				<path d="M4 10H16M16 10L11 5M16 10L11 15" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"/>
			</svg>
		</button>
	</footer>
	
	<div class="label">Embedded SDK</div>
</article>

<style>
	.widget {
		position: relative;
		border-radius: var(--radius-xl);
		overflow: hidden;
		max-width: 720px;
		width: 100%;
		box-shadow: var(--shadow-xl);
		background: #1a2e1a;
	}

	.background-image {
		position: absolute;
		inset: 0;
		z-index: 0;
	}

	.background-image img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		opacity: 0.85;
	}

	.widget-content {
		position: relative;
		z-index: 1;
		padding: 24px;
		display: flex;
		justify-content: space-between;
		align-items: flex-start;
		gap: 16px;
	}

	.card {
		background: rgba(30, 60, 35, 0.85);
		backdrop-filter: blur(8px);
		border-radius: var(--radius-lg);
		border: 1px solid rgba(255, 255, 255, 0.1);
		padding: 20px 24px;
		max-width: 480px;
	}

	.card-content {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}

	.title {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--color-accent-bright);
		line-height: 1.2;
	}

	.description {
		font-size: 1rem;
		color: white;
		line-height: 1.5;
		opacity: 0.95;
	}

	.info-button {
		width: 44px;
		height: 44px;
		background: rgba(40, 70, 45, 0.9);
		backdrop-filter: blur(8px);
		border-radius: var(--radius-md);
		display: flex;
		align-items: center;
		justify-content: center;
		color: white;
		transition: all 0.2s ease;
		flex-shrink: 0;
	}

	.info-button:hover {
		background: rgba(50, 85, 55, 0.95);
		transform: scale(1.05);
	}

	.footer {
		position: relative;
		z-index: 1;
		padding: 16px 24px 24px;
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 16px;
	}

	.powered-by {
		display: flex;
		align-items: center;
		gap: 10px;
		font-size: 0.9rem;
		color: rgba(255, 255, 255, 0.8);
	}

	.powered-by :global(.ekko-logo .logo-text) {
		color: white;
	}

	.cta-button {
		display: flex;
		align-items: center;
		gap: 10px;
		padding: 14px 28px;
		background: linear-gradient(135deg, #3b82f6, #2563eb);
		color: white;
		font-size: 1rem;
		font-weight: 600;
		border-radius: 50px;
		transition: all 0.2s ease;
		box-shadow: 0 4px 14px rgba(37, 99, 235, 0.4);
	}

	.cta-button:hover {
		background: linear-gradient(135deg, #4f91ff, #3b82f6);
		transform: translateY(-2px);
		box-shadow: 0 6px 20px rgba(37, 99, 235, 0.5);
	}

	.cta-button svg {
		transition: transform 0.2s ease;
	}

	.cta-button:hover svg {
		transform: translateX(4px);
	}

	.label {
		position: absolute;
		top: 0;
		left: 24px;
		transform: translateY(-100%);
		font-size: 0.75rem;
		color: var(--color-text-muted);
		padding: 4px 0;
	}

	@media (max-width: 640px) {
		.widget-content {
			flex-direction: column;
			align-items: stretch;
		}

		.card {
			max-width: none;
		}

		.info-button {
			position: absolute;
			top: 24px;
			right: 24px;
		}

		.footer {
			flex-direction: column;
			align-items: stretch;
			gap: 16px;
		}

		.cta-button {
			justify-content: center;
		}

		.powered-by {
			justify-content: center;
		}
	}
</style>
