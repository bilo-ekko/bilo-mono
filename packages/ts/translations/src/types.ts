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
    theme: string;
    language: string;
    show: string;
    hide: string;
  };
  navigation: {
    home: string;
    dashboard: string;
    settings: string;
    profile: string;
    logout: string;
    organisation: string;
    console: string;
    documentation: string;
    marketingToolkit: string;
    sections: {
      management: string;
      developer: string;
      resources: string;
    };
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
    pages: {
      dashboard: {
        title: string;
        description: string;
        cardTitle: string;
        cardDescription: string;
      };
      organisation: {
        title: string;
        description: string;
        tabs: {
          info: string;
          hierarchy: string;
          users: string;
        };
        info: {
          basicInformation: string;
          organisationName: string;
          organisationNamePlaceholder: string;
          legalName: string;
          legalNamePlaceholder: string;
          billingAddress: string;
          streetAddress: string;
          streetAddressPlaceholder: string;
          city: string;
          cityPlaceholder: string;
          postalCode: string;
          postalCodePlaceholder: string;
          country: string;
          countryPlaceholder: string;
          contactInformation: string;
          email: string;
          emailPlaceholder: string;
          phoneNumber: string;
          phoneNumberPlaceholder: string;
          saveChanges: string;
        };
        hierarchy: {
          organisationStructure: string;
          description: string;
          chartPlaceholder: string;
          departments: string;
          departmentsDescription: string;
        };
        users: {
          usersList: string;
          addUser: string;
          description: string;
          rolesPermissions: string;
          rolesDescription: string;
        };
      };
      console: {
        title: string;
        description: string;
        apiKeys: {
          title: string;
          description: string;
        };
        webhooks: {
          title: string;
          description: string;
        };
        logs: {
          title: string;
          description: string;
        };
      };
      documentation: {
        title: string;
        description: string;
        gettingStarted: {
          title: string;
          description: string;
          readMore: string;
        };
        apiReference: {
          title: string;
          description: string;
          readMore: string;
        };
        sdksLibraries: {
          title: string;
          description: string;
          readMore: string;
        };
        tutorials: {
          title: string;
          description: string;
          readMore: string;
        };
      };
      marketingToolkit: {
        title: string;
        description: string;
        brandAssets: {
          title: string;
          description: string;
          downloadLogos: string;
          viewGuidelines: string;
        };
        socialMediaTemplates: {
          title: string;
          description: string;
          browseTemplates: string;
        };
        emailTemplates: {
          title: string;
          description: string;
          viewTemplates: string;
        };
      };
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
