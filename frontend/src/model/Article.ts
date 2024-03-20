import { z } from "zod";

import { Tag } from "./Tag";

export const Article = z.object({
  id: z.string(),
  title: z.string(),
  content: z.string(),
  tags: z.array(z.nullable(Tag)),
  viewers: z.number().int().min(0),
  favorites: z.number().int().min(0),
});

export type Article = z.infer<typeof Article>;
