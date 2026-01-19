/**
 * Example usage of the @bilo/translations package
 * 
 * This file demonstrates various ways to use the translations package
 * in your frontend applications.
 */

import { createTranslator, getAvailableLocales, isValidLocale } from '../src';
import type { Locale } from '../src';

// Example 1: Basic translation
function example1() {
  console.log('=== Example 1: Basic Translation ===');
  
  const t = createTranslator('en-GB');
  console.log(t('common.welcome')); // "Welcome"
  console.log(t('navigation.home')); // "Home"
  console.log(t('errors.notFound')); // "Not found"
}

// Example 2: German translations
function example2() {
  console.log('\n=== Example 2: German Translation ===');
  
  const t = createTranslator('de-DE');
  console.log(t('common.welcome')); // "Willkommen"
  console.log(t('navigation.home')); // "Startseite"
  console.log(t('errors.notFound')); // "Nicht gefunden"
}

// Example 3: Template interpolation
function example3() {
  console.log('\n=== Example 3: Template Interpolation ===');
  
  const t = createTranslator('en-GB');
  console.log(t('validation.minLength', { min: 8 }));
  // "Minimum length is 8 characters"
  
  console.log(t('validation.maxLength', { max: 100 }));
  // "Maximum length is 100 characters"
}

// Example 4: Dynamic locale switching
function example4() {
  console.log('\n=== Example 4: Dynamic Locale Switching ===');
  
  const userLocale: Locale = 'de-DE'; // Could come from user preferences
  const t = createTranslator(userLocale);
  
  console.log(`Locale: ${userLocale}`);
  console.log(t('common.save')); // "Speichern"
}

// Example 5: Locale validation
function example5() {
  console.log('\n=== Example 5: Locale Validation ===');
  
  const locales = getAvailableLocales();
  console.log('Available locales:', locales); // ['en-GB', 'de-DE']
  
  console.log('Is "en-GB" valid?', isValidLocale('en-GB')); // true
  console.log('Is "fr-FR" valid?', isValidLocale('fr-FR')); // false
}

// Example 6: React component (pseudo-code)
function example6React() {
  console.log('\n=== Example 6: React Component Pattern ===');
  console.log(`
// components/LocaleProvider.tsx
import { createTranslator, type Locale, type TranslateFunction } from '@bilo/translations';
import { createContext, useContext } from 'react';

const TranslationContext = createContext<TranslateFunction | null>(null);

export function LocaleProvider({ 
  locale, 
  children 
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

export function useTranslation() {
  const t = useContext(TranslationContext);
  if (!t) throw new Error('useTranslation must be used within LocaleProvider');
  return t;
}

// Usage in a component:
function WelcomeMessage() {
  const t = useTranslation();
  return <h1>{t('common.welcome')}</h1>;
}
  `);
}

// Example 7: SvelteKit pattern (pseudo-code)
function example7Svelte() {
  console.log('\n=== Example 7: SvelteKit Pattern ===');
  console.log(`
// src/routes/+layout.ts
import { createTranslator } from '@bilo/translations';

export const load = ({ params }) => {
  const locale = params.locale || 'en-GB';
  const t = createTranslator(locale);
  
  return { t, locale };
};

// Usage in a Svelte component:
<script lang="ts">
  export let data;
  const { t } = data;
</script>

<h1>{t('common.welcome')}</h1>
<nav>
  <a href="/">{t('navigation.home')}</a>
</nav>
  `);
}

// Run all examples
if (require.main === module) {
  example1();
  example2();
  example3();
  example4();
  example5();
  example6React();
  example7Svelte();
}
