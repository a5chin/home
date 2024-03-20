"use client";
import type { FC } from "react";
import { ChevronLeftIcon } from "@heroicons/react/20/solid";
import { useRouter } from "next/navigation";

import { Flex, Box, Heading, IconButton, Icon } from "@/lib/@chakra-ui/react";

type Props = {
  title: string;
  fixed?: boolean;
  withBackButton?: boolean;
  withBorderBottom?: boolean;
};

export const Header: FC<Props> = ({
  title,
  fixed = false,
  withBackButton = false,
  withBorderBottom = true,
}) => {
  const router = useRouter();
  const back = () => {
    router.back();
  };

  return (
    <Flex
      as="header"
      justifyContent="space-between"
      px="2"
      alignItems="center"
      h="100px"
      w="full"
      bg="brand.white"
      {...(withBorderBottom
        ? { borderBottom: "1px", borderColor: "brand.gray.10" }
        : {})}
      {...(fixed ? { position: "fixed", top: 0, left: 0, zIndex: 20 } : {})}
    >
      <Box w="10" h="10">
        {withBackButton && (
          <IconButton
            isRound
            bg="transparent"
            icon={<Icon as={ChevronLeftIcon} boxSize={6} />}
            aria-label="戻る"
            onClick={back}
          />
        )}
      </Box>
      <Box>
        <Heading size="sm">{title}</Heading>
      </Box>
      <Box w="10" h="10">
        {/* right */}
      </Box>
    </Flex>
  );
};
