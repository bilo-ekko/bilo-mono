<script lang="ts">
	import { page } from '$app/stores';
	import { createTranslator, isValidLocale } from '@bilo/translations';
	import CheckoutWidget from '$lib/components/CheckoutWidget.svelte';
	
	// Get locale from query string parameter, fallback to en-GB
	$: localeParam = $page.url.searchParams.get('locale');
	$: locale = localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB';
	$: t = createTranslator(locale);
</script>

<svelte:head>
	<title>{t('sdks.meta.checkoutTitle')}</title>
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
