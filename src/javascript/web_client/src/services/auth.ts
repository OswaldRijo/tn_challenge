import axios from '@/utils/axios';

export const whoami = async (): Promise<object> => {
  const respose = await axios.post('/users/whoami', { withCredentials: true });
  return respose.data;
};

export const logout = async () => {
  await axios.post('/auth/logout', { withCredentials: true });
};
