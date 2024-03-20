import { ArticleDetail } from "@/app/articles/[id]/_component/server/ArticleDetail";
import { Container } from "@/lib/@chakra-ui/react";
import { BasicLayout } from "@/component/layout/BasicLayout";

export const dynamic = "force-dynamic";
export default function ArticleDetailPage({
  params,
}: {
  params: { articleId: string };
}) {
  return (
    <BasicLayout title="記事詳細" withBackButton>
      <Container py="8" h="100%">
        <ArticleDetail articleId={params.articleId} />
      </Container>
    </BasicLayout>
  );
}
