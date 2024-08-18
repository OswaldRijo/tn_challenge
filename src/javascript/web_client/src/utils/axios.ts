/**
 * axios setup to use mock service
 */

import axios, { AxiosError } from 'axios';
import Cookies from 'js-cookie';

const axiosServices = axios.create({
  baseURL: process.env.REACT_APP_BACKEND_PATH,
  timeout: 50000,
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: true
});

const AUTH_TOKEN_KEY = 'accessToken';
axiosServices.interceptors.request.use((config) => {
  const token = Cookies.get(AUTH_TOKEN_KEY);

  config.headers.set('Authorization', `Bearer ${token}`);
  return config;
});

// Change response data/error here
axiosServices.interceptors.response.use(
  (response) => response,
  (error) => {
    return Promise.reject(error);
  }
);

export class ServerError extends AxiosError<{ message: string; code: string }, unknown> {}

export default axiosServices;
