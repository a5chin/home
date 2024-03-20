import { redirect } from "next/navigation";
import type { FC } from "react";

import { Box, SkeletonText } from "@/lib/@chakra-ui/react";
import { getArticle } from "@/service/article/query";

type Props = {
  articleId: string;
};

export const ArticleDetail: FC<Props> = async ({ articleId }) => {
  const article = await getArticle({ articleId: articleId });
  if (article === undefined) {
    redirect("/home");
  }

  return <Box>{article.title}</Box>;
};

export const ArticleDetailFallback = () => (
  <Box>
    <SkeletonText w="80%" noOfLines={1} skeletonHeight="6" />
    <Box mt="8">
      <SkeletonText noOfLines={8} skeletonHeight="4" mt="4" />
    </Box>
  </Box>
);
