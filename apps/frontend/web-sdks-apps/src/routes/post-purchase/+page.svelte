<script lang="ts">
	import { page } from '$app/stores';
	import { createTranslator, isValidLocale } from '@bilo/translations';
	import PostPurchaseWidget from '$lib/components/PostPurchaseWidget.svelte';
	
	// Get locale from query string parameter, fallback to en-GB
	$: localeParam = $page.url.searchParams.get('locale');
	$: locale = localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB';
	$: t = createTranslator(locale);
</script>

<svelte:head>
	<title>{t('sdks.meta.postPurchaseTitle')}</title>
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
