# Next.js Dashboard Integration Examples

This guide shows how to use `@bilo/translations` in the Next.js web-dashboard app.

> **Note:** These are example patterns. Copy the relevant code into your actual app files.

---

## Example 1: Metadata with Translations

Generate metadata using translations:

```typescript
import { createTranslator, type Locale } from '@bilo/translations';
import type { Metadata } from 'next';

export function generateMetadata(locale: Locale = 'en-GB'): Metadata {
  const t = createTranslator(locale);
  
  return {
    title: t('dashboard.meta.title'),
    description: t('dashboard.meta.description'),
  };
}
```

---

## Example 2: Page Component with Translations

Use translations in a page component:

```tsx
import { createTranslator, type Locale } from '@bilo/translations';

export function DashboardHomePage({ locale = 'en-GB' }: { locale?: Locale }) {
  const t = createTranslator(locale);
  
  return (
    <div className="flex min-h-screen items-center justify-center">
      <main className="flex min-h-screen w-full max-w-3xl flex-col items-center justify-between py-32 px-16">
        {/* Hero Section */}
        <div className="flex flex-col items-center gap-6">
          <h1 className="text-3xl font-semibold">
            {t('dashboard.hero.heading')}
          </h1>
          <p className="text-lg">
            {t('dashboard.hero.description', {
              templatesLink: t('dashboard.hero.templatesLinkText'),
              learningLink: t('dashboard.hero.learningLinkText'),
            })}
          </p>
        </div>

        {/* Actions */}
        <div className="flex flex-col gap-4">
          <button className="flex h-12 items-center justify-center rounded-full">
            {t('dashboard.actions.deployNow')}
          </button>
          <button className="flex h-12 items-center justify-center rounded-full">
            {t('dashboard.actions.documentation')}
          </button>
        </div>
      </main>
    </div>
  );
}
```

---

## Example 3: Server Component with Dynamic Locale

Use locale from route params:

```tsx
import { createTranslator, type Locale } from '@bilo/translations';

export async function DashboardLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { locale: Locale };
}) {
  const t = createTranslator(params.locale);
  
  return (
    <html lang={params.locale}>
      <head>
        <title>{t('dashboard.meta.title')}</title>
        <meta name="description" content={t('dashboard.meta.description')} />
      </head>
      <body>{children}</body>
    </html>
  );
}
```

---

## Example 4: React Context Provider Pattern

Create a translation context for use throughout your app:

```tsx
import { createContext, useContext } from 'react';
import { createTranslator, type Locale, type TranslateFunction } from '@bilo/translations';

const TranslationContext = createContext<TranslateFunction | null>(null);

export function TranslationProvider({
  locale,
  children,
}: {
  locale: Locale;
  children: React.ReactNode;
}) {
  const t = createTranslator(locale);
  
  return (
    <TranslationContext.Provider value={t}>
      {children}
    </TranslationContext.Provider>
  );
}

export function useTranslation(): TranslateFunction {
  const t = useContext(TranslationContext);
  if (!t) {
    throw new Error('useTranslation must be used within TranslationProvider');
  }
  return t;
}
```

### Usage in Components

```tsx
export function ExampleComponent() {
  const t = useTranslation();
  
  return (
    <div>
      <h1>{t('dashboard.hero.heading')}</h1>
      <button>{t('dashboard.actions.deployNow')}</button>
    </div>
  );
}
```

---

## Example 5: Reading Locale from Query String

Get locale from URL search params:

```tsx
'use client';

import { useSearchParams } from 'next/navigation';
import { createTranslator, isValidLocale } from '@bilo/translations';

export function ClientComponent() {
  const searchParams = useSearchParams();
  const localeParam = searchParams.get('locale');
  const locale = localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB';
  
  const t = createTranslator(locale);
  
  return (
    <div>
      <h1>{t('dashboard.hero.heading')}</h1>
    </div>
  );
}
```

---

## Available Translation Keys

### Metadata
- `dashboard.meta.title`
- `dashboard.meta.description`

### Hero Section
- `dashboard.hero.heading`
- `dashboard.hero.description`
- `dashboard.hero.templatesLinkText`
- `dashboard.hero.learningLinkText`

### Actions
- `dashboard.actions.deployNow`
- `dashboard.actions.documentation`

### Image Alt Text
- `dashboard.imageAlt.nextjsLogo`
- `dashboard.imageAlt.vercelLogomark`
