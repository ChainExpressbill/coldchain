import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useLogout } from 'data/account/account.hooks';

const Header = () => {
  const navigator = useNavigate();

  const { mutate } = useLogout({
    onSuccess() {
      navigator('/');
    },
    onError(err) {
      console.error('error', err);
      window.alert('로그아웃에 실패했습니다! 관리자에게 문의하세요');
    },
  });

  const onLogout = () => {
    mutate({ id: 'admin' });
  };

  return (
    <header className="w-full h-12 flex flex-row px-4 items-center justify-end border-b-[1px] border-main">
      <div className="cursor-pointer" onClick={onLogout}>
        Logout
      </div>
    </header>
  );
};

export default Header;
