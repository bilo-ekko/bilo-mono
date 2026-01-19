/**
 * Supported locales in the application
 */
export type Locale = 'en-GB' | 'de-DE';

/**
 * Translation dictionary structure
 */
export interface TranslationDictionary {
  common: {
    welcome: string;
    hello: string;
    goodbye: string;
    yes: string;
    no: string;
    save: string;
    cancel: string;
    delete: string;
    edit: string;
    create: string;
    update: string;
    search: string;
    loading: string;
    error: string;
    success: string;
    learnMore: string;
    poweredBy: string;
    with: string;
  };
  navigation: {
    home: string;
    dashboard: string;
    settings: string;
    profile: string;
    logout: string;
  };
  errors: {
    notFound: string;
    unauthorized: string;
    serverError: string;
    networkError: string;
    validationError: string;
  };
  validation: {
    required: string;
    invalidEmail: string;
    minLength: string;
    maxLength: string;
  };
  dashboard: {
    meta: {
      title: string;
      description: string;
    };
    hero: {
      heading: string;
      description: string;
      templatesLinkText: string;
      learningLinkText: string;
    };
    actions: {
      deployNow: string;
      documentation: string;
    };
    imageAlt: {
      nextjsLogo: string;
      vercelLogomark: string;
    };
  };
  sdks: {
    meta: {
      homeTitle: string;
      checkoutTitle: string;
      postPurchaseTitle: string;
    };
    home: {
      heading: string;
      description: string;
      checkoutLink: string;
      postPurchaseLink: string;
    };
    checkout: {
      title: string;
      subtitle: string;
      environmentalProjects: string;
      tree: string;
      year: string;
      climateAction: string;
      roundUp: string;
      thankYou: string;
      imageAlt: string;
    };
    postPurchase: {
      title: string;
      description: string;
      findOutMore: string;
      embeddedSdk: string;
      imageAlt: string;
      moreInformation: string;
    };
  };
}

/**
 * Translation key path helper type
 * Generates dot-notation paths like "common.welcome" or "dashboard.hero.heading"
 */
type NestedKeyOf<T> = T extends object
  ? {
      [K in keyof T]: K extends string
        ? T[K] extends object
          ? T[K] extends any[]
            ? never
            : `${K}.${NestedKeyOf<T[K]>}` | (T[K] extends { [key: string]: string } ? `${K}.${Extract<keyof T[K], string>}` : never)
          : `${K}`
        : never;
    }[keyof T]
  : never;

export type TranslationKey = NestedKeyOf<TranslationDictionary>;

/**
 * Translation function type
 */
export type TranslateFunction = (
  key: TranslationKey,
  params?: Record<string, string | number>
) => string;
