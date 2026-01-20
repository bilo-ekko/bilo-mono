"use client";

import * as React from "react";
import Link from "next/link";
import { usePathname, useSearchParams } from "next/navigation";
import {
  LayoutDashboard,
  Users,
  Terminal,
  BookOpen,
  Briefcase,
  ChevronRight,
  Settings,
} from "lucide-react";
import { ThemeToggle } from "./theme-toggle";
import { LanguagePicker } from "./language-picker";
import { createTranslator, type TranslationKey } from "@bilo/translations";

interface NavItem {
  titleKey: TranslationKey;
  href: string;
  icon: React.ReactNode;
}

interface NavSection {
  titleKey: TranslationKey;
  items: NavItem[];
}

const navSections: NavSection[] = [
  {
    titleKey: "navigation.sections.management",
    items: [
      {
        titleKey: "navigation.dashboard",
        href: "/dashboard",
        icon: <LayoutDashboard className="h-5 w-5" />,
      },
      {
        titleKey: "navigation.organisation",
        href: "/organisation",
        icon: <Users className="h-5 w-5" />,
      },
    ],
  },
  {
    titleKey: "navigation.sections.developer",
    items: [
      {
        titleKey: "navigation.console",
        href: "/console",
        icon: <Terminal className="h-5 w-5" />,
      },
      {
        titleKey: "navigation.documentation",
        href: "/documentation",
        icon: <BookOpen className="h-5 w-5" />,
      },
    ],
  },
  {
    titleKey: "navigation.sections.resources",
    items: [
      {
        titleKey: "navigation.marketingToolkit",
        href: "/marketing-toolkit",
        icon: <Briefcase className="h-5 w-5" />,
      },
    ],
  },
];

export function Sidebar() {
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const [mounted, setMounted] = React.useState(false);
  const [isSettingsExpanded, setIsSettingsExpanded] = React.useState(false);

  const locale = searchParams.get("locale") || "en-GB";
  const t = createTranslator(locale as "en-GB" | "de-DE");

  React.useEffect(() => {
    setMounted(true);
  }, []);

  return (
    <aside className="flex h-screen w-64 flex-col border-r border-zinc-200 bg-white dark:border-zinc-800 dark:bg-zinc-900">
      {/* Logo / Brand with Settings Toggle */}
      <div className="flex h-16 items-center justify-between border-b border-zinc-200 px-6 dark:border-zinc-800">
        <h1 className="text-xl font-bold text-zinc-900 dark:text-white">
          ekko
        </h1>
        <button
          onClick={() => setIsSettingsExpanded(!isSettingsExpanded)}
          className="rounded-lg p-2 transition-colors hover:bg-zinc-100 dark:hover:bg-zinc-800"
          title={
            isSettingsExpanded
              ? t("common.hide") || "Hide settings"
              : t("common.show") || "Show settings"
          }
        >
          <Settings className="h-5 w-5 text-zinc-600 dark:text-zinc-400" />
        </button>
      </div>

      {/* Settings - Collapsible */}
      <div
        className={`overflow-hidden border-b border-zinc-200 transition-all duration-300 dark:border-zinc-800 ${
          isSettingsExpanded ? "max-h-96 p-4" : "max-h-0 p-0"
        }`}
      >
        <div className="space-y-3">
          <div className="flex flex-col gap-2">
            <label className="text-xs font-semibold uppercase tracking-wider text-zinc-500 dark:text-zinc-400">
              {t("common.theme") || "Theme"}
            </label>
            <ThemeToggle />
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-xs font-semibold uppercase tracking-wider text-zinc-500 dark:text-zinc-400">
              {t("common.language") || "Language"}
            </label>
            <LanguagePicker />
          </div>
        </div>
      </div>

      {/* Navigation */}
      <nav className="flex-1 overflow-y-auto px-4 py-4">
        {navSections.map((section, sectionIndex) => (
          <div key={section.titleKey} className={sectionIndex > 0 ? "mt-6" : ""}>
            <h2 className="mb-2 px-2 text-xs font-semibold uppercase tracking-wider text-zinc-500 dark:text-zinc-400">
              {t(section.titleKey)}
            </h2>
            <div className="space-y-1">
              {section.items.map((item) => {
                const isActive = mounted && pathname === item.href;
                return (
                  <Link
                    key={item.href}
                    href={item.href}
                    className={`flex items-center gap-3 rounded-lg px-3 py-2 text-sm font-medium transition-colors ${
                      isActive
                        ? "bg-zinc-100 text-zinc-900 dark:bg-zinc-800 dark:text-white"
                        : "text-zinc-600 hover:bg-zinc-50 hover:text-zinc-900 dark:text-zinc-400 dark:hover:bg-zinc-800/50 dark:hover:text-white"
                    }`}
                  >
                    {item.icon}
                    <span className="flex-1">{t(item.titleKey)}</span>
                    {isActive && <ChevronRight className="h-4 w-4" />}
                  </Link>
                );
              })}
            </div>
          </div>
        ))}
      </nav>
    </aside>
  );
}
