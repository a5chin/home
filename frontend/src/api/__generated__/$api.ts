import type { AspidaClient, BasicHeaders } from "aspida";
import { dataToURLString } from "aspida";

import type { Methods as Methods0 } from "./articles";
import type { Methods as Methods1 } from "./articles/_articleId@string";
import type { Methods as Methods2 } from "./articles/totalViewers";

const api = <T>({ baseURL, fetch }: AspidaClient<T>) => {
  const prefix = (baseURL === undefined ? "" : baseURL).replace(/\/$/, "");
  const PATH0 = "/articles";
  const PATH1 = "/articles/totalViewers";
  const GET = "GET";

  return {
    articles: {
      _articleId: (val1: string) => {
        const prefix1 = `${PATH0}/${val1}`;

        return {
          /**
           * @returns OK
           */
          get: (option?: { config?: T | undefined } | undefined) =>
            fetch<
              Methods1["get"]["resBody"],
              BasicHeaders,
              Methods1["get"]["status"]
            >(prefix, prefix1, GET, option).json(),
          /**
           * @returns OK
           */
          $get: (option?: { config?: T | undefined } | undefined) =>
            fetch<
              Methods1["get"]["resBody"],
              BasicHeaders,
              Methods1["get"]["status"]
            >(prefix, prefix1, GET, option)
              .json()
              .then((r) => r.body),
          $path: () => `${prefix}${prefix1}`,
        };
      },
      totalViewers: {
        /**
         * @returns OK
         */
        get: (option?: { config?: T | undefined } | undefined) =>
          fetch<
            Methods2["get"]["resBody"],
            BasicHeaders,
            Methods2["get"]["status"]
          >(prefix, PATH1, GET, option).json(),
        /**
         * @returns OK
         */
        $get: (option?: { config?: T | undefined } | undefined) =>
          fetch<
            Methods2["get"]["resBody"],
            BasicHeaders,
            Methods2["get"]["status"]
          >(prefix, PATH1, GET, option)
            .json()
            .then((r) => r.body),
        $path: () => `${prefix}${PATH1}`,
      },
      /**
       * @returns OK
       */
      get: (
        option?:
          | {
              query?: Methods0["get"]["query"] | undefined;
              config?: T | undefined;
            }
          | undefined
      ) =>
        fetch<
          Methods0["get"]["resBody"],
          BasicHeaders,
          Methods0["get"]["status"]
        >(prefix, PATH0, GET, option).json(),
      /**
       * @returns OK
       */
      $get: (
        option?:
          | {
              query?: Methods0["get"]["query"] | undefined;
              config?: T | undefined;
            }
          | undefined
      ) =>
        fetch<
          Methods0["get"]["resBody"],
          BasicHeaders,
          Methods0["get"]["status"]
        >(prefix, PATH0, GET, option)
          .json()
          .then((r) => r.body),
      $path: (
        option?:
          | { method?: "get" | undefined; query: Methods0["get"]["query"] }
          | undefined
      ) =>
        `${prefix}${PATH0}${
          option && option.query ? `?${dataToURLString(option.query)}` : ""
        }`,
    },
  };
};

export type ApiInstance = ReturnType<typeof api>;
export default api;
