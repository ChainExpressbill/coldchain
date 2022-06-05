import React from 'react';
import ModalContainer from './ModalContainer';

interface AlertModalProps {
  msg: string;
  type: string;
  closeModal: () => void;
}

const AlertModal = ({ msg, type, closeModal }: AlertModalProps) => {
  const msgStyle = `text-2xl ${
    type === 'success' ? 'text-black' : 'text-red-500'
  }`;
  return (
    <ModalContainer>
      <div className={msgStyle}>{msg}</div>
      <button
        className="w-16 h-12 border border-black rounded-md text-black"
        onClick={closeModal}
      >
        확인
      </button>
    </ModalContainer>
  );
};

export default AlertModal;
