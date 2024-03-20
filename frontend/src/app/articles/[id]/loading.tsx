import { Container } from "@/lib/@chakra-ui/react";
import { BasicLayout } from "@/component/layout/BasicLayout";
import { ArticleDetailFallback } from "@/app/articles/[id]/_component/server/ArticleDetail";

export default async function Articles() {
  return (
    <BasicLayout title="お悩み詳細" withBackButton>
      <Container py="8" h="100%">
        <ArticleDetailFallback />
      </Container>
    </BasicLayout>
  );
}
