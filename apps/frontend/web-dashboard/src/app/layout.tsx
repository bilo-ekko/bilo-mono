import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import { createTranslator } from "@bilo/translations";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

// TODO: Get locale from user preferences/URL when locale switching is implemented
const t = createTranslator("en-GB");

export const metadata: Metadata = {
  title: t("dashboard.meta.title"),
  description: t("dashboard.meta.description"),
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
