"use server";

import { Article } from "@/model/Article";
import { client } from "@/api";
import type { Tag } from "@/model/Tag";
import type { RecursivePartial } from "@/util/type";

export const getArticle = async ({
  articleId,
}: {
  articleId: string;
}): Promise<Article> => {
  console.log(client.articles._articleId(articleId).$path())
  const articleRes = (await client.articles._articleId(articleId).$get()).article;
  console.log(articleRes)

  if (articleRes === undefined) {
    throw new Error("parse error");
  }
  const article: Article = convertToArticle({
    ...articleRes,
  });

  return article;
};

type ArticleData = RecursivePartial<{
  id: string;
  title: string;
  content: string;
  tags: RecursivePartial<Tag[]>;
  viewers: number;
  favorites: number;
}>;

const convertToArticle = (articleData: ArticleData) => {
  return Article.parse(articleData);
};
