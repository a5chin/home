import type { FC, ReactNode } from "react";

import { Header } from "@/component/layout/BasicLayout/Header";
import { Box } from "@/lib/@chakra-ui/react";

type Props = {
  title: string;
  children: ReactNode;
  bg?: "white" | "gray";
  withBackButton?: boolean;
  withBorderBottom?: boolean;
};

export const BasicLayout: FC<Props> = ({
  children,
  title,
  bg = "white",
  withBackButton = false,
  withBorderBottom = true,
}) => {
  const mainBg = bg === "white" ? "brand.white" : "brand.gray.10";
  return (
    <>
      <Header
        title={title}
        fixed
        withBackButton={withBackButton}
        withBorderBottom={withBorderBottom}
      />
      <Box
        pt="100px"
        as="main"
        minH="100vh"
        bg={mainBg}
        display="flex"
        flexDirection="column"
      >
        {children}
      </Box>
    </>
  );
};
