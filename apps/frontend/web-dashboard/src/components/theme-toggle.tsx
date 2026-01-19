"use client";

import * as React from "react";
import { Moon, Sun, Monitor } from "lucide-react";
import { useTheme } from "next-themes";

export function ThemeToggle() {
  const { theme, setTheme } = useTheme();
  const [mounted, setMounted] = React.useState(false);

  React.useEffect(() => {
    setMounted(true);
  }, []);

  if (!mounted) {
    return (
      <div className="flex items-center gap-1 rounded-lg bg-zinc-100 p-1 dark:bg-zinc-800">
        <button className="rounded-md p-2 hover:bg-zinc-200 dark:hover:bg-zinc-700">
          <Monitor className="h-4 w-4" />
        </button>
      </div>
    );
  }

  return (
    <div className="flex items-center gap-1 rounded-lg bg-zinc-100 p-1 dark:bg-zinc-800">
      <button
        onClick={() => setTheme("light")}
        className={`rounded-md p-2 transition-colors ${
          theme === "light"
            ? "bg-white shadow-sm dark:bg-zinc-900"
            : "hover:bg-zinc-200 dark:hover:bg-zinc-700"
        }`}
        title="Light theme"
      >
        <Sun className="h-4 w-4" />
      </button>
      <button
        onClick={() => setTheme("dark")}
        className={`rounded-md p-2 transition-colors ${
          theme === "dark"
            ? "bg-white shadow-sm dark:bg-zinc-900"
            : "hover:bg-zinc-200 dark:hover:bg-zinc-700"
        }`}
        title="Dark theme"
      >
        <Moon className="h-4 w-4" />
      </button>
      <button
        onClick={() => setTheme("system")}
        className={`rounded-md p-2 transition-colors ${
          theme === "system"
            ? "bg-white shadow-sm dark:bg-zinc-900"
            : "hover:bg-zinc-200 dark:hover:bg-zinc-700"
        }`}
        title="System theme"
      >
        <Monitor className="h-4 w-4" />
      </button>
    </div>
  );
}
