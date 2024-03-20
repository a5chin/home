import { z } from "zod";

import dayjs from "@/lib/dayjs";

export const DateTime = z.string().transform((v) => dayjs(v));

export type DateTime = z.infer<typeof DateTime>;
