# SvelteKit SDKs Integration Examples

This guide shows how to use `@bilo/translations` in the SvelteKit web-sdks-apps.

> **Note:** These are example patterns. Copy the relevant code into your actual app files.

---

## Example 1: Load Function with Translations

Create a layout load function that provides translations:

**File:** `src/routes/+layout.ts`

```typescript
import { createTranslator, type Locale } from '@bilo/translations';

export function load({ params }: { params: { locale?: string } }) {
  const locale: Locale = params.locale === 'de-DE' ? 'de-DE' : 'en-GB';
  const t = createTranslator(locale);
  
  return {
    t,
    locale,
  };
}
```

---

## Example 2: Page Load with Metadata

Load translations for a specific page:

**File:** `src/routes/checkout/+page.ts`

```typescript
import { createTranslator, type Locale } from '@bilo/translations';

export function load({ params }: { params: { locale?: string } }) {
  const locale: Locale = params.locale === 'de-DE' ? 'de-DE' : 'en-GB';
  const t = createTranslator(locale);
  
  return {
    t,
    locale,
    meta: {
      title: t('sdks.meta.checkoutTitle'),
    },
  };
}
```

---

## Example 3: Component Usage in Svelte

Use translations in a Svelte component:

```svelte
<script lang="ts">
  import type { PageData } from './$types';
  
  export let data: PageData;
  const { t } = data;
  
  const carbonFootprint = 21;
</script>

<svelte:head>
  <title>{t('sdks.meta.checkoutTitle')}</title>
</svelte:head>

<main>
  <h1>{t('sdks.home.heading')}</h1>
  <p>{t('sdks.home.description')}</p>
  
  <nav>
    <a href="/checkout">{t('sdks.home.checkoutLink')}</a>
    <a href="/post-purchase">{t('sdks.home.postPurchaseLink')}</a>
  </nav>
</main>
```

---

## Example 4: Checkout Widget Component

Full example with interpolated values:

**File:** `src/lib/components/CheckoutWidget.svelte`

```svelte
<script lang="ts">
  import { createTranslator } from '@bilo/translations';
  
  export let locale: 'en-GB' | 'de-DE' = 'en-GB';
  
  const t = createTranslator(locale);
  const carbonFootprint = 21;
  
  let climateActionEnabled = $state(true);
  let roundUpEnabled = $state(false);
</script>

<article class="widget">
  <div class="widget-content">
    <header class="header">
      <h2 class="title">{t('sdks.checkout.title')}</h2>
      <p class="subtitle">
        Support <span class="highlight">{t('sdks.checkout.environmentalProjects')}</span>
        and act on the ~{carbonFootprint} kgCO₂e footprint of this purchase - 
        about what <span class="highlight">{t('sdks.checkout.tree')}</span> 
        can capture in <span class="highlight">{t('sdks.checkout.year')}</span>!
      </p>
    </header>
    
    <div class="options">
      <div class="option">
        <span class="option-title">{t('sdks.checkout.climateAction')}</span>
      </div>
      <div class="option">
        <span class="option-title">{t('sdks.checkout.roundUp')}</span>
      </div>
    </div>
    
    <footer class="footer">
      <a href="#learn" class="learn-more">{t('common.learnMore')}</a>
      <div class="powered-by">
        <span>{t('common.poweredBy')}</span>
      </div>
    </footer>
  </div>
  
  {#if climateActionEnabled || roundUpEnabled}
    <div class="thank-you active">
      <span>{t('sdks.checkout.thankYou')}</span>
    </div>
  {/if}
</article>
```

---

## Example 5: Post-Purchase Widget Component

**File:** `src/lib/components/PostPurchaseWidget.svelte`

```svelte
<script lang="ts">
  import { createTranslator } from '@bilo/translations';
  
  export let locale: 'en-GB' | 'de-DE' = 'en-GB';
  
  const t = createTranslator(locale);
  const carbonFootprint = 21;
  
  function handleFindOutMore() {
    // Handle click
  }
</script>

<article class="widget">
  <div class="widget-content">
    <div class="card">
      <h2 class="title">{t('sdks.postPurchase.title')}</h2>
      <p class="description">
        {t('sdks.postPurchase.description', {
          carbonFootprint: carbonFootprint.toString(),
        })}
      </p>
    </div>
  </div>
  
  <footer class="footer">
    <div class="powered-by">
      <span>{t('common.poweredBy')}</span>
    </div>
    
    <button class="cta-button" onclick={handleFindOutMore}>
      <span>{t('sdks.postPurchase.findOutMore')}</span>
    </button>
  </footer>
  
  <div class="label">{t('sdks.postPurchase.embeddedSdk')}</div>
</article>
```

---

## Example 6: Reading Locale from Query String

Get locale from URL search params:

```typescript
import { page } from '$app/stores';
import { derived } from 'svelte/store';
import { createTranslator, isValidLocale, type Locale } from '@bilo/translations';

// Create a derived store for the current locale
export const currentLocale = derived(page, ($page) => {
  const localeParam = $page.url.searchParams.get('locale');
  return localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB';
});

// Create a derived store for the translator function
export const translator = derived(currentLocale, ($locale) => {
  return createTranslator($locale as Locale);
});
```

### Usage in Component

```svelte
<script lang="ts">
  import { translator } from '$lib/stores/locale';
</script>

<h1>{$translator('sdks.home.heading')}</h1>
```

---

## Example 7: Locale Switcher Helper

Create a utility for switching locales:

```typescript
import { goto } from '$app/navigation';
import type { Locale } from '@bilo/translations';

export function createLocaleSwitcher(
  currentLocale: Locale,
  currentUrl: URL
) {
  return {
    currentLocale,
    availableLocales: ['en-GB', 'de-DE'] as const,
    switchLocale: (newLocale: Locale) => {
      if (newLocale !== currentLocale) {
        const newUrl = new URL(currentUrl);
        newUrl.searchParams.set('locale', newLocale);
        goto(newUrl.toString());
      }
    },
  };
}
```

### Usage in Component

```svelte
<script lang="ts">
  import { page } from '$app/stores';
  import { createLocaleSwitcher } from '$lib/utils/locale-switcher';
  
  export let data;
  const { locale } = data;
  
  $: localeSwitcher = createLocaleSwitcher(locale, $page.url);
</script>

<select 
  value={localeSwitcher.currentLocale}
  onchange={(e) => localeSwitcher.switchLocale(e.target.value)}
>
  {#each localeSwitcher.availableLocales as availableLocale}
    <option value={availableLocale}>{availableLocale}</option>
  {/each}
</select>
```

---

## Available Translation Keys

### Meta
- `sdks.meta.homeTitle`
- `sdks.meta.checkoutTitle`
- `sdks.meta.postPurchaseTitle`

### Home Page
- `sdks.home.heading`
- `sdks.home.description`
- `sdks.home.checkoutLink`
- `sdks.home.postPurchaseLink`

### Checkout Widget
- `sdks.checkout.title`
- `sdks.checkout.subtitle`
- `sdks.checkout.environmentalProjects`
- `sdks.checkout.tree`
- `sdks.checkout.year`
- `sdks.checkout.climateAction`
- `sdks.checkout.roundUp`
- `sdks.checkout.thankYou`
- `sdks.checkout.imageAlt`

### Post-Purchase Widget
- `sdks.postPurchase.title`
- `sdks.postPurchase.description`
- `sdks.postPurchase.findOutMore`
- `sdks.postPurchase.embeddedSdk`
- `sdks.postPurchase.imageAlt`
- `sdks.postPurchase.moreInformation`

### Common
- `common.learnMore`
- `common.poweredBy`
- `common.with`
