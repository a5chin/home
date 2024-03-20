import type { AxiosResponse } from "axios";
import axios, { isAxiosError } from "axios";
import aspida from "@aspida/axios";
import { redirect } from "next/navigation";

import api from "./__generated__/$api";

// if (process.env.API_URL === undefined) {
//   throw new Error("API_URL environment variable is not defined.");
// }

axios.interceptors.response.use(
  (response: AxiosResponse) => response,
  async (error: unknown) => {
    /* AxiosError */
    if (isAxiosError(error)) {
      if (error?.response?.status === 401) {
        redirect("/auth/signout");
      }
      console.dir(error?.response?.data);
      return Promise.reject();
    }

    console.error("Unexpected error occured");
    console.error(error);

    return Promise.reject();
  },
);

const axiosConfig = { timeout: 1000, baseURL: "http://0.0.0.0:4010" };
export const client = api(aspida(axios, axiosConfig));
