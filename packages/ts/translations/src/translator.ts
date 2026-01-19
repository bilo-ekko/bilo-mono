import type { Locale, TranslationDictionary, TranslationKey, TranslateFunction } from './types';
import enGB from './locales/en-GB.json';
import deDE from './locales/de-DE.json';

/**
 * Available translations mapped by locale
 */
const translations: Record<Locale, TranslationDictionary> = {
  'en-GB': enGB as TranslationDictionary,
  'de-DE': deDE as TranslationDictionary,
};

/**
 * Get nested value from object using dot-notation path
 */
function getNestedValue(obj: unknown, path: string): string | undefined {
  const keys = path.split('.');
  let current: unknown = obj;

  for (const key of keys) {
    if (current && typeof current === 'object' && key in current) {
      current = (current as Record<string, unknown>)[key];
    } else {
      return undefined;
    }
  }

  return typeof current === 'string' ? current : undefined;
}

/**
 * Replace template variables in translation strings
 * e.g., "Hello {{name}}" with params {name: "World"} => "Hello World"
 */
function interpolate(template: string, params?: Record<string, string | number>): string {
  if (!params) {
    return template;
  }

  return template.replace(/\{\{(\w+)\}\}/g, (match, key) => {
    const value = params[key];
    return value !== undefined ? String(value) : match;
  });
}

/**
 * Create a translator function for a specific locale
 */
export function createTranslator(locale: Locale): TranslateFunction {
  const dictionary = translations[locale];

  if (!dictionary) {
    throw new Error(`Unsupported locale: ${locale}`);
  }

  return (key: TranslationKey, params?: Record<string, string | number>): string => {
    const value = getNestedValue(dictionary, key);

    if (value === undefined) {
      console.warn(`Translation missing for key "${key}" in locale "${locale}"`);
      return key;
    }

    return interpolate(value, params);
  };
}

/**
 * Get all available locales
 */
export function getAvailableLocales(): Locale[] {
  return Object.keys(translations) as Locale[];
}

/**
 * Get the full translation dictionary for a locale
 */
export function getTranslations(locale: Locale): TranslationDictionary {
  const dictionary = translations[locale];

  if (!dictionary) {
    throw new Error(`Unsupported locale: ${locale}`);
  }

  return dictionary;
}

/**
 * Check if a locale is supported
 */
export function isValidLocale(locale: string): locale is Locale {
  return locale === 'en-GB' || locale === 'de-DE';
}
