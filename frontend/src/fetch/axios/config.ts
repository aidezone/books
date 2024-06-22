// import { AxiosRequestConfig } from "axios";

const axiosDefaultConfig: any = {
  timeout: 1000 * 600,
  headers: {
    "Cache-Control": "no-cache",
    Pragma: "no-cache",
    Expires: "-1",
    language: "en"
  },
  withCredentials: true
};

export default axiosDefaultConfig;
