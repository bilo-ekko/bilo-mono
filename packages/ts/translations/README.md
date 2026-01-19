# @bilo/translations

Internationalization (i18n) package providing translations and utilities for bilo-mono frontend applications.

## Features

- 🌍 Multi-language support (English, German)
- 🔒 Type-safe translation keys
- 📦 Framework-agnostic (works with Next.js, SvelteKit, etc.)
- 🎯 Simple interpolation for dynamic values
- 🚀 Zero runtime dependencies

## Supported Locales

- `en-GB` - English (United Kingdom)
- `de-DE` - German (Germany)

## Translation Scopes

The translations are organized into the following scopes:

- **common**: Shared UI elements (buttons, actions, etc.)
- **navigation**: Navigation menu items
- **errors**: Error messages
- **validation**: Form validation messages
- **dashboard**: Next.js web-dashboard app specific translations
- **sdks**: SvelteKit web-sdks-apps specific translations

## Installation

This package is part of the bilo-mono workspace. Reference it in your package.json:

```json
{
  "dependencies": {
    "@bilo/translations": "workspace:*"
  }
}
```

## Usage

### Basic Usage

```typescript
import { createTranslator } from '@bilo/translations';

// Create a translator for a specific locale
const t = createTranslator('en-GB');

// Use the translator
console.log(t('common.welcome')); // "Welcome"
console.log(t('navigation.home')); // "Home"
console.log(t('errors.notFound')); // "Not found"
```

### With Interpolation

```typescript
const t = createTranslator('en-GB');

// Use template variables
console.log(t('validation.minLength', { min: 8 }));
// "Minimum length is 8 characters"

console.log(t('validation.maxLength', { max: 100 }));
// "Maximum length is 100 characters"
```

### Locale Utilities

```typescript
import { getAvailableLocales, isValidLocale, getTranslations } from '@bilo/translations';

// Get all available locales
const locales = getAvailableLocales(); // ['en-GB', 'de-DE']

// Check if a locale is valid
if (isValidLocale('en-GB')) {
  console.log('Valid locale');
}

// Get full translation dictionary
const translations = getTranslations('de-DE');
console.log(translations.common.welcome); // "Willkommen"
```

## Framework Integration Examples

### Next.js (App Router)

```typescript
// app/[locale]/layout.tsx
import { createTranslator, type Locale } from '@bilo/translations';

export default function LocaleLayout({
  children,
  params,
}: {
  children: React.ReactNode;
  params: { locale: Locale };
}) {
  const t = createTranslator(params.locale);
  
  return (
    <html lang={params.locale}>
      <body>{children}</body>
    </html>
  );
}
```

### SvelteKit

```typescript
// src/routes/+layout.ts
import { createTranslator } from '@bilo/translations';

export const load = ({ params }) => {
  const locale = params.locale || 'en-GB';
  const t = createTranslator(locale);
  
  return { t, locale };
};
```

## Adding New Translations

1. Add the translation key to both `src/locales/en-GB.json` and `src/locales/de-DE.json`
2. Update the `TranslationDictionary` type in `src/types.ts`
3. Run `pnpm build` to regenerate type definitions

## Development

```bash
# Build the package
pnpm build

# Watch mode for development
pnpm dev
```

## Type Safety

All translation keys are type-checked at compile time. Invalid keys will produce TypeScript errors:

```typescript
const t = createTranslator('en-GB');

t('common.welcome'); // ✅ Valid
t('invalid.key');    // ❌ TypeScript error
```

## License

MIT
