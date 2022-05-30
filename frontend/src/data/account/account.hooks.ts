import { useMutation, UseMutationOptions } from 'react-query';
import { AxiosError } from 'axios';
import { LoginForm, JoinForm } from './account.model';
import { login, logout, join } from './account.api';

export function useLogin(
  options: UseMutationOptions<unknown, AxiosError, LoginForm>,
) {
  return useMutation<unknown, AxiosError, LoginForm>('/login', login, {
    retry: false,
    ...options,
  });
}

export function useLogout(
  options: UseMutationOptions<unknown, AxiosError, { id: string }>,
) {
  return useMutation<unknown, AxiosError, { id: string }>('/logout', logout, {
    retry: false,
    ...options,
  });
}

export function useJoin(
  options: UseMutationOptions<unknown, AxiosError, JoinForm>,
) {
  return useMutation<unknown, AxiosError, JoinForm>('/join', join, {
    retry: false,
    ...options,
  });
}
