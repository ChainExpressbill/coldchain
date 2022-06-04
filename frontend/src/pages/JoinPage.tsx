import React from 'react';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { JoinForm } from 'data/account/account.model';
import { useJoin } from 'data/account/account.hooks';

function Join() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<JoinForm>();
  const navigator = useNavigate();

  // TODO: 성공 및 실패 처리
  const { mutate } = useJoin({
    onSuccess(data) {
      console.log('success', data);
      navigator('/');
    },
    onError(err) {
      console.log('error', err);
    },
  });

  const onSubmit = (data: JoinForm) => {
    mutate(data);
  };

  return (
    <div className="text-black h-full flex flex-col align-items-center justify-center gap-12">
      <figure className="w-80 ml-auto mr-auto">
        <img src="logo_square_white.jpg" />
      </figure>
      <form
        onSubmit={handleSubmit(onSubmit)}
        className="w-96 ml-auto mr-auto grid grid-cols-1 gap-8"
      >
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
            placeholder="아이디를 입력해주세요."
            {...register('id', { required: true })}
          />
          {errors.id && (
            <span className="text-red-600 text-sm">아이디를 확인해주세요.</span>
          )}
        </div>
        <div className="h-20">
          <input
            className="border border-gray w-full h-14 p-2"
            placeholder="비밀번호를 입력해주세요."
            type="password"
            {...register('password', {
              required: true,
              minLength: 8,
              maxLength: 15,
            })}
          />
          {errors.password && errors.password.type === 'minLength' && (
            <span className="text-red-600 text-sm">
              비밀번호는 최소 8글자 이상이여야 합니다.
            </span>
          )}
          {errors.password && errors.password.type === 'maxLength' && (
            <span className="text-red-600 text-sm">
              비밀번호는 15글자를 초과할 수 없습니다.
            </span>
          )}
        </div>
        <button
          className="mr-auto ml-auto w-24 p-2 border border-gray"
          type="submit"
        >
          회원가입
        </button>
      </form>
      <div className="w-96 h-11"></div>
    </div>
  );
}

export default Join;
