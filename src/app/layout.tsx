import type { Metadata } from 'next';
import './globals.css';

export const metadata: Metadata = {
  title: 'RIT Racing Music Dashboard',
  description: 'Music dashboard for RIT racing.',
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
