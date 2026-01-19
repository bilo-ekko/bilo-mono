"use client";

import * as React from "react";
import { useRouter, usePathname, useSearchParams } from "next/navigation";

const languages = [
  { code: "en-GB", label: "English" },
  { code: "de-DE", label: "German" },
//   { code: "fr-FR", label: "FR" },
//   { code: "es-ES", label: "ES" },
] as const;

export function LanguagePicker() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const [mounted, setMounted] = React.useState(false);

  const currentLocale = searchParams.get("locale") || "en-GB";

  React.useEffect(() => {
    setMounted(true);
  }, []);

  const handleLanguageChange = (langCode: string) => {
    const params = new URLSearchParams(searchParams.toString());
    params.set("locale", langCode);
    router.push(`${pathname}?${params.toString()}`);
  };

  if (!mounted) {
    return (
      <div className="grid grid-cols-2 gap-1 rounded-lg bg-zinc-100 p-1 dark:bg-zinc-800">
        {languages.map((lang) => (
          <button
            key={lang.code}
            className="rounded-md px-3 py-2 text-xs font-medium transition-colors hover:bg-zinc-200 dark:hover:bg-zinc-700"
          >
            {lang.label}
          </button>
        ))}
      </div>
    );
  }

  return (
    <div className="grid grid-cols-2 gap-1 rounded-lg bg-zinc-100 p-1 dark:bg-zinc-800">
      {languages.map((lang) => (
        <button
          key={lang.code}
          onClick={() => handleLanguageChange(lang.code)}
          className={`rounded-md px-3 py-2 text-xs font-medium transition-colors ${
            currentLocale === lang.code
              ? "bg-white shadow-sm dark:bg-zinc-900"
              : "hover:bg-zinc-200 dark:hover:bg-zinc-700"
          }`}
          title={lang.code}
        >
          {lang.label}
        </button>
      ))}
    </div>
  );
}
