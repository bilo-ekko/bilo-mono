export type { 
  Locale, 
  TranslationDictionary, 
  TranslationKey, 
  TranslateFunction 
} from './types';

export { 
  createTranslator, 
  getAvailableLocales, 
  getTranslations,
  isValidLocale 
} from './translator';
