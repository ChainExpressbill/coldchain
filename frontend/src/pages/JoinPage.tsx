import React from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { JoinForm } from 'data/account/account.model';
import { useJoin } from 'data/account/account.hooks';
import { ModalPortal, AlertModal } from 'components/modal';
import useModal from 'hooks/useModal';
import { AxiosError } from 'axios';

interface CheckPasswordJoinForm extends JoinForm {
  passwordRepeat: string;
}

function Join() {
  const {
    register,
    handleSubmit,
    formState: { errors: formErrors, isValid: isFormValid },
    getValues,
  } = useForm<CheckPasswordJoinForm>({ mode: 'onChange' });
  const navigator = useNavigate();

  const { isOpen, setIsOpen, modalMsg, setModalMsg, modalType, setModalType } =
    useModal();

  const { mutate } = useJoin({
    onSuccess() {
      setModalMsg(() => '회원가입에 성공하였습니다.');
      setModalType(() => 'success');
      setIsOpen(() => true);
    },
    onError(joinError: AxiosError) {
      if (joinError.response !== undefined) {
        const code: string = joinError.response.data.code;
        switch (code) {
          case '500':
            setModalMsg(
              () => '회원가입에 실패했습니다. 관리자에게 문의해주세요.',
            );
            break;
        }
      }
      setModalType(() => 'error');
      setIsOpen(() => true);
    },
  });

  const onSubmit = (data: CheckPasswordJoinForm) => {
    const joinForm: JoinForm = {
      name: data.name,
      emailAddress: data.emailAddress,
      id: data.id,
      password: data.password,
    };
    mutate(joinForm);
  };

  return (
    <div className="text-black h-full flex flex-col align-items-center justify-center gap-12">
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-96 ml-auto mr-auto grid grid-cols-1 gap-8"
      >
        <figure className="w-60 ml-auto mr-auto">
          <img src="logo_square_white.jpg" />
        </figure>
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
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
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
            placeholder="이름을 입력해주세요."
            {...register('name', {
              required: { value: true, message: '이름을 확인해주세요.' },
            })}
          />
          {formErrors.name && (
            <span className="text-red-600 text-sm">
              {' '}
              {formErrors.name.message}
            </span>
          )}
        </div>
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
            placeholder="이메일 주소를 입력해주세요."
            type="email"
            {...register('emailAddress', {
              required: { value: true, message: '이메일 주소를 확인해주세요.' },
            })}
          />
          {formErrors.emailAddress && (
            <span className="text-red-600 text-sm">
              {formErrors.emailAddress.message}
            </span>
          )}
        </div>
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
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
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
            placeholder="비밀번호를 한번 더 입력해주세요."
            type="password"
            {...register('passwordRepeat', {
              required: true,
              validate: (value) =>
                value === getValues().password ||
                '비밀번호가 일치하지 않습니다.',
            })}
          />
          {formErrors.passwordRepeat && (
            <span className="text-red-600 text-sm">
              {formErrors.passwordRepeat?.message}
            </span>
          )}
        </div>
        <button
          className="mr-auto ml-auto mb-10 w-24 p-2 border border-gray disabled:text-white disabled:bg-red-500 disabled:cursor-not-allowed"
          type="submit"
          disabled={!isFormValid}
        >
          회원가입
        </button>
      </form>
      <ModalPortal>
        {isOpen && (
          <AlertModal
            msg={modalMsg}
            type={modalType}
            closeModal={() => {
              setIsOpen(() => false);
              if (modalType === 'success') {
                navigator('/');
              }
            }}
          />
        )}
      </ModalPortal>
    </div>
  );
}

export default Join;
