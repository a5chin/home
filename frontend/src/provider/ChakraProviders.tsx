/* eslint-disable no-restricted-imports */
"use client";

import { CacheProvider } from "@chakra-ui/next-js";
import { ChakraProvider } from "@chakra-ui/react";
import { Noto_Sans_JP } from "next/font/google";

const notoSansJp = Noto_Sans_JP({ subsets: ["latin"] });

export function ChakraProviders({ children }: { children: React.ReactNode }) {
  return (
    <CacheProvider>
      <ChakraProvider>
        {children}
      </ChakraProvider>
    </CacheProvider>
  );
}
