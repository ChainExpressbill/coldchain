import axios from 'utils/axios';
import { JoinForm, LoginForm } from './account.model';

export async function login(loginData: LoginForm) {
  const { data } = await axios.post('/login', loginData);
  return data;
}

export async function logout(id: string) {
  const { data } = await axios.post('/logout', { id });
  return data;
}

export async function join(joinData: JoinForm) {
  const { data } = await axios.post('/join', joinData);
  return data;
}
