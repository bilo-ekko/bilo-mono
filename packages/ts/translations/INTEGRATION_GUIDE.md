# Integration Guide

This guide walks you through integrating `@bilo/translations` into your frontend apps.

## Quick Start

### 1. Install the Package

Add the dependency to your app's `package.json`:

```json
{
  "dependencies": {
    "@bilo/translations": "workspace:*"
  }
}
```

Then run:

```bash
pnpm install
```

### 2. Basic Usage

```typescript
import { createTranslator } from '@bilo/translations';

const t = createTranslator('en-GB');
console.log(t('common.welcome')); // "Welcome"
```

---

## Integration by App

### Web Dashboard (Next.js)

**Location**: `apps/frontend/web-dashboard`

#### Step 1: Add dependency

```bash
cd apps/frontend/web-dashboard
pnpm add @bilo/translations@workspace:*
```

#### Step 2: Update layout for metadata

```typescript
// src/app/layout.tsx
import { createTranslator } from '@bilo/translations';

const t = createTranslator('en-GB');

export const metadata = {
  title: t('dashboard.meta.title'),
  description: t('dashboard.meta.description'),
};
```

#### Step 3: Update page component

```typescript
// src/app/page.tsx
import { createTranslator } from '@bilo/translations';

export default function Home() {
  const t = createTranslator('en-GB');
  
  return (
    <div>
      <h1>{t('dashboard.hero.heading')}</h1>
      <button>{t('dashboard.actions.deployNow')}</button>
    </div>
  );
}
```

#### Available Translation Keys

All dashboard translations are under the `dashboard` scope:

- `dashboard.meta.*` - Metadata (title, description)
- `dashboard.hero.*` - Hero section content
- `dashboard.actions.*` - Button labels
- `dashboard.imageAlt.*` - Image alt text

See `TRANSLATION_MAP.md` for the complete mapping.

---

### Web SDKs Apps (SvelteKit)

**Location**: `apps/frontend/web-sdks-apps`

#### Step 1: Add dependency

```bash
cd apps/frontend/web-sdks-apps
pnpm add @bilo/translations@workspace:*
```

#### Step 2: Create a layout load function

```typescript
// src/routes/+layout.ts
import { createTranslator } from '@bilo/translations';

export function load() {
  const locale = 'en-GB'; // Could come from user preferences
  const t = createTranslator(locale);
  
  return { t, locale };
}
```

#### Step 3: Use in components

```svelte
<!-- src/routes/+page.svelte -->
<script lang="ts">
  export let data;
  const { t } = data;
</script>

<main>
  <h1>{t('sdks.home.heading')}</h1>
  <p>{t('sdks.home.description')}</p>
</main>
```

#### Step 4: Update widgets

```svelte
<!-- src/lib/components/CheckoutWidget.svelte -->
<script lang="ts">
  import { createTranslator } from '@bilo/translations';
  
  const t = createTranslator('en-GB');
  const carbonFootprint = 21;
</script>

<article class="widget">
  <h2>{t('sdks.checkout.title')}</h2>
  <p>
    {t('sdks.checkout.subtitle', {
      environmentalProjects: t('sdks.checkout.environmentalProjects'),
      carbonFootprint: carbonFootprint.toString(),
      subscript: '2',
      tree: t('sdks.checkout.tree'),
      year: t('sdks.checkout.year'),
    })}
  </p>
  <button>{t('sdks.checkout.climateAction')}</button>
  <a href="#learn">{t('common.learnMore')}</a>
</article>
```

#### Available Translation Keys

All SDK translations are under the `sdks` scope:

- `sdks.meta.*` - Page titles
- `sdks.home.*` - Home page content
- `sdks.checkout.*` - Checkout widget content
- `sdks.postPurchase.*` - Post-purchase widget content

See `TRANSLATION_MAP.md` for the complete mapping.

---

## Advanced Features

### Locale Switching

To support multiple languages, you can implement a locale switcher:

```typescript
// Next.js example
export default function LocaleSwitcher() {
  const [locale, setLocale] = useState<Locale>('en-GB');
  const t = createTranslator(locale);
  
  return (
    <select value={locale} onChange={(e) => setLocale(e.target.value as Locale)}>
      <option value="en-GB">English</option>
      <option value="de-DE">Deutsch</option>
    </select>
  );
}
```

### Server-Side Rendering

For Next.js App Router, you can use translations in Server Components:

```typescript
// app/[locale]/page.tsx
import { createTranslator, type Locale } from '@bilo/translations';

export default function Page({ params }: { params: { locale: Locale } }) {
  const t = createTranslator(params.locale);
  
  return <h1>{t('dashboard.hero.heading')}</h1>;
}
```

### React Context Pattern

For client components that need translations throughout the tree:

```typescript
// contexts/TranslationContext.tsx
import { createContext, useContext } from 'react';
import { createTranslator, type Locale, type TranslateFunction } from '@bilo/translations';

const TranslationContext = createContext<TranslateFunction | null>(null);

export function TranslationProvider({ 
  locale, 
  children 
}: { 
  locale: Locale; 
  children: React.ReactNode;
}) {
  const t = createTranslator(locale);
  return <TranslationContext.Provider value={t}>{children}</TranslationContext.Provider>;
}

export function useTranslation() {
  const t = useContext(TranslationContext);
  if (!t) throw new Error('useTranslation must be used within TranslationProvider');
  return t;
}
```

---

## Testing Translations

You can test translations in your components:

```typescript
import { createTranslator } from '@bilo/translations';

describe('Component with translations', () => {
  it('renders English text', () => {
    const t = createTranslator('en-GB');
    expect(t('common.welcome')).toBe('Welcome');
  });
  
  it('renders German text', () => {
    const t = createTranslator('de-DE');
    expect(t('common.welcome')).toBe('Willkommen');
  });
});
```

---

## Adding New Translations

1. Add the key to both `src/locales/en-GB.json` and `src/locales/de-DE.json`
2. Update the `TranslationDictionary` interface in `src/types.ts`
3. Rebuild the package: `pnpm build`
4. The new keys will be type-checked in your apps

---

## Troubleshooting

### "Translation missing" warning

If you see a warning like `Translation missing for key "xxx" in locale "en-GB"`:

1. Check that the key exists in `src/locales/en-GB.json`
2. Verify you're using the correct key path (e.g., `dashboard.hero.heading`)
3. Rebuild the translations package

### TypeScript errors

If you get type errors about invalid keys:

1. Make sure you've rebuilt the translations package after adding new keys
2. Restart your TypeScript server in your IDE
3. Check that the key exists in the `TranslationDictionary` interface

### Package not found

If your app can't find `@bilo/translations`:

1. Make sure you've added it to your app's `package.json`
2. Run `pnpm install` in your app directory
3. Check that the package is built (`dist/` folder exists)

---

## Next Steps

1. ✅ Package created with all current app content
2. 🔄 Integrate into `web-dashboard` app
3. 🔄 Integrate into `web-sdks-apps` app
4. 🔄 Add locale switching UI
5. 🔄 Test both English and German translations
6. 🔄 Consider adding more locales (French, Spanish, etc.)

For detailed examples, see:
- `examples/nextjs-dashboard.tsx` - Next.js integration patterns
- `examples/sveltekit-sdks.ts` - SvelteKit integration patterns
- `TRANSLATION_MAP.md` - Complete mapping of original text to keys
