<script lang="ts">
	import { page } from '$app/stores';
	import { createTranslator, isValidLocale } from '@bilo/translations';
	import Toggle from './Toggle.svelte';
	import EkkoLogo from './EkkoLogo.svelte';
	import GoldStandardLogo from './GoldStandardLogo.svelte';
	import PartnerLogos from './PartnerLogos.svelte';

	// Get locale from query string parameter, fallback to en-GB
	let localeParam = $derived($page.url.searchParams.get('locale'));
	let locale = $derived(localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB');
	let t = $derived(createTranslator(locale));

	let climateActionEnabled = $state(true);
	let roundUpEnabled = $state(false);
	
	const carbonFootprint = 21;
	
	let isActive = $derived(climateActionEnabled || roundUpEnabled);
	const climateActionCost = 0.65;
	const roundUpCost = 0.85;
</script>

<article class="widget">
	<div class="widget-content">
		<div class="image-section">
		<img
			src="https://images.unsplash.com/photo-1529963183134-61a90db47eaf?w=400&h=500&fit=crop&q=80"
			alt={t('sdks.checkout.imageAlt')}
			loading="lazy"
		/>
		</div>
		
		<div class="content-section">
			<header class="header">
				<h2 class="title">{t('sdks.checkout.title')}</h2>
				<p class="subtitle">
					Support <span class="highlight">{t('sdks.checkout.environmentalProjects')}</span> and act on the ~{carbonFootprint}
					kgCO<sub>2</sub>e footprint of this purchase - about what <span class="highlight">{t('sdks.checkout.tree')}</span>
					can capture in <span class="highlight">{t('sdks.checkout.year')}</span>!
				</p>
			</header>

			<div class="options">
				<div class="option">
					<div class="option-content">
						<div class="option-header">
							<span class="option-title">{t('sdks.checkout.climateAction')}</span>
							<span class="option-price">£{climateActionCost.toFixed(2)}</span>
						</div>
						<div class="option-partner">
							<span class="with-text">{t('common.with')}</span>
							<GoldStandardLogo />
						</div>
					</div>
					<Toggle bind:checked={climateActionEnabled} />
				</div>

				<div class="option">
					<div class="option-content">
						<div class="option-header">
							<span class="option-title">{t('sdks.checkout.roundUp')}</span>
							<span class="option-price">£{roundUpCost.toFixed(2)}</span>
						</div>
						<div class="option-partner">
							<span class="with-text">{t('common.with')}</span>
							<PartnerLogos />
						</div>
					</div>
					<Toggle bind:checked={roundUpEnabled} />
				</div>
			</div>

			<footer class="footer">
				<a href="#learn" class="learn-more">{t('common.learnMore')}</a>
				<div class="powered-by">
					<span>{t('common.poweredBy')}</span>
					<EkkoLogo size="sm" />
				</div>
			</footer>
		</div>
	</div>
	
	<div class="thank-you" class:active={isActive}>
		<span>{t('sdks.checkout.thankYou')}</span>
		<span class="leaf">🌿</span>
	</div>
</article>

<style>
	.widget {
		background: var(--color-card);
		border-radius: var(--radius-xl);
		box-shadow: var(--shadow-xl);
		overflow: hidden;
		max-width: 720px;
		width: 100%;
	}

	.widget-content {
		display: grid;
		grid-template-columns: 200px 1fr;
	}

	.image-section {
		position: relative;
		overflow: hidden;
	}

	.image-section img {
		width: 100%;
		height: 100%;
		object-fit: cover;
		min-height: 340px;
	}

	.content-section {
		padding: 24px 28px;
		display: flex;
		flex-direction: column;
		gap: 20px;
	}

	.header {
		display: flex;
		flex-direction: column;
		gap: 8px;
	}

	.title {
		font-size: 1.5rem;
		font-weight: 700;
		color: var(--color-primary);
		line-height: 1.2;
	}

	.subtitle {
		font-size: 0.9rem;
		color: var(--color-text);
		line-height: 1.5;
	}

	.subtitle sub {
		font-size: 0.7em;
	}

	.highlight {
		color: var(--color-primary);
		font-weight: 600;
	}

	.options {
		display: flex;
		flex-direction: column;
		gap: 16px;
	}

	.option {
		display: flex;
		align-items: center;
		justify-content: space-between;
		gap: 16px;
		padding-bottom: 16px;
		border-bottom: 1px solid var(--color-border);
	}

	.option:last-child {
		border-bottom: none;
		padding-bottom: 0;
	}

	.option-content {
		display: flex;
		flex-direction: column;
		gap: 6px;
	}

	.option-header {
		display: flex;
		align-items: center;
		gap: 12px;
	}

	.option-title {
		font-size: 0.95rem;
		font-weight: 600;
		color: var(--color-text);
	}

	.option-price {
		font-size: 0.9rem;
		font-weight: 600;
		color: var(--color-text-light);
	}

	.option-partner {
		display: flex;
		align-items: center;
		gap: 8px;
	}

	.with-text {
		font-size: 0.8rem;
		color: var(--color-text-muted);
	}

	.footer {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding-top: 4px;
	}

	.learn-more {
		font-size: 0.9rem;
		color: var(--color-primary);
		font-weight: 600;
		text-decoration: underline;
		text-underline-offset: 2px;
	}

	.learn-more:hover {
		color: var(--color-primary-light);
	}

	.powered-by {
		display: flex;
		align-items: center;
		gap: 8px;
		font-size: 0.8rem;
		color: var(--color-text-muted);
	}

	.thank-you {
		background: #e8ebe8;
		color: #1f2937;
		padding: 12px 24px;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 8px;
		font-size: 0.9rem;
		font-weight: 500;
		transition: background-color 0.3s ease, color 0.3s ease;
	}

	.thank-you.active {
		background: var(--color-primary);
		color: white;
	}

	.leaf {
		font-size: 1.1rem;
	}

	@media (max-width: 640px) {
		.widget-content {
			grid-template-columns: 1fr;
		}

		.image-section {
			height: 180px;
		}

		.image-section img {
			min-height: auto;
		}

		.content-section {
			padding: 20px;
		}

		.title {
			font-size: 1.25rem;
		}

		.option-header {
			flex-direction: column;
			align-items: flex-start;
			gap: 4px;
		}
	}
</style>
