/* eslint-disable */
export type Controller_GetArticleByIDResponse = {
  article?: Entity_Article | undefined;
};

export type Controller_GetArticlesResponse = {
  articles?: Entity_Article[] | undefined;
};

export type Controller_GetTotalViewersResponse = {
  totalViewers?: number | undefined;
};

export type Entity_Article = {
  content?: string | undefined;
  favorites?: number | undefined;
  id?: string | undefined;
  tags?: Entity_Tag[] | undefined;
  title?: string | undefined;
  viewers?: number | undefined;
};

export type Entity_ErrorResponse = {
  message?: string | undefined;
};

export type Entity_Tag = {
  id?: string | undefined;
  name?: string | undefined;
};
