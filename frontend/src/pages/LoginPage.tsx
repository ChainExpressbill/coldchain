import React from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { useLogin } from 'data/account/account.hooks';
import { LoginForm } from 'data/account/account.model';
import useModal from 'hooks/useModal';
import AlertModal from 'components/modal/AlertModal';
import { AxiosError } from 'axios';

function Login() {
  const {
    register,
    handleSubmit,
    formState: { errors: formErrors, isValid: isFormValid },
  } = useForm<LoginForm>({ mode: 'onChange' });
  const navigator = useNavigate();
  const { isOpen, setIsOpen, modalMsg, setModalMsg, modalType, setModalType } =
    useModal();

  const { mutate } = useLogin({
    onSuccess() {
      navigator('/dashboard');
    },
    onError(loginError: AxiosError) {
      if (loginError.response !== undefined) {
        const code: string = loginError.response.data.code;
        switch (code) {
          case '500':
            setModalMsg(() => '아이디 또는 비밀번호를 잘못 입력했습니다.');
            break;
        }
      }
      setModalType(() => 'error');
      setIsOpen(() => true);
    },
  });

  const onSubmit = (data: LoginForm) => mutate(data);

  return (
    <div className="text-black h-full flex flex-col align-items-center justify-center gap-12">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-96 ml-auto mr-auto grid grid-cols-1 gap-6 Rounded__Shadow__Border p-6"
      >
        <figure className="w-60 ml-auto mr-auto">
          <img src="logo_square_white.jpg" />
        </figure>
        <div>
          <input
            className="Form__Input__Border__Box"
            placeholder="아이디를 입력해주세요."
            {...register('id', {
              required: { value: true, message: '아이디를 확인해주세요.' },
            })}
          />
          {formErrors.id && (
            <span className="text-red-600 text-sm">
              {' '}
              {formErrors.id.message}
            </span>
          )}
        </div>
        <div>
          <input
            className="Form__Input__Border__Box"
            placeholder="비밀번호를 입력해주세요."
            type="password"
            {...register('password', {
              required: true,
              minLength: {
                value: 8,
                message: '비밀번호는 최소 8글자 이상이여야 합니다.',
              },
              maxLength: {
                value: 15,
                message: '비밀번호는 15글자를 초과할 수 없습니다.',
              },
            })}
          />
          {formErrors.password && (
            <span className="text-red-600 text-sm">
              {formErrors.password.message}
            </span>
          )}
        </div>
        <button
          className="mr-auto ml-auto w-full p-2 border border-gray-300 text-white bg-red-500 disabled:opacity-60 disabled:cursor-not-allowed"
          type="submit"
          disabled={!isFormValid}
        >
          로그인
        </button>
      </form>
      <div className="w-96 h-11 mr-auto ml-auto text-right">
        <button
          className="w-24 p-2 border border-gray-300"
          onClick={() => navigator('/join')}
        >
          회원가입
        </button>
      </div>
      {isOpen && (
        <AlertModal
          msg={modalMsg}
          type={modalType}
          closeModal={() => {
            setIsOpen(() => false);
          }}
        />
      )}
    </div>
  );
}

export default Login;
